# npn

[npn](https://npn.dev) helps you build stuff

https://github.com/kyleu/npn

## Documentation

- [downloads.md](doc/downloads.md)
- [features.md](doc/features.md)
- [installing.md](doc/installing.md)
- [scripts.md](doc/scripts.md)

## Installing

- Download from [Github Releases](https://github.com/kyleu/npn/releases/latest), or see the [downloads page](doc/downloads.md) 

## Building

- Run `bin/bootstrap.sh` to install required Go utilities
- Run `make build` to produce a binary in `./build`, or run `bin/dev.sh` to recompile and restart automatically

For full stack development, you'll need some tools installed:

- For TypeScript changes, use `bin/build-client.sh`; you'll need `tsc` and `closure-compiler` installed
- For SCSS changes, use `bin/build-css.sh`; you'll need `sass` installed
- For a developer environment, run `bin/workspace.sh`, which will watch all files and hot-reload (iTerm2 required)

For macOS, you can install all dependencies with `brew install md5sha1sum closure-compiler typescript sass/sass/sass`
