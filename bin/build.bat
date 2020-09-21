go-embed -input npnasset/vendor -output npnasset/assets/assets.go
go-embed -input web/assets -output app/assets/assets.go

md build
md build\release

go build -o build/release github.com/kyleu/npn/cmd/npn

git checkout app/assets/assets.go
