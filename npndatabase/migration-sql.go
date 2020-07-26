package npndatabase

import (
	"fmt"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/fevo-tech/charybdis/app/util"
	"github.com/fevo-tech/charybdis/gen/query"
	"golang.org/x/text/language"
	"logur.dev/logur"
)

type migrationFile struct {
	Title string
	F     func(*strings.Builder)
}

var initialSchemaMigrations = []migrationFile{
	{Title: "reset", F: func(sb *strings.Builder) { query.ResetDatabase(sb) }},
	{Title: "create-types", F: func(sb *strings.Builder) { query.CreateTypes(sb) }},
	{Title: "create-tables", F: func(sb *strings.Builder) { query.CreateTables(sb) }},
	{Title: "seed-data", F: func(sb *strings.Builder) { query.SeedData(sb) }},
}

var databaseMigrations = []migrationFile{
	{Title: "first-migration", F: func(sb *strings.Builder) { query.Migration1(sb) }},
}

func exec(file migrationFile, s *Service, logger logur.Logger) (string, error) {
	sb := &strings.Builder{}
	file.F(sb)
	sql := sb.String()
	sqls := strings.Split(sql, ";")
	startNanos := time.Now().UnixNano()
	for _, q := range sqls {
		if len(strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(q), "--"))) > 0 {
			_, err := s.Exec(q, nil, -1)
			if err != nil {
				return "", errors.Wrap(err, "cannot execute ["+file.Title+"]")
			}
		}
	}
	elapsed := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
	ms := util.MicrosToMillis(language.AmericanEnglish, int(elapsed))
	logger.Debug(fmt.Sprintf("ran query [%s] in [%v]", file.Title, ms))
	return sql, nil
}
