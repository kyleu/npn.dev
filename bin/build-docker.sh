#!/bin/bash

## Makes a release build, builds a docker image, then exports and zips the output
## Requires docker

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

docker build -m 4g -t kyleu/npn .

mkdir -p build/docker
docker save -o build/docker/npn.docker.tar kyleu/npn
cd build/docker/
rm -f npn.docker.tar.gz
gzip npn.docker.tar
