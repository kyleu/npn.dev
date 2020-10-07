package npncore

var (
	AppKey      = "SetAppKey"
	AppName     = "SetAppName"
	AppPlatform = "SetAppPlatform"
	AppVersion  = "0.0.0"
)

type key int

const (
	RoutesKey key = iota + 1
	InfoKey
	ContextKey
)
