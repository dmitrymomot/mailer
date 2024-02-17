package mailer

import "errors"

// Predefined errors.
var (
	ErrInvalidTaskPayload     = errors.New("invalid task payload")
	ErrFailedToSendEmail      = errors.New("failed to send email")
	ErrFailedToMarshalPayload = errors.New("failed to marshal payload")
	ErrFailedToEnqueueTask    = errors.New("failed to enqueue task")
	ErrMissingEmail           = errors.New("recipient email address must not be empty")
	ErrMissingSubject         = errors.New("email subject must not be empty")
	ErrMissingHTMLBody        = errors.New("email body cannot be empty")
)
