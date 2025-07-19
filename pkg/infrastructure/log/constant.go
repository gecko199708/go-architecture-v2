package log

type Level int

const (
	DebugLevel Level = 4*iota - 4
	InfoLevel
	WarnLevel
	ErrorLevel

	Default Level = InfoLevel
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return "unknown"
	}
}
