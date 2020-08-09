package npncore

var AppKey = "SetAppKey"
var AppName = "SetAppName"

type key int

const (
	ContextKey key = iota
	RoutesKey  key = iota
	InfoKey    key = iota
)
