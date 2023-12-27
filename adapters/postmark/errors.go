package postmark

import "errors"

// Predefined postmark adapter errors.
var (
	ErrFailedToSendEmail    = errors.New("failed to send email")
	ErrMissingFromEmail     = errors.New("missing from email address")
	ErrFailedToCreateClient = errors.New("failed to create client")
	ErrMissingCredentials   = errors.New("missing server or account token")
)
