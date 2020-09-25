#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "Embedding assets..."
go-embed -input npnasset/vendor -output npnasset/assets/assets.go
go-embed -input web/assets -output app/assets/assets.go
