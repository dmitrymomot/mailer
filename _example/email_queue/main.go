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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, _ := errgroup.WithContext(ctx)

	// Run a new queue server with redis as the broker.
	eg.Go(asyncer.RunQueueServer(
		ctx, redisConnString, nil,
		// Register a handler for the task.
		mailer.SendEmailHandler(postmarkAdapter),
		// ... add more handlers here ...
	))

	// Create a new enqueuer with redis as the broker.
	enqueuer := asyncer.MustNewEnqueuer(redisConnString)
	defer enqueuer.Close()

	// Create a new mail enqueuer.
	mailEnqueuer := mailer.NewEnqueuer(enqueuer)

	// HTTP server
	{
		// Create a new HTTP handler
		httpHandler := newHTTPHandler(mailEnqueuer)

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
