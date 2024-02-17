package mailer

import (
	"braces.dev/errtrace"
	"context"
	"errors"
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
		EnqueueTask(ctx context.Context, taskName string, payload any) error
	}
)

// NewEnqueuer creates a new email enqueuer.
func NewEnqueuer(e workerEnqueuer) *Enqueuer {
	return &Enqueuer{queue: e}
}

// SendEmail enqueues a task to send an email.
// It implements the emailAdapter interface. So, it can be used as a mailer adapter in your application.
// This function returns an error if the task could not be enqueued.
func (e *Enqueuer) SendEmail(ctx context.Context, payload SendEmailPayload) error {
	// Validate the payload
	if err := payload.Validate(); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToEnqueueTask, err))
	}

	// Enqueue the task
	if err := e.queue.EnqueueTask(ctx, SendEmailTask, payload); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToEnqueueTask, err))
	}

	return nil
}
