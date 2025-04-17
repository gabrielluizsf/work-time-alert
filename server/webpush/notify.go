package webpush

import (
	"encoding/json"
	"work-time-alert/server"

	webpush "github.com/i9si-sistemas/web-push"
)

type SubscriptionMessage struct {
	SenderId string `json:"senderId"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

type SessionManager interface {
	Find(id string) (server.Subscription, error)
}

func Notify(
	sessionId string,
	message SubscriptionMessage,
	sessionManager SessionManager,
) error {
	sub, err := sessionManager.Find(sessionId)
	if err != nil {
		return err
	}
	var (
		publicKey  = Keys().PublicKey
		privateKey = Keys().PrivateKey
		ttl        = 30
	)
	b, _ := json.Marshal(&message)
	webpushSub := sub.WebPush()
	_, err = webpushClient.SendNotification(b, webpushSub, &webpush.Options{
		VAPIDPublicKey:  publicKey,
		VAPIDPrivateKey: privateKey,
		TTL:             ttl,
	})
	return err
}
