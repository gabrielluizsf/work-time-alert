package main

import "slices"

const webFlag = "--web"

// Environment is a way of running the app: desktop window or web server.
type Environment interface {
	Start()
}

// environmentFor picks the web server when --web is passed, desktop otherwise.
func environmentFor(args []string) Environment {
	if slices.Contains(args, webFlag) {
		return Web{}
	}
	return Desktop{}
}
