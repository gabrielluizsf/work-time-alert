package main

import (
	"embed"
	"io/fs"
	"log"
	"os"

	"github.com/gabrielluizsf/work-time-alert/server/logger"
	"github.com/gabrielluizsf/work-time-alert/server/session"

	"github.com/i9si-sistemas/nine"
	i9 "github.com/i9si-sistemas/nine/pkg/server"
)

//go:embed public/*
var publicFiles embed.FS

func main() {
	logger := logger.New()
	server := nine.NewServer(os.Getenv("PORT"))
	fs, _ := fs.Sub(publicFiles, "public")
	server.ServeFilesWithFS("/", fs)
	server.Use(func(c *i9.Context) error {
		logger.Request(c)
		return nil
	})
	server.Get("/hello", func(c *i9.Context) error {
		return c.Send([]byte("Hello World!"))
	})
	session.Routes(server, logger)
	log.Fatal(server.Listen())
}
