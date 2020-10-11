#!/bin/bash

## Temp deploy

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build.sh linux amd64
../kyleu.dev/deploy/npn.sh
../kyleu.dev/shell.sh
