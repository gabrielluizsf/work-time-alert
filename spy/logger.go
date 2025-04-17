package spy

import "github.com/gabrielluizsf/work-time-alert/server/logger"

type Logger struct {
	request logger.Request
	data []any
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Request(req logger.Request, data ...any) {
	l.request = req
	l.data = data
}

func (l *Logger) Data(data ...any) {
	l.data = data
}