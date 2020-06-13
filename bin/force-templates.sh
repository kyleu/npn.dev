#!/bin/sh

## Builds all the templates using hero

echo "updating [web/components] templates"
rm -rf gen/components
hero -extensions .html -source "web/components" -pkgname components -dest gen/components

echo "updating [web/templates] templates"
rm -rf gen/templates
hero -extensions .html -source "web/templates" -pkgname templates -dest gen/templates
