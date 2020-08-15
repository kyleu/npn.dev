package npndatabase

import (
	"fmt"
	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	// load postgres driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"logur.dev/logur"
)

type DBParams struct {
	Username string
	Password string
	DBName   string
	Debug    bool
	Wipe     bool
	Migrate  bool
	Logger   logur.Logger
}

func OpenDatabase(params DBParams) (*Service, error) {
	params.Logger = logur.WithFields(params.Logger, map[string]interface{}{npncore.KeySvc: "database"})

	host := "localhost"
	port := 5432
	user := params.Username
	password := params.Password
	dbname := params.DBName

	template := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	url := fmt.Sprintf(template, host, port, user, password, dbname)

	db, err := sqlx.Open("pgx", url)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	svc := NewService(params.Debug, db, params.Logger)

	if params.Wipe {
		err = DBWipe(svc, params.Logger)
		if err != nil {
			return nil, errors.Wrap(err, "error applying initial schema")
		}
	}

	if params.Migrate {
		err = Migrate(svc)
		if err != nil {
			return nil, errors.Wrap(err, "error applying database migrations")
		}
	}

	return svc, nil
}
