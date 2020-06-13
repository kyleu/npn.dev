# Scripts

There are scripts in the `./bin` directory to help you build, run, test, and publish npn

They're designed for macOS, but should work on Linux or Windows (via WSL).

- `bootstrap.sh`: Downloads and installs the Go libraries and tools needed in other scripts
- `build-client.sh`: Uses `scss` to compile the TypeScript files in `./client`
- `build-client-watch.sh`: Builds the TypeScript resources using `build-client`, then watches for changes in `./client`
- `build-css.sh`: Uses `scss` to compile the stylesheets in `./web/stylesheets`
- `build-css-watch.sh`: Builds the css resources using `build-css`, then watches for changes in `./web/stylesheets`
- `build-docker.sh`: Makes a release build, builds a docker image, then exports and zips the output
- `build-linux.sh`: Makes a release build for 64-bit Linux
- `build-windows.bat`: Makes a release build, builds a docker image, then exports and zips the output
- `check.sh`: Runs code statistics, checks for outdated dependencies, then runs various linters
- `dev.sh`: Watches the project directories, and runs the main application, restarting when changes are detected
- `doc.sh`: Runs godoc for all projects, linking between projects and using custom logos and styling
- `format.sh`: Runs `gofmt` on all projects
- `package.sh`: Runs the code formatter, checks all the projects, then builds binaries for Linux, macOS, and Windows
- `run-docker.sh`: Runs the Docker image produced by `build-docker`, exposing an HTTP port
- `run-release.sh`: Builds the project in release mode and runs it
- `templates.sh`: Builds hero templates, writing output to `./gen`
- `workspace.sh`: Opens all build in separate panes (iTerm2 only)
