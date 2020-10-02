package npndatabase

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
)

func (s *Service) Insert(q string, tx *sqlx.Tx, values ...interface{}) error {
	if s.debug {
		logQuery(s, "inserting row", q, values)
	}
	aff, err := s.execUnknown(q, tx, values...)
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New(fmt.Sprintf("no rows affected by insert using sql [%v] and %v values", q, len(values)))
	}
	return nil
}

func (s *Service) Update(q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	return s.process("updating", "updated", q, tx, expected, values...)
}

func (s *Service) UpdateOne(q string, tx *sqlx.Tx, values ...interface{}) error {
	_, err := s.Update(q, tx, 1, values...)
	return err
}

func (s *Service) Delete(q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	return s.process("deleting", "deleted", q, tx, expected, values...)
}

func (s *Service) DeleteOne(q string, tx *sqlx.Tx, values ...interface{}) error {
	_, err := s.Delete(q, tx, 1, values...)
	if err != nil {
		return errors.Wrap(err, errMessage("delete", q, values))
	}
	return err
}

func (s *Service) Exec(q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	return s.process("executing", "executed", q, tx, expected, values...)
}

func (s *Service) execUnknown(q string, tx *sqlx.Tx, values ...interface{}) (int, error) {
	var err error
	var ret sql.Result
	if tx == nil {
		r, e := s.db.Exec(q, values...)
		ret = r
		err = e
	} else {
		r, e := tx.Exec(q, values...)
		ret = r
		err = e
	}
	if err != nil {
		return 0, errors.Wrap(err, errMessage("exec", q, values))
	}
	aff, err := ret.RowsAffected()
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return int(aff), nil
}

func (s *Service) process(key string, past string, q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	if s.debug {
		logQuery(s, fmt.Sprintf("%v [%v] rows", key, expected), q, values)
	}

	aff, err := s.execUnknown(q, tx, values...)
	if err != nil {
		return 0, errors.Wrap(err, errMessage(past, q, values))
	}
	if expected > -1 && aff != expected {
		const msg = "expected [%v] %v row(s), but [%v] records affected from sql [%v] with values [%s]"
		return aff, errors.New(fmt.Sprintf(msg, expected, past, aff, q, npncore.ValueStrings(values)))
	}
	return aff, nil
}
