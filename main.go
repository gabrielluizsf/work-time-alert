package main

import (
	"log"

	"github.com/i9si-sistemas/nine"
)

func main() {
	server := nine.NewServer("")
	server.ServeFiles("/", "./public")
	server.Use(func(c *nine.Context) error {
		log.Printf("method=[%s] path=[%s] ips=%s", c.Method(), c.Path(), c.IPs())
		return nil
	})
	log.Fatal(server.Listen())
}
