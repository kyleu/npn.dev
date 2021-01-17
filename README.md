# npn

[npn](https://npn.dev) helps you explore and test HTTP APIs, with a focus on speed and correctness.

It's basically Postman, but much smaller (8MB download) and faster. You can run npn as an HTTP server, or use a native desktop or mobile app.

## Download

https://npn.dev/download

## Source code

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

For macOS, you can install all dependencies with `brew install md5sha1sum sass/sass/sass`
