package constants

import "errors"

var (
	ErrDB = errors.New("database error")
)

func WrapAsErrDB(err error) error {
	return errors.Join(ErrDB, err)
}

var (
	ErrHabitNotFound = errors.New("habit not found")
	ErrEventNotFound = errors.New("event not found")
)
