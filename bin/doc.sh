#!/bin/bash

## Produces documentation for the project

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "=== generating documentation ==="
