#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

# require argument

[ "$#" -eq 2 ] || [ "$#" -eq 3 ] || die "pass [os arch] as arguments "

os=$1
arch=$2
fn=${3:-npn}

# default argument

echo "Building [$os $arch]..."
env GOOS=$os GOARCH=$arch make build-release
mkdir -p ./build/$os/$arch
mv "./build/release/$fn" "./build/$os/$arch/$fn"
