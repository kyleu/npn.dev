package cli

import (
	"github.com/kirsle/configdir"
	"github.com/kyleu/npn/npncore"
)

func defaultDirectory() string {
	dir := configdir.LocalConfig(npncore.AppName)
	_ = configdir.MakePath(dir) // Ensure it exists.
	return dir
}
