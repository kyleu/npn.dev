#!/bin/bash

## Updates the go.mod version, deletes go.sum, tags the git repo

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

TGT=$1

[ "$TGT" ] || (echo `must provide one argument, like "0.0.1"` && exit)

rplc () {
  rm -f go.sum
  find . -type f -name "go.mod" -print0 | xargs -0 sed -i '' -e "s/v[01]\.[0-9]*[0-9]\.[0-9]*[0-9] \/\/ npn/v${TGT} \/\/ npn/g"
  sed -i '' -e "s/v[01]\.[0-9]*[0-9]\.[0-9]*[0-9]/v${TGT}/g" doc/downloads.md
}

bld () {
  make build
  git add .
  git commit -m "v${TGT}"
}

gt () {
  git tag "$1/v${TGT}"
}

tagall () {
  git tag "v${TGT}"

  gt "npnasset"
  gt "npnconnection"
  gt "npncontroller"
  gt "npncore"
  gt "npndatabase"
  gt "npnexport"
  gt "npngraphql"
  gt "npnscript"
  gt "npnservice"
  gt "npnservice-db"
  gt "npnservice-fs"
  gt "npntemplate"
  gt "npnuser"
  gt "npnweb"
}

rplc
bld
tagall
