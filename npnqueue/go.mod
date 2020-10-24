module github.com/kyleu/npn/npnqueue

go 1.14

require (
	github.com/kyleu/npn/npncore v0.0.60 // fevgo
	github.com/Shopify/sarama v1.27.1
)

replace github.com/kyleu/npn/npncore => ../npncore
