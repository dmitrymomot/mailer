package main

import (
	"time"

	"github.com/dmitrymomot/go-env"
)

// The app environment variables.
var (
	// Postmark.
	postmarkServerToken  = env.MustString("POSTMARK_SERVER_TOKEN")
	postmarkAccountToken = env.MustString("POSTMARK_ACCOUNT_TOKEN")

	// Redis.
	redisConnString = env.MustString("REDIS_CONNECTION_STRING")

	// Queue.
	queueName         = env.GetString("QUEUE_NAME", "mailer")
	queueTaskDeadline = env.GetDuration("QUEUE_TASK_DEADLINE", 5*time.Minute)
	queueMaxRetry     = env.GetInt("QUEUE_MAX_RETRY", 3)
	workerConcurrency = env.GetInt("QUEUE_WORKER_CONCURRENCY", 3)
)
