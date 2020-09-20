SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "Building [ios]..."

mkdir -p build/macos/

bin/build.sh darwin amd64
cp build/darwin/amd64/npn projects/macos/npn/npn/npn-server

cd projects/macos/npn

xcodebuild -project npn.xcodeproj

cd build/Release/

zip -r ../../../../../build/macos/npn.macos.zip npn.app
