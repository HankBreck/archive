#!/usr/bin/env bash

set -eo pipefail

# get protoc executions
go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos 2>/dev/null

# # get cosmos sdk from github
# go get github.com/cosmos/cosmos-sdk 2>/dev/null

echo "Generating gogo proto code"
cd proto
echo "1"
proto_dirs=$(find ./archive -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
echo "2"
# for dir in $proto_dirs; do
#   for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
#     if grep go_package $file &>/dev/null; then
#       buf generate --template buf.gen.gogo.yaml $file
#     fi
#   done
# done
echo "Running cmd"
buf generate --template buf.gen.gogo.yaml
echo "3"

cd ..

# move proto files to the right places
#
# Note: Proto files are suffixed with the current binary version.
cp -r github.com/HankBreck/archive/* ./
rm -rf github.com

echo "4"

go mod tidy -compat=1.18