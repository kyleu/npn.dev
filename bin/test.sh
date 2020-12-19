#!/bin/bash

## Runs all the tests

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

t () {
  go test github.com/kyleu/npn/$1
}

t "npncore"

t "npnconnection"
t "npncontroller"
t "npndatabase"
t "npnexport/pdf"
t "npnexport/xls"
t "npngraphql"
t "npnqueue"
t "npnscript/js"
t "npnscript/lua"
# t "npnservice"
t "npnservice-db/authdb"
t "npnservice-db/userdb"
t "npnservice-fs/authfs"
t "npnservice-fs/userfs"
t "npntemplate/gen/npntemplate"
t "npnuser"
t "npnweb"

t "app"
