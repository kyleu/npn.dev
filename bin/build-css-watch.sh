#!/bin/bash

## Builds the css resources using `build-css`, then watches for changes in `stylesheets`
## Requires SCSS available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

bin/build-css.sh
echo "Watching sass compilation for [web/stylesheets/style.scss]..."
sass --watch --no-source-map web/stylesheets/style.scss web/assets/vendor/style.css
