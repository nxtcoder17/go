package errors

import (
	"errors"
	"fmt"
	"github.com/yext/yerrors"
)

func NewEf(err error, msg string, a ...interface{}) error {
	return yerrors.WrapFrame(yerrors.Errorf("%s while %w", fmt.Sprintf(msg, a...), err), 1)
}

func Newf(msg string, a ...interface{}) error {
	if len(a) > 0 {
		return yerrors.Wrap(yerrors.Errorf(msg, a))
	}
	return yerrors.New(msg)
}

func NewE(err error) error {
	return yerrors.Wrap(err)
}

func Is(err error, err2 error) bool {
	return errors.Is(err, err2)
}
