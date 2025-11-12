# THIS FILE WAS AUTOMATICALLY GENERATED, PLEASE DO NOT EDIT.
#
# Generated on 2025-11-12T17:46:46Z by kres 911d166.

# common variables

SHA := $(shell git describe --match=none --always --abbrev=8 --dirty)
TAG := $(shell git describe --tag --always --dirty --match v[0-9]\*)
ABBREV_TAG := $(shell git describe --tags >/dev/null 2>/dev/null && git describe --tag --always --match v[0-9]\* --abbrev=0 || echo 'undefined')
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
ARTIFACTS := _out
IMAGE_TAG ?= $(TAG)
OPERATING_SYSTEM := $(shell uname -s | tr '[:upper:]' '[:lower:]')
GOARCH := $(shell uname -m | sed 's/x86_64/amd64/' | sed 's/aarch64/arm64/')
REGISTRY ?= ghcr.io
USERNAME ?= siderolabs
REGISTRY_AND_USERNAME ?= $(REGISTRY)/$(USERNAME)
KRES_IMAGE ?= ghcr.io/siderolabs/kres:latest
CONFORMANCE_IMAGE ?= ghcr.io/siderolabs/conform:latest

# source date epoch of first commit

INITIAL_COMMIT_SHA := $(shell git rev-list --max-parents=0 HEAD)
SOURCE_DATE_EPOCH := $(shell git log $(INITIAL_COMMIT_SHA) --pretty=%ct)

# sync bldr image with pkgfile

BLDR_RELEASE := v0.5.5
BLDR_IMAGE := ghcr.io/siderolabs/bldr:$(BLDR_RELEASE)
BLDR := docker run --rm --user $(shell id -u):$(shell id -g) --volume $(PWD):/src --entrypoint=/bldr $(BLDR_IMAGE) --root=/src

# docker build settings

BUILD := docker buildx build
PLATFORM ?= linux/amd64,linux/arm64
PROGRESS ?= auto
PUSH ?= false
CI_ARGS ?=
BUILD_ARGS = --build-arg=SOURCE_DATE_EPOCH=$(SOURCE_DATE_EPOCH)
BUILD_ARGS += --build-arg=TAG="$(TAG)"
BUILD_ARGS += --build-arg=PKGS="$(PKGS)"
BUILD_ARGS += --build-arg=PKGS_PREFIX="$(PKGS_PREFIX)"
BUILD_ARGS += --build-arg=TOOLS="$(TOOLS)"
BUILD_ARGS += --build-arg=TOOLS_PREFIX="$(TOOLS_PREFIX)"
COMMON_ARGS = --file=Pkgfile
COMMON_ARGS += --provenance=false
COMMON_ARGS += --progress=$(PROGRESS)
COMMON_ARGS += --platform=$(PLATFORM)
COMMON_ARGS += $(BUILD_ARGS)

# extra variables

EXTENSIONS_IMAGE_REF ?= $(REGISTRY_AND_USERNAME)/extensions:$(TAG)
PKGS ?= v1.12.0
PKGS_PREFIX ?= ghcr.io/siderolabs
TOOLS ?= v1.12.0
TOOLS_PREFIX ?= ghcr.io/siderolabs
IMAGE_SIGNER_RELEASE ?= v0.1.1

# targets defines all the available targets

TARGETS = amazon-ena
TARGETS += amdgpu
TARGETS += amd-ucode
TARGETS += binfmt-misc
TARGETS += bnx2-bnx2x
TARGETS += btrfs
TARGETS += chelsio-drivers
TARGETS += chelsio-firmware
TARGETS += cloudflared
TARGETS += crun
TARGETS += ctr
TARGETS += drbd
TARGETS += dvb-cx23885
TARGETS += dvb-m88ds3103
TARGETS += ecr-credential-provider
TARGETS += fuse3
TARGETS += gasket-driver
TARGETS += glibc
TARGETS += gvisor
TARGETS += gvisor-debug
TARGETS += hailort
TARGETS += hello-world-service
TARGETS += i915
TARGETS += intel-ice-firmware
TARGETS += intel-ucode
TARGETS += iscsi-tools
TARGETS += kata-containers
TARGETS += lldpd
TARGETS += mdadm
TARGETS += mei
TARGETS += metal-agent
TARGETS += nebula
TARGETS += netbird
TARGETS += newt
TARGETS += nfs-utils
TARGETS += nfsd
TARGETS += nfsrahead
TARGETS += nut-client
TARGETS += nvidia-container-toolkit-lts
TARGETS += nvidia-container-toolkit-production
TARGETS += nvidia-fabricmanager-lts
TARGETS += nvidia-fabricmanager-production
TARGETS += nvidia-gdrdrv-device
TARGETS += nvidia-open-gpu-kernel-modules-lts
TARGETS += nvidia-open-gpu-kernel-modules-production
TARGETS += nvme-cli
TARGETS += panfrost
TARGETS += qemu-guest-agent
TARGETS += qlogic-firmware
TARGETS += realtek-firmware
TARGETS += revpi-firmware
TARGETS += spin
TARGETS += stargz-snapshotter
TARGETS += tailscale
TARGETS += tenstorrent
TARGETS += thunderbolt
TARGETS += uinput
TARGETS += usb-modem-drivers
TARGETS += usb-audio-drivers
TARGETS += util-linux-tools
TARGETS += v4l-uvc-drivers
TARGETS += vc4
TARGETS += vmtoolsd-guest-agent
TARGETS += wasmedge
TARGETS += xdma-driver
TARGETS += xe
TARGETS += xen-guest-agent
TARGETS += youki
TARGETS += zerotier
TARGETS += zfs
NONFREE_TARGETS = nonfree-kmod-nvidia-lts
NONFREE_TARGETS += nonfree-kmod-nvidia-production

# help menu

export define HELP_MENU_HEADER
# Getting Started

To build this project, you must have the following installed:

- git
- make
- docker (19.03 or higher)

## Creating a Builder Instance

The build process makes use of experimental Docker features (buildx).
To enable experimental features, add 'experimental: "true"' to '/etc/docker/daemon.json' on
Linux or enable experimental features in Docker GUI for Windows or Mac.

To create a builder instance, run:

	docker buildx create --name local --use

If running builds that needs to be cached aggresively create a builder instance with the following:

	docker buildx create --name local --use --config=config.toml

config.toml contents:

[worker.oci]
  gc = true
  gckeepstorage = 50000

  [[worker.oci.gcpolicy]]
    keepBytes = 10737418240
    keepDuration = 604800
    filters = [ "type==source.local", "type==exec.cachemount", "type==source.git.checkout"]
  [[worker.oci.gcpolicy]]
    all = true
    keepBytes = 53687091200

If you already have a compatible builder instance, you may use that instead.

## Artifacts

All artifacts will be output to ./$(ARTIFACTS). Images will be tagged with the
registry "$(REGISTRY)", username "$(USERNAME)", and a dynamic tag (e.g. $(IMAGE):$(IMAGE_TAG)).
The registry and username can be overridden by exporting REGISTRY, and USERNAME
respectively.

endef

all: $(TARGETS)  ## Builds all targets defined.

$(ARTIFACTS):  ## Creates artifacts directory.
	@mkdir -p $(ARTIFACTS)

.PHONY: clean
clean:  ## Cleans up all artifacts.
	@rm -rf $(ARTIFACTS)

target-%:  ## Builds the specified target defined in the Pkgfile. The build result will only remain in the build cache.
	@$(BUILD) --target=$* $(COMMON_ARGS) $(TARGET_ARGS) $(CI_ARGS) .

local-%:  ## Builds the specified target defined in the Pkgfile using the local output type. The build result will be output to the specified local destination.
	@$(MAKE) target-$* TARGET_ARGS="--output=type=local,dest=$(DEST) $(TARGET_ARGS)"

docker-%:  ## Builds the specified target defined in the Pkgfile using the docker output type. The build result will be loaded into Docker.
	@$(MAKE) target-$* TARGET_ARGS="$(TARGET_ARGS)"

reproducibility-test:  ## Builds the reproducibility test target
	@$(MAKE) reproducibility-test-local-reproducibility

reproducibility-test-local-%:  ## Builds the specified target defined in the Pkgfile using the local output type with and without cahce. The build result will be output to the specified local destination
	@rm -rf $(ARTIFACTS)/build-a $(ARTIFACTS)/build-b
	@$(MAKE) local-$* DEST=$(ARTIFACTS)/build-a
	@$(MAKE) local-$* DEST=$(ARTIFACTS)/build-b TARGET_ARGS="--no-cache"
	@touch -ch -t $$(date -d @$(SOURCE_DATE_EPOCH) +%Y%m%d0000) $(ARTIFACTS)/build-a $(ARTIFACTS)/build-b
	@diffoscope $(ARTIFACTS)/build-a $(ARTIFACTS)/build-b
	@rm -rf $(ARTIFACTS)/build-a $(ARTIFACTS)/build-b

$(ARTIFACTS)/bldr: $(ARTIFACTS)  ## Downloads bldr binary.
	@curl -sSL https://github.com/siderolabs/bldr/releases/download/$(BLDR_RELEASE)/bldr-$(OPERATING_SYSTEM)-$(GOARCH) -o $(ARTIFACTS)/bldr
	@chmod +x $(ARTIFACTS)/bldr

.PHONY: update-checksums
update-checksums: $(ARTIFACTS)/bldr  ## Updates the checksums in the Pkgfile/vars.yaml based on the changed version variables.
	@git diff -U0 | $(ARTIFACTS)/bldr update

nonfree: $(NONFREE_TARGETS)  ## Builds all nonfree targets defined.

.PHONY: $(TARGETS) $(NONFREE_TARGETS)
$(TARGETS) $(NONFREE_TARGETS): $(ARTIFACTS)/bldr
	@$(MAKE) docker-$@ TARGET_ARGS="--tag=$(REGISTRY)/$(USERNAME)/$@:$(shell $(ARTIFACTS)/bldr eval --target $@ --build-arg TAG=$(TAG) '{{.VERSION}}' 2>/dev/null) --push=$(PUSH)"

.PHONY: deps.svg
deps.svg:  ## Generates a dependency graph of the Pkgfile.
	@rm -f deps.png
	@$(BLDR) graph $(BUILD_ARGS) | dot -Tsvg -o deps.svg

.PHONY: extensions
extensions: internal/extensions/descriptions.yaml
	@$(MAKE) docker-$@ TARGET_ARGS="--tag=$(EXTENSIONS_IMAGE_REF) --push=$(PUSH)"

.PHONY: extensions-catalog
extensions-catalog: $(ARTIFACTS)/bldr
	@$(ARTIFACTS)/bldr dump --build-arg TAG=VERSION --template hack/catalog.template > $(ARTIFACTS)/catalog.md 2>/dev/null
	@lead='^<!-- ### BEGIN GENERATED CONTENT -->$$'; tail='^<!-- ### END GENERATED CONTENT -->$$'; sed -i -e "/$$lead/,/$$tail/{ /$$lead/{p; r $(ARTIFACTS)/catalog.md" -e "}; /$$tail/p; d }" README.md

.PHONY: check-dirty
check-dirty:
	@if test -n "`git status --porcelain`"; then echo "Source tree is dirty"; git status; git diff; exit 1 ; fi

.PHONY: extensions-metadata
extensions-metadata: $(ARTIFACTS)/bldr
	@rm -f _out/extensions-metadata
	@$(foreach target,$(TARGETS),echo $(REGISTRY)/$(USERNAME)/$(target):$(shell $(ARTIFACTS)/bldr eval --target $(target) --build-arg TAG=$(TAG) '{{.VERSION}}' 2>/dev/null) >> _out/extensions-metadata;)
	@$(foreach target,$(NONFREE_TARGETS),echo $(REGISTRY)/$(USERNAME)/$(target):$(shell $(ARTIFACTS)/bldr eval --target $(target) --build-arg TAG=$(TAG) '{{.VERSION}}' 2>/dev/null) >> _out/extensions-metadata;)

.PHONY: internal/extensions/image-digests
internal/extensions/image-digests: extensions-metadata
	@echo "Generating image digests..."
	@cat _out/extensions-metadata | xargs -I{} sh -c 'echo {}@$$(crane digest {})' > internal/extensions/image-digests

.PHONY: internal/extensions/descriptions.yaml
internal/extensions/descriptions.yaml: internal/extensions/image-digests
	@echo "Generating image descriptions..."
	@echo -n "" > internal/extensions/descriptions.yaml
	@for image in $(shell cat internal/extensions/image-digests); do \
	  crane export $$image - | tar x -O --occurrence=1 manifest.yaml | yq -r ". += {\"$$image\": {\"author\": .metadata.author, \"description\": .metadata.description}} | del(.metadata, .version)" - >> internal/extensions/descriptions.yaml; \
	done

.PHONY: $(ARTIFACTS)/image-signer
$(ARTIFACTS)/image-signer:
	@curl -sSL https://github.com/siderolabs/go-tools/releases/download/$(IMAGE_SIGNER_RELEASE)/image-signer-$(OPERATING_SYSTEM)-$(GOARCH) -o $(ARTIFACTS)/image-signer
	@chmod +x $(ARTIFACTS)/image-signer

.PHONY: sign-images
sign-images: $(ARTIFACTS)/image-signer
	@$(ARTIFACTS)/image-signer sign --timeout=15m $(shell crane export $(EXTENSIONS_IMAGE_REF) | tar x --to-stdout image-digests) $(EXTENSIONS_IMAGE_REF)@$$(crane digest $(EXTENSIONS_IMAGE_REF))

.PHONY: grype-scan
grype-scan:
	@$(MAKE) local-$@ DEST=$(ARTIFACTS)/grype-scan PLATFORM=linux/amd64

.PHONY: rekres
rekres:
	@docker pull $(KRES_IMAGE)
	@docker run --rm --net=host --user $(shell id -u):$(shell id -g) -v $(PWD):/src -w /src -e GITHUB_TOKEN $(KRES_IMAGE)

.PHONY: help
help:  ## This help menu.
	@echo "$$HELP_MENU_HEADER"
	@grep -E '^[a-zA-Z%_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: release-notes
release-notes: $(ARTIFACTS)
	@ARTIFACTS=$(ARTIFACTS) ./hack/release.sh $@ $(ARTIFACTS)/RELEASE_NOTES.md $(TAG)

.PHONY: conformance
conformance:
	@docker pull $(CONFORMANCE_IMAGE)
	@docker run --rm -it -v $(PWD):/src -w /src $(CONFORMANCE_IMAGE) enforce

