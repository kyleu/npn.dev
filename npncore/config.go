package npncore

var (
	// The global key for this application. Must be lowercase with only alphanumeric characters
	AppKey = "SetAppKey"
	// The global name for this application. Go nuts
	AppName = "SetAppName"
	// The global platform for this application. Usually an empty string, sometimes "wasm"
	AppPlatform = "SetAppPlatform"
	// The global version for this application. Should probably be pulled from git
	AppVersion = "0.0.0"
	// When set to a domain, makes sure all OAuth logins originate from accounts in that domain
	RequireLoginsFrom = "none"
	// Extra scripts to include in the basic HTML layout template
	IncludedScripts = []string{}
	// Extra stylesheets to include in the basic HTML layout template
	IncludedStylesheets = []string{}
)

type key int

const (
	// Used as a session key for looking up the Router
	RoutesKey key = iota + 1
	// Used as a session key for looking up the AppInfo
	InfoKey
	// Used as a session key for looking up the RequestContext
	ContextKey
)
