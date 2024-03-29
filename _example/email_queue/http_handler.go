package main

import (
	"context"
	"net/http"

	"github.com/dmitrymomot/mailer"
)

type (
	// httpHandler is a HTTP handler.
	httpHandler struct {
		mailSender mailSender
	}

	// mailSender is a mail sender.
	mailSender interface {
		SendEmail(ctx context.Context, payload mailer.SendEmailPayload) error
	}
)

// newHTTPHandler creates a new HTTP handler.
func newHTTPHandler(mailSender mailSender) *httpHandler {
	return &httpHandler{
		mailSender: mailSender,
	}
}

// SendEmail is a HTTP handler for sending email.
// It takes a request and returns a response.
// It is used in the HTTP server.
func (h *httpHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	// Send email
	if err := h.mailSender.SendEmail(r.Context(), mailer.SendEmailPayload{
		Email:    email,
		Subject:  "Test Email",
		HTMLBody: "<h1>Hi there!</h1>",
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
