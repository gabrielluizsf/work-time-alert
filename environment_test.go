package main

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestEnvironmentDefaultsToDesktop(t *testing.T) {
	assert.Equal(t, environmentFor(nil), Environment(Desktop{}))
	assert.Equal(t, environmentFor([]string{"--unknown"}), Environment(Desktop{}))
}

func TestEnvironmentWebFlagSelectsWeb(t *testing.T) {
	assert.Equal(t, environmentFor([]string{"--web"}), Environment(Web{}))
}
