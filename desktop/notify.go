package desktop

import (
	"runtime"
	"strings"

	"github.com/i9si-sistemas/command"
)

// Notify shows a system notification, so the alert is seen even when the
// window is hidden behind others. It is best effort: a system without a
// notifier still has the dialog inside the window.
func Notify(title, body string) error {
	cmd := command.New()
	switch runtime.GOOS {
	case "darwin":
		script := "display notification " + appleString(body) + " with title " + appleString(title)
		cmd = cmd.Execute("osascript", "-e", script)
	case "windows":
		return nil
	default:
		cmd = cmd.Execute("notify-send", "--app-name=Work Time Alert", title, body)
	}
	return cmd.Run()
}

// appleString quotes s as an AppleScript string literal.
func appleString(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return `"` + s + `"`
}
