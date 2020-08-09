#!/bin/bash

## Builds all the templates using hero, skipping if unchanged

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

FORCE="$1"

function tmpl {
  echo "updating [$2] templates"
  mv "$ftgt" "$fsrc"
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

  find -s "$2" -type f -exec md5sum {} \; | md5sum > "$ftgt"

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
