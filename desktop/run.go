package desktop

import (
	"log"
	"runtime"
	"time"

	"github.com/gabrielluizsf/antui"
	"github.com/gabrielluizsf/work-time-alert/store"
)

const appTitle = "Work Time Alert"

// Cocoa wants its window driven from the main thread, and Win32 from the one
// that made it; locking here, during init, pins the main goroutine to it.
func init() { runtime.LockOSThread() }

// Run opens the desktop window and runs it until it is closed. icon is the
// app's logo as PNG bytes; nil keeps the system's default icon.
func Run(icon []byte) {
	win, err := antui.OpenWith(antui.Options{
		Title:  appTitle,
		Width:  1100,
		Height: 760,
		Points: true,
		Limits: antui.Limits{MinWidth: 400, MinHeight: 560},
	})
	if err != nil {
		log.Fatalf("could not open the window: %v", err)
	}
	defer win.Close()
	win.SetFPS(30)
	if icons, err := decodeIcon(icon); err == nil {
		win.SetIcon(icons...)
	}

	chime := NewChime()
	defer chime.Close()

	app := NewApp(Options{
		Lang:  SystemLang(),
		Store: store.NewLocal(),
		Now:   time.Now,
		OnDone: func(title, message string) {
			chime.Play()
			win.SendNotification(title, message)
		},
	})

	v := &view{win: win}
	var p page
	title := ""
	for win.Begin() {
		if s := win.ContentScale(); s != v.scale || v.body == nil {
			v.scale = s
			v.loadFaces()
		}
		v.pal = paletteFor(win)
		v.applyTheme()
		p.frame(v, app)

		if t := windowTitle(app); t != title {
			title = t
			win.SetTitle(t)
		}
		win.End()
	}
}
