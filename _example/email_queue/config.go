package main

import (
	"github.com/dmitrymomot/go-env"
)

// The app environment variables.
var (
	// Postmark.
	postmarkServerToken  = env.MustString("POSTMARK_SERVER_TOKEN")
	postmarkAccountToken = env.MustString("POSTMARK_ACCOUNT_TOKEN")

	// Redis.
	redisConnString = env.MustString("REDIS_CONNECTION_STRING")
)
