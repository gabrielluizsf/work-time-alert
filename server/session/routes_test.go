package session

import (
	"testing"

	"github.com/gabrielluizsf/work-time-alert/spy"
	"github.com/i9si-sistemas/assert"
)

func TestRoutes(t *testing.T) {
	serverSpy := spy.NewServer()
	Routes(serverSpy)
	assert.Equal(t, len(serverSpy.GroupCalls), 1)
	assert.Equal(t, len(serverSpy.GetCalls), 1)
	assert.Equal(t, len(serverSpy.PostCalls), 3)
	serviceWorkerGroup := serverSpy.GroupCalls[0]
	assert.Equal(t, "/notifier", serviceWorkerGroup.Prefix)
	assert.Equal(t, serverSpy.PostCalls[0].Path, "/session")
	assert.Equal(t, serverSpy.GetCalls[0].Path, "/public")
	assert.Equal(t, serverSpy.PostCalls[1].Path, "/subscribe")
	assert.Equal(t, serverSpy.PostCalls[2].Path, "/notify")
}
