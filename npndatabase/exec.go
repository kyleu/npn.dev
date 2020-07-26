package npndatabase

import (
	"database/sql"
	"fmt"

	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
	"github.com/fevo-tech/charybdis/app/util"
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
	if s.debug {
		logQuery(s, fmt.Sprintf("updating [%v] rows", expected), q, values)
	}

	aff, err := s.execUnknown(q, tx, values...)
	if err != nil {
		return 0, errors.Wrap(err, errMessage("update", q, values))
	}
	if expected > -1 && aff != expected {
		msg := "expected [%v] updated row(s), but [%v] records affected from sql [%v] with values [%s]"
		return aff, errors.New(fmt.Sprintf(msg, expected, aff, q, util.ValueStrings(values)))
	}
	return aff, nil
}

func (s *Service) UpdateOne(q string, tx *sqlx.Tx, values ...interface{}) error {
	_, err := s.Update(q, tx, 1, values...)
	return err
}

func (s *Service) Delete(q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	if s.debug {
		logQuery(s, fmt.Sprintf("deleting [%v] rows", expected), q, values)
	}
	aff, err := s.execUnknown(q, tx, values...)
	if err != nil {
		return 0, errors.Wrap(err, errMessage("delete", q, values))
	}
	if expected > -1 && aff != expected {
		msg := "expected [%v] deleted row(s), but [%v] records affected from sql [%v] with values [%s]"
		return aff, errors.New(fmt.Sprintf(msg, expected, aff, q, util.ValueStrings(values)))
	}
	return aff, err
}

func (s *Service) DeleteOne(q string, tx *sqlx.Tx, values ...interface{}) error {
	_, err := s.Delete(q, tx, 1, values...)
	if err != nil {
		return errors.Wrap(err, errMessage("delete", q, values))
	}
	return err
}

func (s *Service) Exec(q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	if s.debug {
		logQuery(s, fmt.Sprintf("executing [%v] rows", expected), q, values)
	}
	aff, err := s.execUnknown(q, tx, values...)
	if err != nil {
		return 0, errors.Wrap(err, errMessage("exec", q, values))
	}
	if expected > -1 && aff != expected {
		msg := "expected [%v] exec row(s), but [%v] records affected from sql [%v] with values [%s]"
		return aff, errors.New(fmt.Sprintf(msg, expected, aff, q, util.ValueStrings(values)))
	}
	return aff, nil
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
