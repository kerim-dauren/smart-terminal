package loggerx

const (
	XRequestIDKey string = "X-Request-Id"
	XTraceIDKey   string = "X-Trace-Id"
)

var standardHeaderKeys = []string{XRequestIDKey, XTraceIDKey}

type Env string

const (
	Dev  Env = "dev"
	Test Env = "test"
	Prod Env = "prod"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)

type Handler string

const (
	Text Handler = "text"
	JSON Handler = "json"
)
