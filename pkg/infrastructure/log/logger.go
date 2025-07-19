package log

import (
	"app/pkg/common/log"
	"app/pkg/common/stacktrace"
	"encoding/json"
	"io"
	"os"
	"time"
)

type Params = log.Params

type Logger struct {
	w  io.Writer
	lv Level
}

const defaultTimeFormat = "2006-01-02 15:04:05"

const callLocationOffset = 3

func (l *Logger) Debug(params Params) {
	l.print(DebugLevel, params)
}
func (l *Logger) Info(params Params) {
	l.print(InfoLevel, params)
}
func (l *Logger) Warn(params Params) {
	l.print(WarnLevel, params)
}
func (l *Logger) Error(params Params) {
	l.print(ErrorLevel, params)
}

func (l *Logger) print(level Level, params Params) {
	if level < l.lv {
		return
	}
	var (
		timestamp = time.Now().Local()
		location  = stacktrace.CallLocation(callLocationOffset)
	)
	_params := Params{
		"level":     level.String(),
		"timestamp": timestamp.Format(defaultTimeFormat),
		"location":  location.String(),
		"params":    params,
	}
	encode := json.NewEncoder(l.output()).Encode

	encode(_params)
}

func (l *Logger) output() io.Writer {
	if l.w == nil {
		l.w = os.Stdout
	}
	return l.w
}

func (l *Logger) SetOutput(w io.Writer) {
	l.w = w
}

func (l *Logger) SetLevel(level Level) {
	l.lv = level
}
