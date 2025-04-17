package server

import webpush "github.com/i9si-sistemas/web-push"

type Subscription interface {
	WebPush() *webpush.Subscription
}
