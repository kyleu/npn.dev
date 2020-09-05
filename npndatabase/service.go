package npndatabase

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	debug  bool
	db     *sqlx.DB
	logger logur.Logger
}

func NewService(debug bool, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "db"})
	return &Service{debug: debug, db: db, logger: logger}
}

func (s *Service) StartTransaction() (*sqlx.Tx, error) {
	if s.debug {
		s.logger.Debug("opening transaction")
	}
	return s.db.Beginx()
}

func errMessage(t string, q string, values []interface{}) string {
	return fmt.Sprintf("error running %v sql [%v] with values [%v]", t, strings.TrimSpace(q), npncore.ValueStrings(values))
}

func logQuery(s *Service, msg string, q string, values []interface{}) {
	s.logger.Debug(fmt.Sprintf("%v {\n  SQL: %v\n  Values: %v\n}", msg, strings.TrimSpace(q), npncore.ValueStrings(values)))
}

func (s *Service) Tables() ([]string, error) {
	type table struct {
		Name string `db:"n"`
	}
	res := []*table{}
	sql := "select tablename as n from pg_catalog.pg_tables where schemaname != 'pg_catalog' and schemaname != 'information_schema'"
	err := s.db.Select(&res, sql)

	ret := make([]string, 0, len(res))
	for _, t := range res {
		ret = append(ret, t.Name)
	}
	sort.Strings(ret)
	return ret, err
}

func (s *Service) Indexes(tableName string) ([]*Index, error) {
	ret := []*Index{}
	sql := "select indexname n, indexdef d from pg_indexes where schemaname = 'public' and tablename = $1 order by indexname"
	err := s.db.Select(&ret, sql, tableName)
	return ret, err
}

func ArrayToString(a []string) string {
	return "{" + strings.Join(a, ",") + "}"
}
