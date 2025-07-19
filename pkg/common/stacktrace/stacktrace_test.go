package stacktrace_test

import (
	"app/pkg/common/stacktrace"
	"errors"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var (
		ErrFailure = errors.New("failure")
		ErrUnknown = errors.New("unknown")
	)

	var err error
	if err = testMain(); err != nil {
		err = stacktrace.At(ErrFailure, ErrUnknown, err)
	}
	fmt.Println(err.Error())

	_err, ok := err.(stacktrace.Error)
	if !ok {
		t.Fatal("")
	}

	s := _err.StackTraceJSON()
	fmt.Println(s)

	s = _err.StackTrace()
	fmt.Println(s)
}

func testMain() error {
	if err := testSub(); err != nil {
		return stacktrace.At(fmt.Errorf("Main.error"), err)
	}
	return nil
}

func testSub() error {
	// do something error
	proc := func() error {
		return stacktrace.At(fmt.Errorf("Sub.error"))
	}
	err := proc()
	return stacktrace.At(err)
}
