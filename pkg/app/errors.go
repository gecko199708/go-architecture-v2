package app

import "errors"

var (
	ErrFailureServerStart       = errors.New("server start or runtime failure")
	ErrFailureCloseDependencies = errors.New("failure closing dependencies")
	ErrFailureGracefulShutdown  = errors.New("graceful shutdown failed")
)
