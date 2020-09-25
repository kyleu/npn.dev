SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "Building [ios]..."

mkdir -p build/ios/

bin/asset-embed.sh
echo "gomobile..."
gomobile bind -o build/ios/NpnServer.framework -target=ios github.com/kyleu/npn/lib
bin/asset-reset.sh

echo "Building [ios] app..."

cd projects/ios/npn

xcodebuild -project npn.xcodeproj

cd build/Release-iphoneos/

cp -R npn.app ../../../../../build/ios
