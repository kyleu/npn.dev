SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "Building [macos] app..."

mkdir -p build/darwin/

bin/build.sh darwin amd64
cp build/darwin/amd64/npn projects/macos/npn/npn/npn-server

cd projects/macos/npn

xcodebuild -project npn.xcodeproj

cd build/Release/

cp -R npn.app ../../../../../build/darwin
