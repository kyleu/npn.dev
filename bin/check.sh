#!/bin/bash

## Runs code statistics, checks for outdated dependencies, then runs linters

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

# echo "=== outdated dependecies ==="
# go list -u -m -json all | go-mod-outdated -update

echo "=== linting ==="
golangci-lint run \
  -E asciicheck \
  -E bodyclose \
  -E deadcode \
  -E depguard \
  -E dogsled \
  -D dupl \
  -E errcheck \
  -D funlen \
  -D gochecknoglobals \
  -E gochecknoinits \
  -E gosimple \
  -D gocognit \
  -E goconst \
  -E gocritic \
  -E gocyclo \
  -D godot \
  -E godox \
  -D goerr113 \
  -E gofmt \
  -E goimports \
  -E golint \
  -D gomnd \
  -E gomodguard \
  -E goprintffuncname \
  -D gosec \
  -E gosimple \
  -E govet \
  -E ineffassign \
  -D interfacer \
  -D lll \
  -E maligned \
  -E misspell \
  -E nakedret \
  -E nestif \
  -E nolintlint \
  -E prealloc \
  -E rowserrcheck \
  -E scopelint \
  -E staticcheck \
  -E structcheck \
  -E stylecheck \
  -E testpackage \
  -E typecheck \
  -E unconvert \
  -E unparam \
  -E unused \
  -E varcheck \
  -E whitespace \
  -D wsl \
./...

