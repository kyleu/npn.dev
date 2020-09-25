package npncore

var (
	AppKey = "SetAppKey"
  AppName = "SetAppName"
  AppPlatform = "SetAppName"
  AppVersion = "0"
)

type key int

const (
	ContextKey key = iota
	RoutesKey  key = iota
	InfoKey    key = iota
)
