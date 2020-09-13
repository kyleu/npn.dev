SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

# ./bin/build-all.sh

mkdir -p ./build/release/stage

rm -rf ./build/publish
mkdir -p ./build/publish

cd build/release/stage

# cp -R "$DIR/data" ./data

cp "$DIR/build/linux/amd64/npn" ./npn
zip -r "$DIR/build/publish/npn.linux.zip" *
rm ./npn

cp "$DIR/build/linux/arm64/npn" ./npn
zip -r "$DIR/build/publish/npn.linux.arm.zip" *
rm ./npn

cp "$DIR/build/darwin/amd64/npn" ./npn
zip -r "$DIR/build/publish/npn.macos.zip" *
rm ./npn

cp "$DIR/build/windows/amd64/npn.exe" ./npn.exe
zip -r "$DIR/build/publish/npn.windows.zip" *
rm ./npn.exe

cp "$DIR/build/docker/npn.docker.tar.gz" "$DIR/build/publish/npn.docker.tar.gz"

rm -rf "$DIR/build/release/stage"
