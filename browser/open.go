package browser

import (
	"runtime"

	"github.com/i9si-sistemas/command"
)

// Open opens the specified URL in the default web browser of the user.
func Open(url string) error {
	cmd := command.New()

	switch runtime.GOOS {
	case "windows":
		cmd = cmd.Execute("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = cmd.Execute("open", url)
	default:
		cmd = cmd.Execute("xdg-open", url)
	}

	return cmd.Run()
}
