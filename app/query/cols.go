package query

import "github.com/kyleu/npn/app/util"

var (
	allowedUserSortColumns       = []string{util.KeyID, util.KeyName, util.KeyRole, util.KeyTheme, "navColor", "linkColor", "picture", "locale", util.KeyCreated}
)

var allowedColumns = map[string][]string{
	util.KeyUser:         allowedUserSortColumns,
}
