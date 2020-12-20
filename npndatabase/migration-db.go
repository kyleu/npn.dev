package npndatabase

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
)

// Lists the available Migrations defined in DatabaseMigrations
func (s *Service) ListMigrations(params *npncore.Params) Migrations {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyMigration, params, npncore.DefaultCreatedOrdering...)

	var dtos []migrationDTO
	q := SQLSelect("*", npncore.KeyMigration, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.Select(&dtos, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving migrations: %+v", err))
		return nil
	}

	return toMigrations(dtos)
}

// Gets a migration at the provided index
func (s *Service) GetMigrationByIdx(idx int) *Migration {
	var dto = &migrationDTO{}
	q := SQLSelectSimple("*", "migration", "idx = $1")
	err := s.Get(dto, q, nil, idx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting migration by idx [%v]: %+v", idx, err))
		return nil
	}
	return dto.toMigration()
}

// Removes a migration for some reason
func (s *Service) RemoveMigrationByIdx(idx int) error {
	q := SQLDelete("migration", "idx = $1")
	_, err := s.Delete(q, nil, 1, idx)
	if err != nil {
		return errors.Wrap(err, "error removing migration")
	}
	return nil
}

func newMigration(s *Service, e Migration) error {
	q := SQLInsert("migration", []string{"idx", "title", "src"}, 1)
	return s.Insert(q, nil, e.Idx, e.Title, e.Src)
}

func maxMigrationIdx(s *Service) int {
	q := SQLSelectSimple("max(idx) as x", "migration")
	max, err := s.SingleInt(q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting migrations: %+v", err))
		return -1
	}
	return int(max)
}

func toMigrations(dtos []migrationDTO) Migrations {
	ret := make(Migrations, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.toMigration())
	}

	return ret
}
