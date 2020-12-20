package npndatabase

import (
	"fmt"
	"strings"
	"time"

	"emperror.dev/errors"
	"logur.dev/logur"
)

// Wipes the database, probably
func DBWipe(s *Service, logger logur.Logger) error {
	for _, file := range InitialSchemaMigrations {
		_, err := exec(file, s, logger)
		if err != nil {
			return err
		}
	}
	return nil
}

// Runs database SQL migrations provided in DatabaseMigrations, if needed
func Migrate(s *Service) error {
	var err error

	maxIdx := maxMigrationIdx(s)

	if len(DatabaseMigrations) > maxIdx+1 {
		s.logger.Info(fmt.Sprintf("applying [%v] database migrations...", len(DatabaseMigrations)-maxIdx+1))
	}

	for i, file := range DatabaseMigrations {
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

// Applies a single migration
func applyMigration(s *Service, idx int, file *MigrationFile) error {
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
