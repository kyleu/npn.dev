#!/bin/bash

## Formatting code from all projects

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "=== formatting ==="
gofmt -w .
echo "=== organizing imports ==="
goimports -w .
