package db

import "errors"

var (
	NotFindErr = errors.New("record not find")
	HaveExistedErr = errors.New("record have existed")
)

