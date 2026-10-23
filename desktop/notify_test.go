package desktop

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestAppleStringEscapesQuotes(t *testing.T) {
	assert.Equal(t, appleString(`Hora "de" \ foco`), `"Hora \"de\" \\ foco"`)
}
