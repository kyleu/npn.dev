#!/bin/bash

## Runs the Docker image produced by `build-docker`, exposing an HTTP port; Requires docker

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

docker run -m 4g -P kyleu/npn
