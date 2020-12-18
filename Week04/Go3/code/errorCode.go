package code

import "errors"

var (
	NotFound   error = errors.New("NotFound")
	SqlQuery   error = errors.New("SqlQuery")
	ParseError error = errors.New("ParseError")

	SqlPrepare    error = errors.New("SqlPrepare")
	SqlExec    error = errors.New("SqlExec")
)
