package npndatabase

import (
	"fmt"
	"strings"
	"time"

	"emperror.dev/errors"
	"logur.dev/logur"
)

func DBWipe(s *Service, logger logur.Logger) error {
	for _, file := range initialSchemaMigrations {
		_, err := exec(file, s, logger)
		if err != nil {
			return err
		}
	}
	return nil
}

func Migrate(s *Service) error {
	var err error

	maxIdx := maxMigrationIdx(s)
	// s.logger.Info(fmt.Sprintf("migrating database schema: %v", maxIdx))

	for i, file := range databaseMigrations {
		idx := i + 1
		switch {
		case idx == maxIdx:
			m := s.GetMigrationByIdx(maxIdx)
			if m == nil {
				continue
			}
			if m.Title != file.Title {
				s.logger.Info(fmt.Sprintf("migration [%v] name has changed from [%v] to [%v]", idx, m.Title, file.Title))
				err = s.RemoveMigrationByIdx(idx)
				if err != nil {
					return err
				}
				err = applyMigration(s, idx, file)
				continue
			}
			sb := &strings.Builder{}
			file.F(sb)
			nc := sb.String()
			if nc != m.Src {
				s.logger.Info(fmt.Sprintf("migration [%v:%v] content has changed from [%vB] to [%vB]", idx, file.Title, len(nc), len(m.Src)))
				err = s.RemoveMigrationByIdx(idx)
				if err != nil {
					return err
				}
				err = applyMigration(s, idx, file)
			}
		case idx > maxIdx:
			err = applyMigration(s, idx, file)
		default:
			// noop
		}
	}

	return errors.Wrap(err, "error running database migration")
}

func applyMigration(s *Service, idx int, file migrationFile) error {
	s.logger.Info(fmt.Sprintf("applying database migration [%v]: %v", idx, file.Title))
	sql, err := exec(file, s, s.logger)
	if err != nil {
		return err
	}
	return newMigration(s, Migration{
		Idx:     idx,
		Title:   file.Title,
		Src:     sql,
		Created: time.Time{},
	})
}
