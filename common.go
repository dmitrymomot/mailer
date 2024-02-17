package mailer

// Predefined task names.
const (
	SendEmailTask = "mailer.send_email" // SendEmailTask is a task name for sending an email.
)

// Predefined task payloads.
type (
	// SendEmailPayload is a payload for SendEmailTask.
	SendEmailPayload struct {
		// Email is the email address of the recipient.
		Email string `json:"email,omitempty"`
		// Subject is the subject of the email.
		Subject string `json:"subject,omitempty"`
		// HTMLBody is the HTML body of the email.
		HTMLBody string `json:"html_body,omitempty"`
		// Attachments is an optional encoded file to send along with an email (see Attachment).
		Attachments []Attachment `json:"attachments,omitempty"`
	}

	// Attachment is an optional encoded file to send along with an email
	Attachment struct {
		// Name: attachment name
		Name string `json:"name,omitempty"`
		// Content: Base64 encoded attachment data
		Content string `json:"content,omitempty"`
		// ContentType: attachment MIME type
		ContentType string `json:"content_type,omitempty"`
		// ContentId: populate for inlining images with the images cid
		ContentID string `json:"content_id,omitempty"`
	}
)

// Validate validates the SendEmailPayload.
func (p SendEmailPayload) Validate() error {
	if p.Email == "" {
		return ErrMissingEmail
	}
	if p.Subject == "" {
		return ErrMissingSubject
	}
	if p.HTMLBody == "" {
		return ErrMissingHTMLBody
	}
	return nil
}
