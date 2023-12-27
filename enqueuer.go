package mailer

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hibiken/asynq"
)

type (
	// Enqueuer is a helper struct for enqueuing email tasks.
	// It encapsulates the worker.Enqueuer struct and adds queue methods.
	// See pkg/worker/enqueuer.go.
	Enqueuer struct {
		queue workerEnqueuer
	}

	// workerEnqueuer is an interface for enqueuing tasks.
	workerEnqueuer interface {
		EnqueueTask(ctx context.Context, task *asynq.Task) error
	}
)

// NewEnqueuer creates a new email enqueuer.
func NewEnqueuer(e workerEnqueuer) *Enqueuer {
	return &Enqueuer{queue: e}
}

// SendEmail enqueues a task to send an email.
// This function returns an error if the task could not be enqueued.
// Parameters:
// - emailAddr: email address of the recipient
// - subject: subject of the email
// - htmlBody: HTML body of the email
func (e *Enqueuer) SendEmail(ctx context.Context, emailAddr, subject string, htmlBody string) error {
	payload, err := json.Marshal(SendEmailPayload{
		Email:    emailAddr,
		Subject:  subject,
		HTMLBody: htmlBody,
	})
	if err != nil {
		return errors.Join(ErrFailedToMarshalPayload, err)
	}

	if err := e.queue.EnqueueTask(ctx, asynq.NewTask(SendEmailTask, payload)); err != nil {
		return errors.Join(ErrFailedToEnqueueTask, err)
	}

	return nil
}
