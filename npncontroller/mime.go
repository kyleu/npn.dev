package npncontroller

import "mime"

// Common MIME types
func InitMime() {
	_ = mime.AddExtensionType(".ico", "image/x-icon")
	_ = mime.AddExtensionType(".eot", "font/eot")
	_ = mime.AddExtensionType(".tff", "font/tff")
	_ = mime.AddExtensionType(".woff", "application/font-woff")
	_ = mime.AddExtensionType(".woff2", "application/font-woff")
}
