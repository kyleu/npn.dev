go-embed -input web/assets -output app/assets/assets.go

md build
md build\release

md build\release\amd64
set /A GOARCH=amd64
go build -o build/release/amd64 github.com/kyleu/npn

md build\release\386
set /A GOARCH=386
go build -o build/release/386 github.com/kyleu/npn

md build\release\arm
set /A GOARCH=arm
go build -o build/release/arm github.com/kyleu/npn

git checkout app/assets/assets.go
