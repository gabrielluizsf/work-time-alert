package session

import (
	"errors"
	"work-time-alert/server"

	"github.com/i9si-sistemas/nine"
	webpush "github.com/i9si-sistemas/web-push"
)

var (
	Manager = manager{}
    ErrNotFound = errors.New("session not found")
)

type manager nine.GenericJSON[string, *Subscription]

type Subscription struct {
	SenderId     string `json:"senderId"`
	Subscription Sub    `json:"subscription"`
}

func (sub Subscription) WebPush() *webpush.Subscription {
	return &webpush.Subscription{
		Endpoint: sub.Subscription.Endpoint,
		Keys: webpush.Keys{
			Auth:   sub.Subscription.Keys.Auth,
			P256dh: sub.Subscription.Keys.P256dh,
		},
	}
}

type Sub struct {
	Endpoint string `json:"endpoint"`
	Keys     Keys   `json:"keys"`
}

type Keys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

func (m manager) Register(
	subscription Subscription,
) {
	m[subscription.SenderId] = &subscription
}


func (m manager) Update(id string, sub Sub) error {
	data, err := m.Get(id)
	if err != nil {
		return err
	}
	data.Subscription = sub
	return nil
}


func (m manager) Find(
	id string,
) (server.Subscription, error) {
	sub, err := m.Get(id)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (m manager) Get(id string) (*Subscription, error) {
	sub, ok := m[id]
	if !ok {
		return nil, ErrNotFound
	}
	return sub, nil
}