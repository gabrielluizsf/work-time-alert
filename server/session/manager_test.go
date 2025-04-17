package session

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestManager(t *testing.T) {
	m := manager{}
	m.Register(Subscription{
		SenderId: "141234214231432",
		Subscription: Sub{
			Endpoint: "google.com/notifier",
			Keys: Keys{
				Auth:   "3dj189d481",
				P256dh: "fhbih3921f4",
			},
		},
	})
	m.Register(Subscription{
		SenderId: "2931924803491",
		Subscription: Sub{
			Endpoint: "test.com",
			Keys: Keys{
				Auth:   "31239103141",
				P256dh: "4312483214983",
			},
		},
	})
	sub, err := m.Find("141234214231432")
	assert.NoError(t, err)
	webpushSubscription := sub.WebPush()
	assert.Equal(t, webpushSubscription.Endpoint, "google.com/notifier")
	assert.Equal(t, webpushSubscription.Keys.Auth, "3dj189d481")
	assert.Equal(t, webpushSubscription.Keys.P256dh, "fhbih3921f4")
}
