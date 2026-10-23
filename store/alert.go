package store

import (
	"fmt"
	"strings"
)

// Alert is one custom timer. It is the same shape the web version keeps, so a
// name plus an hour/minute/second triplet travels between interfaces cleanly.
type Alert struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Hours   int    `json:"hours"`
	Minutes int    `json:"minutes"`
	Seconds int    `json:"seconds"`
}

// Total is the whole duration in seconds.
func (a Alert) Total() int { return a.Hours*3600 + a.Minutes*60 + a.Seconds }

// Valid reports whether the alert can be saved: a name that is not blank and
// a duration longer than zero.
func (a Alert) Valid() bool {
	return strings.TrimSpace(a.Name) != "" && a.Total() > 0
}

// Duration renders the alert's length as HH:MM:SS, e.g. "01:05:00".
func (a Alert) Duration() string { return ClockTime(a.Total()) }

// ClockTime renders a number of seconds as HH:MM:SS.
func ClockTime(total int) string {
	if total < 0 {
		total = 0
	}
	h := total / 3600
	m := (total % 3600) / 60
	s := total % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
