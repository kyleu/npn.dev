REM https://git-scm.com/download/win
REM https://golang.org/doc/install
REM http://tdm-gcc.tdragon.net/download

hero -extensions .html -source web/templates -pkgname templates -dest gen/templates
hero -extensions .sql -source query/sql -pkgname query -dest gen/query

go-embed -input web/assets -output app/controllers/assets/assets.go

go build -o build/ ./cmd/...
