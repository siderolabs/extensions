REGISTRY ?= ghcr.io
USERNAME ?= siderolabs
SHA ?= $(shell git describe --match=none --always --abbrev=8 --dirty)
TAG ?= $(shell git describe --tag --always --dirty)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
REGISTRY_AND_USERNAME := $(REGISTRY)/$(USERNAME)
# inital commit time
# git rev-list --max-parents=0 HEAD
# git log a46b3f24d158614d582da5e6e7e34b596d10cb8e --pretty=%ct
SOURCE_DATE_EPOCH ?= "1642703752"
ARTIFACTS ?= _out/
OPERATING_SYSTEM := $(shell uname -s | tr "[:upper:]" "[:lower:]")
GOARCH := $(shell uname -m | tr '[:upper:]' '[:lower:]')
EXTENSIONS_IMAGE_REF := $(REGISTRY)/$(USERNAME)/extensions:$(TAG)

ifeq ($(GOARCH),x86_64)
GOARCH := amd64
endif

# keep in sync with Pkgfile
BLDR_RELEASE ?= v0.2.2
PKGS ?= v1.6.0-alpha.0-20-g0c84090

BUILD := docker buildx build
PLATFORM ?= linux/amd64,linux/arm64
PROGRESS ?= auto
PUSH ?= false
COMMON_ARGS := --file=Pkgfile
COMMON_ARGS += --provenance=false
COMMON_ARGS += --progress=$(PROGRESS)
COMMON_ARGS += --platform=$(PLATFORM)
COMMON_ARGS += --build-arg=http_proxy=$(http_proxy)
COMMON_ARGS += --build-arg=https_proxy=$(https_proxy)
COMMON_ARGS += --build-arg=SOURCE_DATE_EPOCH=$(SOURCE_DATE_EPOCH)
COMMON_ARGS += --build-arg=TAG=$(TAG)
COMMON_ARGS += --build-arg=PKGS=$(PKGS)

, := ,
empty :=
space = $(empty) $(empty)

TARGETS = \
		amdgpu-firmware \
		amd-ucode \
		binfmt-misc \
		bnx2-bnx2x \
		btrfs \
		chelsio-drivers \
		chelsio-firmware \
		drbd \
		fuse3 \
		gasket-driver \
		gvisor \
		hello-world-service \
		i915-ucode \
		intel-ice-firmware \
		intel-ucode \
		iscsi-tools \
		nut-client \
		nvidia-container-toolkit \
		nvidia-fabricmanager \
		nvidia-open-gpu-kernel-modules \
		qemu-guest-agent \
		stargz-snapshotter \
		tailscale \
		thunderbolt \
		usb-modem-drivers \
		util-linux-tools \
		xe-guest-utilities \
		zfs

# Temporarily disabled, as mellanox-ofed fails to build with Linux 6.1
#		mellanox-ofed \

NONFREE_TARGETS = \
	nonfree-kmod-nvidia \

all: $(TARGETS) ## Builds all known pkgs.

.PHONY: nonfree
nonfree: $(NONFREE_TARGETS) ## Builds all known non-free pkgs.

.PHONY: extensions
extensions: internal/extensions/image-digests ## Builds a list of generated extension images as an image.
	@$(MAKE) docker-$@ \
		TARGET_ARGS="--tag=$(EXTENSIONS_IMAGE_REF) --push=$(PUSH)"

.PHONY: help
help: ## This help menu.
	@grep -E '^[a-zA-Z%_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

local-%: ## Builds the specified target defined in the Dockerfile using the local output type. The build result will be output to the specified local destination.
	@$(MAKE) target-$* TARGET_ARGS="--output=type=local,dest=$(DEST) $(TARGET_ARGS)"
	@PLATFORM=$(PLATFORM) \

target-%: ## Builds the specified target defined in the Dockerfile. The build result will only remain in the build cache.
	@$(BUILD) \
		--target=$* \
		$(COMMON_ARGS) \
		$(TARGET_ARGS) .

docker-%: ## Builds the specified target defined in the Dockerfile using the docker output type. The build result will be loaded into docker.
	@$(MAKE) target-$* TARGET_ARGS="$(TARGET_ARGS)"

.PHONY: $(TARGETS) $(NONFREE_TARGETS)
$(TARGETS) $(NONFREE_TARGETS): $(ARTIFACTS)/bldr
	@$(MAKE) docker-$@ \
		TARGET_ARGS="--tag=$(REGISTRY)/$(USERNAME)/$@:$(shell $(ARTIFACTS)/bldr eval --target $@ --build-arg TAG=$(TAG) '{{.VERSION}}' 2>/dev/null) --push=$(PUSH)"

extensions-metadata: $(ARTIFACTS)/bldr
	@rm -f _out/extensions-metadata
	@$(foreach target,$(TARGETS),echo $(REGISTRY)/$(USERNAME)/$(target):$(shell $(ARTIFACTS)/bldr eval --target $(target) --build-arg TAG=$(TAG) '{{.VERSION}}' 2>/dev/null) >> _out/extensions-metadata;)
	@$(foreach target,$(NONFREE_TARGETS),echo $(REGISTRY)/$(USERNAME)/$(target):$(shell $(ARTIFACTS)/bldr eval --target $(target) --build-arg TAG=$(TAG) '{{.VERSION}}' 2>/dev/null) >> _out/extensions-metadata;)

.PHONY: internal/extensions/image-digests
internal/extensions/image-digests: extensions-metadata ## Stores a list of all images built by this Makefile with digests.
	@cat _out/extensions-metadata | xargs -I{} sh -c 'echo {}@$$(crane digest {})' > internal/extensions/image-digests

.PHONY: sign-images
sign-images: ## Run cosign to sign all images built by this Makefile.
	@for image in $(shell crane export $(EXTENSIONS_IMAGE_REF) | tar x --to-stdout image-digests) $(EXTENSIONS_IMAGE_REF)@$$(crane digest $(EXTENSIONS_IMAGE_REF)); do \
		echo '==>' $$image; \
		cosign verify $$image --certificate-identity-regexp '@siderolabs\.com$$' --certificate-oidc-issuer https://accounts.google.com || \
			cosign sign --yes $$image; \
	done

.PHONY: deps.png
deps.png: $(ARTIFACTS)/bldr
	$(ARTIFACTS)/bldr graph | dot -Tpng > deps.png

.PHONY: release-notes
release-notes:
	mkdir -p $(ARTIFACTS)
	@ARTIFACTS=$(ARTIFACTS) ./hack/release.sh $@ $(ARTIFACTS)/RELEASE_NOTES.md $(TAG)

.PHONY: conformance
conformance: ## Performs policy checks against the commit and source code.
	docker run --rm -it -v $(PWD):/src -w /src ghcr.io/siderolabs/conform:latest enforce

$(ARTIFACTS)/bldr:
	@mkdir -p $(ARTIFACTS)
	@curl -L https://github.com/siderolabs/bldr/releases/download/$(BLDR_RELEASE)/bldr-$(OPERATING_SYSTEM)-$(GOARCH) -o $(ARTIFACTS)/bldr
	@chmod +x $(ARTIFACTS)/bldr
