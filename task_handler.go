package mailer

import (
	"context"
	"errors"

	"braces.dev/errtrace"
	"github.com/dmitrymomot/asyncer"
)

type (
	emailAdapter interface {
		// SendEmail sends an email.
		// It returns an error if sending fails.
		// SendEmail sends an email using the Postmark client.
		// It takes a context and a SendEmailPayload as input.
		// The SendEmailPayload contains the necessary information
		// for the email, such as the recipient, subject, body, and attachments.
		// The function returns an error if the email fails to send.
		SendEmail(ctx context.Context, payload SendEmailPayload) error
	}
)

// SendEmailHandler is a function that returns a task handler for sending emails.
// It takes a mail emailAdapter as a parameter and returns a asyncer.TaskHandler.
// See details about asyncer module at github.com/dmitrymomot/asyncer.
// The TaskHandler is responsible for executing the SendEmailTask with the given payload.
// It sends an email using the provided mail adapter and returns an error if the email fails to send.
func SendEmailHandler(mail emailAdapter) asyncer.TaskHandler {
	return asyncer.HandlerFunc(SendEmailTask, func(ctx context.Context, payload SendEmailPayload) error {
		if err := mail.SendEmail(ctx, payload); err != nil {
			return errtrace.Wrap(errors.Join(ErrFailedToSendEmail, err))
		}
		return nil
	})
}
