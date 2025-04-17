package main

import (
	"log"
	"os"
	"work-time-alert/server/logger"
	"work-time-alert/server/session"

	"github.com/i9si-sistemas/nine"
	i9 "github.com/i9si-sistemas/nine/pkg/server"
)

func main() {
	logger := logger.New()
	server := nine.NewServer(os.Getenv("PORT"))
	server.ServeFiles("/", "./public")
	server.Use(func(c *i9.Context) error {
		logger.Request(c)
		return nil
	})
	server.Get("/hello", func(c *i9.Context) error {
		return c.Send([]byte("Hello World!"))
	})
	session.Routes(server)
	log.Fatal(server.Listen())
}
