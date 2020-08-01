SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

./bin/build-macos.sh
./bin/build-linux.sh
./bin/build-windows.sh

mkdir -p ./build/release

# Linux
cd build/linux/amd64
zip -r ../../release/npn.linux.zip npn

cd ../arm64
zip -r ../../release/npn.linux.arm64.zip npn

# macOS
cd ../../../
cd build/darwin/amd64
zip -r ../../release/npn.macos.zip npn

cd ../arm64
zip -r ../../release/npn.macos.arm64.zip npn

# Windows
cd ../../../
cd build/windows/amd64
zip -r ../../release/npn.windows.zip npn.exe

cd ../arm64
zip -r ../../release/npn.windows.arm64.zip npn.exe


cd ../../../
