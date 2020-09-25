SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

echo "Building [android]..."

mkdir -p build/android/

bin/asset-embed.sh
echo "gomobile..."
gomobile bind -o build/android/npn.aar -target=android github.com/kyleu/npn/lib
bin/asset-reset.sh

cd  projects/android/npn/app/libs
rm -f npn.aar npn-sources.jar
cp ../../../../../build/android/npn.aar .
cp ../../../../../build/android/npn-sources.jar .

cd "${DIR}/projects/android/npn"

gradle assembleDebug

cp "app/build/outputs/apk/debug/app-debug.apk" "${DIR}/build/android/npn.apk"

cd "${DIR}/build/android"
