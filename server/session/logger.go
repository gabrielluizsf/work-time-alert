package session

type Logger interface {
	Data(data ...any)
}