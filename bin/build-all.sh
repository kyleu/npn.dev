#!/bin/bash

## Builds the world
## Requires docker

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

bin/build-client.sh
bin/build-css.sh

bin/build-desktop.sh

bin/build-docker.sh

bin/build-wasm.sh

bin/build-android.sh
bin/build-ios.sh

