package npncore

var AppKey = "SetAppKey"
var AppName = "SetAppName"
var AppPlatform = "SetAppName"

type key int

const (
	ContextKey key = iota
	RoutesKey  key = iota
	InfoKey    key = iota
)
