package main

import "github.com/gabrielluizsf/work-time-alert/desktop"

type Desktop struct{}

func (d Desktop) Start() {
	logo, _ := publicFiles.ReadFile("public/logo.png")
	desktop.Run(logo)
}
