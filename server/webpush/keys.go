package webpush

import (
	"context"

	"github.com/i9si-sistemas/nine"
	wp "github.com/i9si-sistemas/web-push"
)

type keys struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

var (
	Context       = context.Background()
	keysCache     = keys{}
	webpushClient = wp.New(nine.New(Context))
)

func Keys() keys {
	if len(keysCache.PrivateKey) > 0 && len(keysCache.PublicKey) > 0 {
		return keysCache
	}
	privateKey, publicKey, err := webpushClient.GenerateVAPIDKeys()
	if err != nil {
		panic(err)
	}
	keysCache.PrivateKey = privateKey
	keysCache.PublicKey = publicKey
	return keysCache
}
