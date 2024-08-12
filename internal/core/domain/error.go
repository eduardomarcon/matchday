package domain

import (
	"errors"
)

var (
	ErrInternal        = errors.New("internal error")
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")
)
