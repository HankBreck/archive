#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
DOCKER := $(shell which docker)
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=archive \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=archived \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

###############################################################################
###                                  Build                                  ###
###############################################################################

all: install

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

# Cross-building for arm64 from amd64 (or viceversa) takes
# a lot of time due to QEMU virtualization but it's the only way (afaik)
# to get a statically linked binary with CosmWasm

build-reproducible: build-reproducible-amd64 build-reproducible-arm64

build-reproducible-amd64: go.sum $(BUILDDIR)/
	$(DOCKER) buildx create --name arcbuilder || true
	$(DOCKER) buildx use arcbuilder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f 2) \
		--platform linux/amd64 \
		-t archive-amd64 \
		--load \
		-f Dockerfile .

build-reproducible-arm64: go.sum $(BUILDDIR)/
	$(DOCKER) buildx create --name arcbuilder || true
	$(DOCKER) buildx use arcbuilder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f 2) \
		--platform linux/arm64 \
		-t archive-arm64 \
		--load \
		-f Dockerfile .
	$(DOCKER) rm -f arcbinary || true
	$(DOCKER) create -ti --name arcbinary archive-arm64
	$(DOCKER) cp arcbinary:/bin/archived $(BUILDDIR)/archived-linux-arm64
	$(DOCKER) rm -f arcbinary

###############################################################################
###                                  Proto                                  ###
###############################################################################

proto-all: proto-format proto-gen

proto:
	@echo
	@echo "=========== Generate Message ============"
	@echo
	./scripts/protocgen.sh
	@echo
	@echo "=========== Generate Complete ============"
	@echo

protoVer=v0.8
protoImageName=hankbreck/archive-proto-gen:$(protoVer)
containerProtoGen=cosmos-sdk-proto-gen-$(protoVer)
containerProtoFmt=cosmos-sdk-proto-fmt-$(protoVer)

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/generate-proto.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi

proto-image-build:
	@DOCKER_BUILDKIT=1 docker build -t $(protoImageName) -f ./proto/Dockerfile ./proto

proto-image-push:
	docker push $(protoImageName)


###############################################################################
###                                Localnet                                 ###
###############################################################################

localnet-init: localnet-clean localnet-build

localnet-build:
	@DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker-compose -f tests/localarchive/docker-compose.yml build

localnet-start:
	@STATE="" docker-compose -f tests/localarchive/docker-compose.yml up

localnet-clean:
	@rm -rf $(HOME)/.archived-local/

