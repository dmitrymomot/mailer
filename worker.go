package mailer

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hibiken/asynq"
)

type (
	// Worker is a task handler for email delivery.
	Worker struct {
		mail emailAdapter
	}

	emailAdapter interface {
		// SendEmail sends an email.
		// It returns an error if sending fails.
		// Parameters:
		// - ctx: context.Context
		// - emailAddr: email address of the recipient
		// - subject: subject of the email
		// - htmlBody: HTML body of the email
		SendEmail(ctx context.Context, emailAddr, subject, htmlBody string, attachments ...Attachment) error
	}
)

// NewWorker creates a new email task handler.
func NewWorker(mail emailAdapter) *Worker {
	return &Worker{mail: mail}
}

// Register registers task handlers for email delivery.
func (w *Worker) Register(mux *asynq.ServeMux) {
	mux.HandleFunc(SendEmailTask, w.SendEmail)
}

// SendEmail sends an example confirmation email.
// It is a task handler for SendEmailTask.
func (w *Worker) SendEmail(ctx context.Context, t *asynq.Task) error {
	var p SendEmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Join(ErrInvalidTaskPayload, err)
	}

	if err := w.mail.SendEmail(ctx, p.Email, p.Subject, p.HTMLBody); err != nil {
		return errors.Join(ErrFailedToSendEmail, err)
	}

	return nil
}
