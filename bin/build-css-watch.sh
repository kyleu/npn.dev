#!/bin/bash

## Builds the css resources using `build-css`, then watches for changes in `stylesheets`
## Requires SCSS available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build-css.sh
echo "Watching sass compilation for [web/stylesheets/style.scss]..."
sass --load-path=npnasset/stylesheets --watch --no-source-map web/stylesheets/style.scss web/assets/vendor/npn.css
