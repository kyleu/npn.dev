#!/bin/bash

## Uses `scss` to compile the stylesheets in `web/stylesheets`
## Requires SCSS available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

sass --no-source-map web/stylesheets/style.scss web/assets/vendor/style.css
sass --style=compressed --no-source-map web/stylesheets/style.scss web/assets/vendor/style.min.css
