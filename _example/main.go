package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dmitrymomot/asyncer"
	"github.com/dmitrymomot/mailer"
	"github.com/dmitrymomot/mailer/adapters/postmark"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Create a new email provider client.
	postmarkAdapter, err := postmark.New(postmarkServerToken, postmarkAccountToken, postmark.Config{
		From:       "sender@email",
		ReplyTo:    "replyto@address",
		TrackOpens: true,
		TrackLinks: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = postmarkAdapter

	// Init client
	asynqClient, asyncConnOpt, err := asyncer.NewClient(redisConnString)
	if err != nil {
		log.Fatal(err)
	}
	defer asynqClient.Close()

	// Create a new enqueuer
	enqueuer := asyncer.NewEnqueuer(
		asynqClient,
		asyncer.WithQueueNameEnq(queueName),
		asyncer.WithTaskDeadline(queueTaskDeadline),
		asyncer.WithMaxRetry(queueMaxRetry),
	)

	// Init error group
	eg, _ := errgroup.WithContext(context.Background())

	// Queue server
	{
		// Create a new queue worker server with the given options
		queueServer := asyncer.NewQueueServer(
			asyncConnOpt,
			asyncer.WithQueueName(queueName),
			asyncer.WithQueueConcurrency(workerConcurrency),
		)
		defer queueServer.Shutdown()

		// Run the worker
		eg.Go(queueServer.Run(
			mailer.NewWorker(postmarkAdapter),
		))
	}

	// Setup mailer enqueuer
	mailerEnqueuer := mailer.NewEnqueuer(enqueuer)

	// HTTP server
	{
		// Create a new HTTP handler
		httpHandler := newHTTPHandler(mailerEnqueuer)

		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Get("/", httpHandler.SendEmail)

		// Run the HTTP server
		eg.Go(func() error {
			if err := http.ListenAndServe(":8080", r); err != nil && err != http.ErrServerClosed {
				return fmt.Errorf("failed to listen and serve: %w", err)
			}
			return nil
		})
	}

	// Wait for the server to finish
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
