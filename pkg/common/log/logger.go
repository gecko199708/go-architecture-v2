package log

type Params map[string]any

type Logger interface {
	Debug(params Params)
	Info(params Params)
	Warn(params Params)
	Error(params Params)
}
