#!/bin/bash

## Formatting code from all projects

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

echo "=== formatting ==="
gofmt -w .
echo "=== organizing imports ==="
goimports -w .
