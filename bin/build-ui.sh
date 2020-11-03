#!/bin/bash

## Attempts to build the UI application

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

cd ui
bin/copy.sh
