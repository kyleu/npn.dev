SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "Building [android]..."

mkdir -p build/android/
gomobile bind -o build/android/npn.aar -target=android github.com/kyleu/npn/lib
