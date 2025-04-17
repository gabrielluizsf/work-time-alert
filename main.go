package main

import (
	"log"
	"os"

	"github.com/i9si-sistemas/nine"
	i9 "github.com/i9si-sistemas/nine/pkg/server"
)

func main() {
	server := nine.NewServer(os.Getenv("PORT"))
	server.ServeFiles("/", "./public")
	server.Use(func(c *i9.Context) error {
		log.Printf("method=[%s] path=[%s] ips=%s", c.Method(), c.Path(), c.IPs())
		return nil
	})
	log.Fatal(server.Listen())
}
