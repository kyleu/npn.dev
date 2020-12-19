package npndatabase

import (
	"fmt"
	"strings"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"golang.org/x/text/language"
	"logur.dev/logur"
)

type MigrationFile struct {
	Title string
	F     func(*strings.Builder)
}

type MigrationFiles []*MigrationFile

var InitialSchemaMigrations = MigrationFiles{}

var DatabaseMigrations = MigrationFiles{}

func exec(file *MigrationFile, s *Service, logger logur.Logger) (string, error) {
	sb := &strings.Builder{}
	file.F(sb)
	sql := sb.String()
	sqls := strings.Split(sql, ";")
	startNanos := npncore.TimerStart()
	for _, q := range sqls {
		if len(strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(q), "--"))) > 0 {
			_, err := s.Exec(q, nil, -1)
			if err != nil {
				return "", errors.Wrap(err, "cannot execute ["+file.Title+"]")
			}
		}
	}
	ms := npncore.MicrosToMillis(language.AmericanEnglish, npncore.TimerEnd(startNanos))
	logger.Debug(fmt.Sprintf("ran query [%s] in [%v]", file.Title, ms))
	return sql, nil
}
