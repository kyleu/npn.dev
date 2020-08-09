SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

./bin/build-client.sh
./bin/build-css.sh

./bin/build-macos.sh
./bin/build-linux.sh
./bin/build-linux-arm.sh
./bin/build-windows.sh

mkdir -p ./build/release/stage

cd build/release/stage

# cp -R "$DIR/data" ./data

cp "$DIR/build/linux/amd64/npn" ./npn
zip -r "$DIR/build/release/npn.linux.zip" *
rm ./npn

cp "$DIR/build/linux/arm64/npn" ./npn
zip -r "$DIR/build/release/npn.linux.arm.zip" *
rm ./npn

cp "$DIR/build/darwin/amd64/npn" ./npn
zip -r "$DIR/build/release/npn.macos.zip" *
rm ./npn

cp "$DIR/build/windows/amd64/npn.exe" ./npn.exe
zip -r "$DIR/build/release/npn.windows.zip" *
rm ./npn.exe

rm -rf "$DIR/build/release/stage"
