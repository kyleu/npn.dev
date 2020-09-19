SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "Building [ios]..."

mkdir -p build/ios/

go-embed -input web/assets -output app/assets/assets.go
echo "gomobile..."
gomobile bind -o build/ios/NpnServer.framework -target=ios github.com/kyleu/npn/lib
git checkout app/assets/assets.go
