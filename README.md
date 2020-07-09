# npn

[npn](https://npn.dev) helps you build stuff

https://github.com/KyleU/npn

## Documentation

- [features.md](doc/features.md)
- [installing.md](doc/installing.md)
- [scripts.md](doc/scripts.md)

## Building

The UI relies on UIkit, run `git submodule init` and `git submodule update` to pull it

For CI servers and go-only changes, simply `make build`. For full stack development, you'll need some tools installed

- Run `bin/bootstrap.sh` to install required Go utilities
- For macOS, you can install all dependencies with `brew install md5sha1sum closure-compiler typescript sass/sass/sass`
- After editing stylesheets, use `bin/build-css.sh`; you'll need `sass` installed
- For TypeScript changes, use `bin/build-client.sh`; you'll need `tsc` and `closure-compiler` installed
- For a developer environment, run `bin/workspace.sh`, which will watch all files and hot-reload
