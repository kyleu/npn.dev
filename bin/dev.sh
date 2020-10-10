#!/bin/bash

## Starts the web server, reloading on changes

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

ulimit -n 2048
air
