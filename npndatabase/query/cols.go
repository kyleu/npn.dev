package query

import "github.com/fevo-tech/charybdis/app/util"

var (
	allowedDataSortColumns = []string{util.KeyID, "namespace", "payload_type", "size", "filename", "error", "created"}
)

var allowedColumns = map[string][]string{
	util.KeyData: allowedDataSortColumns,
}
