#!/bin/bash

## Mega!

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/tag.sh $1
git push

rm -rf build
bin/build-all.sh
bin/package.sh

gh release create "v$1" -n "v$1"
gh release upload "v$1" build/package/*

bin/deploy.sh
