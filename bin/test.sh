#!/bin/bash

## Runs all the tests

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

t () {
  go test github.com/kyleu/npn/$1
}

t "app"
