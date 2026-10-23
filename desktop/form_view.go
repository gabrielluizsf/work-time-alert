package desktop

import "github.com/gabrielluizsf/antui/canvas"

// formHeight is how tall alertForm draws.
func (v *view) formHeight() int {
	label := v.small.Height() + v.px(6)
	return label + v.px(44+16) + label + v.px(44+16+44)
}

// alertForm draws the name, the hours/minutes/seconds row and the actions.
// It returns the areas of its fields, so Enter in any of them can submit.
func (v *view) alertForm(app *App, x, y, width int) []canvas.Area {
	lang := app.Lang
	label := v.small.Height() + v.px(6)
	h := v.px(44)
	gap := v.px(12)

	y += label
	name := canvas.Area{X: x, Y: y, Width: width, Height: h}
	v.input(name, lang.T("nameLabel"), &app.Form.Name)
	y += h + v.px(16) + label

	fw := (width - 2*gap) / 3
	fields := []canvas.Area{name}
	for i, f := range []struct {
		key  string
		text *string
	}{
		{"hoursLabel", &app.Form.Hours},
		{"minutesLabel", &app.Form.Minutes},
		{"secondsLabel", &app.Form.Seconds},
	} {
		area := canvas.Area{X: x + i*(fw+gap), Y: y, Width: fw, Height: h}
		v.input(area, lang.T(f.key), f.text)
		fields = append(fields, area)
	}
	y += h + v.px(16)

	submit, cancel := lang.T("addAlert"), lang.T("cancel")
	if app.Form.Editing() {
		submit, cancel = lang.T("save"), lang.T("cancelEdit")
	}
	sw, cw := v.buttonWidth(submit), v.buttonWidth(cancel)
	right := x + width
	if v.button(canvas.Area{X: right - sw, Y: y, Width: sw, Height: h}, submit, primaryButton, true) {
		app.Submit()
	}
	if v.button(canvas.Area{X: right - sw - gap - cw, Y: y, Width: cw, Height: h}, cancel, textButton, true) {
		app.ClearForm()
	}
	return fields
}
