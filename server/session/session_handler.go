package session

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/i9si-sistemas/nine"
	i9 "github.com/i9si-sistemas/nine/pkg/server"
)

func handler(logger Logger) i9.HandlerWithContext {
	return func(c *i9.Context) error {
		sessionBytes := make([]byte, 32)
		_, err := rand.Read(sessionBytes)
		if err != nil {
			logger.Data("{status:500}", "{error:"+err.Error()+"}")
			return c.SendStatus(http.StatusInternalServerError)
		}
		sessionId := hex.EncodeToString(sessionBytes)
		Manager.Register(Subscription{
			SenderId:     sessionId,
			Subscription: Sub{},
		})
		return c.JSON(nine.JSON{"sessionId": sessionId})
	}
}
