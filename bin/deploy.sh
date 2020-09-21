#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

bin/build.sh linux amd64
../kyleu.dev/deploy/npn.sh
../kyleu.dev/shell.sh
