SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

set -e

# ./bin/build-all.sh

mkdir -p ./build/stage

rm -rf ./build/package
mkdir -p ./build/package

cd build/stage

# cp -R "$DIR/data" ./data

pkg () {
  echo "$4 ($2)..."
  cp "$DIR/build/$1/$2/$3" "./$3"

  if [ $2 = "amd64" ]; then
    zip -r "$DIR/build/package/npn.$4.zip" *
  else
    zip -r "$DIR/build/package/npn.$4.$2.zip" *
  fi

  rm "./$3"
}

# macOS
pkg darwin amd64 npn macos

echo "macOS app..."
cd ../darwin
zip -r "npn.macos.zip" npn.app
mv "npn.macos.zip" "../package"
cd ../stage

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
echo "docker..."
cp "$DIR/build/docker/npn.docker.tar.gz" "$DIR/build/package/npn.docker.tar.gz"

# WASM
echo "wasm..."
cp "$DIR/build/js/wasm/npn.wasm" ./npn.wasm
zip -r "$DIR/build/package/npn.wasm.zip" *
rm ./npn.wasm

# Android
echo "android library..."
cp "$DIR/build/android/npn.aar" ./npn.aar
zip -r "$DIR/build/package/npn.android.aar.zip" npn.aar
rm ./npn.aar

echo "android app..."
cp "$DIR/build/android/npn.apk" ./npn.apk
zip -r "$DIR/build/package/npn.android.apk.zip" npn.apk
rm ./npn.apk

# iOS
echo "ios framework..."
cp  -r "$DIR/build/ios/NpnServer.framework" ./NpnServer.framework

cd NpnServer.framework
rm -rf Headers
rm -rf Modules
rm -rf NpnServer
rm -rf Resources
cp -R Versions/A/* .
rm -rf Versions
cd ..

zip -r "$DIR/build/package/npn.ios.framework.zip" *
rm  -rf ./NpnServer.framework

echo "ios app..."
cd ../ios
zip -r "npn.ios.app.zip" npn.app
mv "npn.ios.app.zip" "../package"
cd ../stage

rm -rf "$DIR/build/stage"
