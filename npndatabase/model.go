package npndatabase

import (
	"time"
)

// A database migration
type Migration struct {
	Idx     int       `json:"idx"`
	Title   string    `json:"title"`
	Src     string    `json:"src"`
	Created time.Time `json:"created"`
}

// Array helper
type Migrations []*Migration

type migrationDTO struct {
	Idx     int       `db:"idx"`
	Title   string    `db:"title"`
	Src     string    `db:"src"`
	Created time.Time `db:"created"`
}

func (dto *migrationDTO) toMigration() *Migration {
	return &Migration{
		Idx:     dto.Idx,
		Title:   dto.Title,
		Src:     dto.Src,
		Created: dto.Created,
	}
}

// Count type for queries, contains a single int64 field
type Count struct {
	C int64 `db:"c"`
}

// Helper class for Index definitions
type Index struct {
	Name       string `db:"n"`
	Definition string `db:"d"`
}
