#!/bin/bash

## Builds all the templates using hero, skipping if unchanged

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

FORCE="$1"

function tmpl {
  echo "updating [$2] templates"
  if test -f "$ftgt"; then
    mv "$ftgt" "$fsrc"
  fi
  if [ "$1" = "npntemplate" ]; then
    cd npntemplate
    rm -rf gen
    hero -extensions .html,.sql -source "html" -pkgname $1 -dest "gen/npntemplate"
    cd ..
  else
    rm -rf $3
    hero -extensions .html,.sql -source "$2" -pkgname $1 -dest $3
  fi
}

function check {
  fsrc="tmp/$1.hashcode"
  ftgt="tmp/$1.hashcode.tmp"

  if [ ! -d "$3" ]; then
    rm -f "$fsrc"
  fi

  mkdir -p tmp/

  find "$2" -type f -exec md5sum {} \; | md5sum > "$ftgt"

  if cmp -s "$fsrc" "$ftgt"; then
    if [ "$FORCE" = "force" ]; then
      tmpl $1 $2 $3 $4
    else
      rm "$ftgt"
    fi
  else
    tmpl $1 $2 $3 $4
  fi
}

check "npntemplate" "npntemplate/html" "npntemplate/gen"
check "templates" "web/templates" "gen/templates"
