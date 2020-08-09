SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

arch=amd64
os=windows

echo "Building [$os $arch]..."
env GOOS=$os GOARCH=$arch make build-release
mkdir -p ./build/$os/$arch
mv ./build/release/npn.exe ./build/$os/$arch/npn.exe
