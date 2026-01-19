package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/gabrielluizsf/work-time-alert/browser"
	"github.com/gabrielluizsf/work-time-alert/server/logger"
	"github.com/gabrielluizsf/work-time-alert/server/session"

	"github.com/i9si-sistemas/nine"
	i9 "github.com/i9si-sistemas/nine/pkg/server"
)

//go:embed public/*
var publicFiles embed.FS

func main() {
	logger := logger.New()
	port := os.Getenv("PORT")
	server := nine.NewServer(port)
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
	go openBrowser(server)
	log.Fatal(server.Listen())
}

type Server interface {
	Port() string
}

func openBrowser(server Server) {
	time.Sleep(1 * time.Millisecond)
	browser.Open(fmt.Sprintf("http://localhost:%s", server.Port()))
}
