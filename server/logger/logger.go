package logger

import "log"

type Request interface {
	Method() string
	Path() string
	IPs() []string
}

type logger struct {}

func (l logger) Request(req Request, data ...any) {
	logRequest(req, data...)
}

func (l logger) Data(data ...any){
	log.Printf("data=%v", data)
}

func New() logger {
	return logger{}
}

func logRequest(req Request, data ...any) {
	if len(data) > 0 {
		log.Printf("method=[%s] path=[%s] ips=%s data=%v", req.Method(), req.Path(), req.IPs(), data)
		return
	}
	log.Printf("method=[%s] path=[%s] ips=%s", req.Method(), req.Path(), req.IPs())
}
