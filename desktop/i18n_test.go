package desktop

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func env(vars map[string]string) func(string) string {
	return func(k string) string { return vars[k] }
}

func TestDetectLang(t *testing.T) {
	assert.Equal(t, DetectLang(env(map[string]string{"LANG": "pt_BR.UTF-8"})), Portuguese)
	assert.Equal(t, DetectLang(env(map[string]string{"LANG": "en_US.UTF-8"})), English)
	assert.Equal(t, DetectLang(env(nil)), English)
	// LC_ALL wins over LANG, as it does for every program on the system.
	assert.Equal(t, DetectLang(env(map[string]string{"LC_ALL": "en_US", "LANG": "pt_BR"})), English)
}

func TestEveryKeyIsTranslatedInBothLanguages(t *testing.T) {
	for key := range translations[Portuguese] {
		_, ok := translations[English][key]
		assert.True(t, ok, key)
	}
	assert.Equal(t, len(translations[Portuguese]), len(translations[English]))
	assert.Equal(t, Portuguese.T("missing"), "missing")
}
