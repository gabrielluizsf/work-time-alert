package webpush

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestKeys(t *testing.T) {
	keysFirstCall := Keys()
	keysLastCall := Keys()
	assert.Equal(t, keysFirstCall, keysLastCall)
}
