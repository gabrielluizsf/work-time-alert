package session

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gabrielluizsf/work-time-alert/server/logger"
	"github.com/gabrielluizsf/work-time-alert/server/webpush"

	"github.com/i9si-sistemas/nine"
	i9 "github.com/i9si-sistemas/nine/pkg/server"
)

func Routes(server nine.Server) {
	logger := logger.New()

	serviceWorker := server.Group("/notifier")

	serviceWorker.Post("/session", func(c *i9.Context) error {
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
	})

	serviceWorker.Get("/public", func(c *i9.Context) error {
		key := webpush.Keys().PublicKey
		return c.JSON(nine.JSON{
			"key": key,
		})
	})

	serviceWorker.Post("/subscribe", func(c *i9.Context) error {
		var body struct {
			SessionId string `json:"sessionId"`
			Sub       Sub    `json:"subscription"`
		}
		if err := c.BodyParser(&body); err != nil {
			logger.Data("{status:400}", "{error:"+err.Error()+"}")
			return c.SendStatus(http.StatusBadRequest)
		}
		err := Manager.Update(body.SessionId, body.Sub)
		if err != nil {
			logger.Data("{status:404}", "{error:"+err.Error()+"}")
			return c.SendStatus(http.StatusNotFound)
		}
		return c.SendStatus(http.StatusCreated)
	})

	serviceWorker.Post("/notify", func(c *i9.Context) error {
		var body struct {
			SessionId string `json:"sessionId"`
			Title     string `json:"title"`
			Body      string `json:"body"`
		}
		if err := c.BodyParser(&body); err != nil {
			logger.Data("{status:400}", "{error:"+err.Error()+"}")
			return c.SendStatus(http.StatusBadRequest)
		}
		err := webpush.Notify(body.SessionId, webpush.SubscriptionMessage{
			SenderId: body.SessionId,
			Title:    body.Title,
			Body:     body.Body,
		}, Manager)
		if err != nil {
			logger.Data("{status:500}", "{error:"+err.Error()+"}")
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.SendStatus(http.StatusCreated)
	})

}
