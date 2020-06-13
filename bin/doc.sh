#!/bin/bash

## Produces documentation for the project

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

echo "=== generating documentation ==="
