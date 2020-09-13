#!/bin/bash

## Builds the world
## Requires docker

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

bin/build-client.sh
bin/build-css.sh

bin/build-macos.sh
# bin/build-macos-arm.sh

bin/build-linux.sh
bin/build-linux-arm.sh

bin/build-windows.sh
# bin/build-windows-arm.sh

bin/build-docker.sh
