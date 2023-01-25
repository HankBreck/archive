#!/usr/bin/env bash

set -eo pipefail

# get protoc executions
go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos 2>/dev/null

# # get cosmos sdk from github
# go get github.com/cosmos/cosmos-sdk 2>/dev/null

echo "Generating gogo proto code"
cd proto
buf generate --template buf.gen.gogo.yaml

cd ..

# move proto files to the right places
#
# Note: Proto files are suffixed with the current binary version.
cp -r github.com/HankBreck/archive/* ./
rm -rf github.com

go mod tidy -compat=1.18