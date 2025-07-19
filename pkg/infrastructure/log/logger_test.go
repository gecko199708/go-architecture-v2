package log_test

import (
	"app/pkg/common/stacktrace"
	"app/pkg/infrastructure/log"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	logger := new(log.Logger)

	logger.Debug(log.Params{
		"test": "test",
	})

	logger.Info(log.Params{
		"test": "test",
	})

	err := stacktrace.At(fmt.Errorf("err1"))
	stacktrace.At(err, fmt.Errorf("err2"))

	logger.Error(log.Params{
		"stacktrace": err,
	})
}
