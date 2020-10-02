# Scripts

There are scripts in the `./bin` directory to help you build, run, test, and publish npn

They're designed for macOS, but should work on Linux or Windows (via WSL).

- `asset-embed.sh`: Embeds assets for building into the project
- `asset-reset.sh`: Resets the assets to load from local filesystem in development
- `bootstrap.sh`: Downloads and installs the Go libraries and tools needed in other scripts
- `build.bat`: Builds the app on Windows (or just use make build)
- `build.sh`: Builds the app (or just use make build)
- `build-all.sh`: Attempts to build for all available platforms and architectures; Requires docker and a bunch of other stuff
- `build-android.sh`: Builds the project as an android framework and builds the native app in `projects/android`
- `build-client.sh`: Uses `scss` to compile the TypeScript files in `./client`
- `build-client-watch.sh`: Builds the TypeScript resources using `build-client`, then watches for changes in `./client`
- `build-css.sh`: Uses `scss` to compile the stylesheets in `./web/stylesheets`
- `build-css-watch.sh`: Builds the css resources using `build-css`, then watches for changes in `./web/stylesheets`
- `build-desktop.sh`: Builds all the desktop apps, XCode required
- `build-docker.sh`: Makes a release build, builds a docker image, then exports and zips the output
- `build-ios.sh`: Builds the project as an iOS framework and builds the native app in `projects/ios`
- `build-linux.sh`: Makes a release build for 64-bit Linux
- `build-macos.sh`: Builds the project as a macOS server and builds the native app in `projects/macos`
- `build-wasm.sh`: Builds the app as a WASM server
- `check.sh`: Runs code statistics, checks for outdated dependencies, then runs various linters
- `dev.sh`: Watches the project directories, and runs the main application, restarting when changes are detected
- `doc.sh`: Runs godoc for all projects, linking between projects and using custom logos and styling
- `format.sh`: Runs `gofmt` on all projects
- `package.sh`: Packages the build output for Github Releases
- `run-docker.sh`: Runs the Docker image produced by `build-docker`, exposing an HTTP port
- `run-release.sh`: Builds the project in release mode and runs it
- `tag.sh`: Updates the go.mod version, deletes go.sum, tags the git repo
- `templates.sh`: Builds hero templates, writing output to `./gen`
- `workspace.sh`: Opens all build in separate panes (iTerm2 only)
