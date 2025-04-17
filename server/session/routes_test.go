package session

import (
	"testing"

	"github.com/gabrielluizsf/work-time-alert/spy"
	"github.com/i9si-sistemas/assert"
)

func TestRoutes(t *testing.T) {
	serverSpy := spy.NewServer()
	Routes(serverSpy, spy.NewLogger())
	assert.Equal(t, len(serverSpy.GroupCalls), 1)
	assert.Equal(t, len(serverSpy.GetCalls), 1)
	assert.Equal(t, len(serverSpy.PostCalls), 3)
	serviceWorkerGroup := serverSpy.GroupCalls[0]
	assert.Equal(t, "/notifier", serviceWorkerGroup.Prefix)
	assert.Equal(t, serverSpy.PostCalls[0].Path, "/notifier/session")
	assert.Equal(t, serverSpy.GetCalls[0].Path, "/notifier/public")
	assert.Equal(t, serverSpy.PostCalls[1].Path, "/notifier/subscribe")
	assert.Equal(t, serverSpy.PostCalls[2].Path, "/notifier/notify")
}
