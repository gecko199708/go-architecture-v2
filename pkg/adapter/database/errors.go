package database

import "errors"

var (
	ErrUnSupportedConfig = errors.New("unsupported config")

	ErrFailureOpenDatabase = errors.New("failure open database")
)
