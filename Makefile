#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
DOCKER := $(shell which docker)
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

# process build tags

# build_tags = netgo
# ifeq ($(LEDGER_ENABLED),true)
#   ifeq ($(OS),Windows_NT)
#     GCCEXE = $(shell where gcc.exe 2> NUL)
#     ifeq ($(GCCEXE),)
#       $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
#     else
#       build_tags += ledger
#     endif
#   else
#     UNAME_S = $(shell uname -s)
#     ifeq ($(UNAME_S),OpenBSD)
#       $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
#     else
#       GCC = $(shell command -v gcc 2> /dev/null)
#       ifeq ($(GCC),)
#         $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
#       else
#         build_tags += ledger
#       endif
#     endif
#   endif
# endif

# ifeq ($(WITH_CLEVELDB),yes)
#   build_tags += gcc
# endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))

# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=archive \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=archived \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags_comma_sep)" -ldflags '$(ldflags)' -trimpath

###############################################################################
###                                  Build                                  ###
###############################################################################

all: build test # TODO: add lint

BUILD_TARGETS := build install

build-x: go.sum 
	go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ ./...

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
	$(DOCKER) rm -f arcbinary || true
	$(DOCKER) create -ti --name arcbinary archive-amd64
	$(DOCKER) cp arcbinary:/bin/archived $(BUILDDIR)/archived-linux-amd64
	$(DOCKER) rm -f arcbinary

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
		sh ./scripts/protocgen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi

proto-image-build:
	@DOCKER_BUILDKIT=1 docker build -t $(protoImageName) -f ./proto/Dockerfile ./proto

proto-image-push:
	docker push $(protoImageName)

###############################################################################
###                               Linting                                   ###
###############################################################################

golangci_lint_cmd=golangci-lint
golangci_version=v1.51.2

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --timeout=10m --config ./.golangci.yml

lint-fix:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0 --config ./.golangci.yml

.PHONY: lint lint-fix

format:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	$(golangci_lint_cmd) run --fix --config ./.golangci.yml
.PHONY: format

###############################################################################
###                                Testing                                  ###
###############################################################################

test:
	@go test -v ./x/...


###############################################################################
###                                Localnet                                 ###
###############################################################################

localnet-init: localnet-clean localnet-build

localnet-build:
	@STATE="" DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker-compose -f tests/localarchive/docker-compose.yml build --progress tty

localnet-start:
	@STATE="" docker-compose -f tests/localarchive/docker-compose.yml up

localnet-clean:
	@rm -rf $(HOME)/.archived-local/

###############################################################################
###                                 Testnet                                 ###
###############################################################################

testnet-init: testnet-clean testnet-build

testnet-build:
	@DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker-compose -f tests/testnet/docker-compose.yml build

testnet-start:
	docker-compose -f tests/testnet/docker-compose.yml up

testnet-clean:
	@rm -rf $(HOME)/.archived-local-0/
	@rm -rf $(HOME)/.archived-local-1/
	@rm -rf $(HOME)/.archived-local-2/
