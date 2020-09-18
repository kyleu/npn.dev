SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "Building [ios]..."

mkdir -p build/ios/
gomobile bind -o build/ios/npn.framework -target=ios github.com/kyleu/npn/lib
