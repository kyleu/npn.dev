package npncore

var (
	AppKey      = "SetAppKey"
	AppName     = "SetAppName"
	AppPlatform = "SetAppName"
	AppVersion  = "0"
)

type key int

const (
	RoutesKey key = iota + 1
	InfoKey
	ContextKey
)
