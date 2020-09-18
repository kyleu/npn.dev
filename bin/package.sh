SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

set -e

# ./bin/build-all.sh

mkdir -p ./build/release/stage

rm -rf ./build/publish
mkdir -p ./build/publish

cd build/release/stage

# cp -R "$DIR/data" ./data

pkg () {
  echo "$4 ($2)..."
  cp "$DIR/build/$1/$2/$3" "./$3"

  if [ $2 = "amd64" ]; then
    zip -r "$DIR/build/publish/npn.$4.zip" *
  else
    zip -r "$DIR/build/publish/npn.$4.$2.zip" *
  fi

  rm "./$3"
}

# macOS
pkg darwin amd64 npn macos

# Linux
pkg linux amd64 npn linux
pkg linux 386 npn linux
pkg linux arm64 npn linux
pkg linux arm npn linux
pkg linux mips npn linux
pkg linux riscv64 npn linux

# FreeBSD
pkg freebsd amd64 npn freebsd
pkg freebsd arm64 npn freebsd
pkg freebsd 386 npn freebsd

# Windows
pkg windows amd64 npn.exe windows
pkg windows 386 npn.exe windows
pkg windows arm npn.exe windows

# Docker
cp "$DIR/build/docker/npn.docker.tar.gz" "$DIR/build/publish/npn.docker.tar.gz"

# WASM
cp "$DIR/build/js/wasm/npn.wasm" ./npn.wasm
zip -r "$DIR/build/publish/npn.wasm.zip" *
rm ./npn.wasm

# Android
cp "$DIR/build/android/npn.aar" ./npn.aar
zip -r "$DIR/build/publish/npn.android.zip" *
rm ./npn.aar

# iOS
cp  -r "$DIR/build/ios/npn.framework" ./npn.framework
zip -r "$DIR/build/publish/npn.ios.zip" *
rm  -rf ./npn.framework

rm -rf "$DIR/build/release/stage"
