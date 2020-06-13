SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

./bin/build-client.sh
./bin/build-css.sh

for arch in amd64 386
do
  for os in darwin linux windows
  do
    echo "Building [$os $arch]..."
    env GOOS=$os GOARCH=$arch make build-release
    mkdir -p ./build/$os/$arch
    if [ "$os" = "windows" ]; then
      mv ./build/release/npn.exe ./build/$os/$arch/npn.exe
    else
      mv ./build/release/npn ./build/$os/$arch/npn
    fi
  done
done
