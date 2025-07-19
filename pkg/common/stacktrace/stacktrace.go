package stacktrace

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"time"
)

type Error interface {
	StackTrace() string
	StackTraceJSON() string
}

type location struct {
	Fn   string `json:"function"`
	File string `json:"file"`
	Line int    `json:"line"`
}

func (l location) String() string {
	return fmt.Sprintf("%s:%d", l.Fn, l.Line)
}

func CallLocation(offset int) location {
	pc, file, line, _ := runtime.Caller(offset)
	fn := runtime.FuncForPC(pc).Name()
	return location{fn, file, line}
}

type callPoint struct {
	Message []string `json:"message"`
	// Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Location  string    `json:"location"`
}

type structuredError struct {
	items []*callPoint
}

func At(err ...error) error {
	var _root *structuredError
	var items []error
	for _, _err := range err {
		if _err, ok := _err.(*structuredError); ok {
			_root = _err
			continue
		}
		items = append(items, _err)
	}
	if _root == nil {
		_root = new(structuredError)
	}
	if len(items) > 1 {
		_root.add(errors.Join(items...))
	} else if len(items) == 1 {
		_root.add(items[0])
	} else {
		_root.add(nil)
	}

	return _root
}

func (e *structuredError) add(err error) {
	const (
		skip = 3
	)
	// t := reflect.ValueOf(err)
	var message []string
	if err != nil {
		if _err, ok := err.(interface{ Unwrap() []error }); ok {
			for _, _err := range _err.Unwrap() {
				message = append(message, _err.Error())
			}
		} else {
			message = []string{err.Error()}
		}
	}

	e.items = append([]*callPoint{
		{
			Message: message,
			// Type:      t.String(),
			Timestamp: time.Now().Local(),
			Location:  CallLocation(skip).String(),
		},
	}, e.items...)
}

func (e *structuredError) Error() string {
	return fmt.Sprintf("struct error. (%d)", len(e.items))
}

func (e *structuredError) StackTrace() string {
	const (
		offset = 2
	)
	if len(e.items) == 0 {
		return ""
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s (%v) \n", e.items[0].Message, e.items[0].Timestamp)
	fmt.Fprintf(&buf, "stacktrace: %s\n", CallLocation(offset).String())
	for _, err := range e.items {
		fmt.Fprintf(&buf, "  at %s\n", err.Location)
	}

	return buf.String()
}

func (e *structuredError) StackTraceJSON() string {
	const (
		offset = 2
	)
	if len(e.items) == 0 {
		return ""
	}
	m := map[string]any{
		"stacktrace": CallLocation(offset).String(),
		"errors":     e.items,
	}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(m)
	return buf.String()
}

func (e *structuredError) MarshalJSON() ([]byte, error) {
	const (
		offset = 2
	)
	m := map[string]any{
		"stacktrace": CallLocation(offset).String(),
		"errors":     e.items,
	}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(m)
	return buf.Bytes(), nil
}
