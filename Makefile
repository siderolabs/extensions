REGISTRY ?= ghcr.io
USERNAME ?= siderolabs
SHA ?= $(shell git describe --match=none --always --abbrev=8 --dirty)
TAG ?= $(shell git describe --tag --always --dirty)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
REGISTRY_AND_USERNAME := $(REGISTRY)/$(USERNAME)

BUILD := docker buildx build
PLATFORM ?= linux/amd64,linux/arm64
PROGRESS ?= auto
PUSH ?= false
COMMON_ARGS := --file=Pkgfile
COMMON_ARGS += --progress=$(PROGRESS)
COMMON_ARGS += --platform=$(PLATFORM)
COMMON_ARGS += --build-arg=http_proxy=$(http_proxy)
COMMON_ARGS += --build-arg=https_proxy=$(https_proxy)

, := ,
empty :=
space = $(empty) $(empty)

TARGETS = amd-ucode bnx2-bnx2x gvisor hello-world-service intel-ucode
NONFREE_TARGETS = nvidia-container-toolkit

all: $(TARGETS) ## Builds all known pkgs.

nonfree: $(NONFREE_TARGETS) ## Builds all known non-free pkgs.

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
$(TARGETS) $(NONFREE_TARGETS):
	@$(MAKE) docker-$@ TARGET_ARGS="--tag=$(REGISTRY)/$(USERNAME)/$@:$(TAG) --push=$(PUSH)"

.PHONY: deps.png
deps.png:
	bldr graph | dot -Tpng > deps.png

.PHONY: conformance
conformance: ## Performs policy checks against the commit and source code.
	docker run --rm -it -v $(PWD):/src -w /src ghcr.io/siderolabs/conform:latest enforce
