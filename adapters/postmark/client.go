package postmark

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitrymomot/mailer"
	"github.com/mrz1836/postmark"
)

type (
	// Postmark service wrapper
	Client struct {
		client *postmark.Client
		config Config
	}

	// Config struct
	Config struct {
		// From is the email address to use for the from header. This must be a verified sender signature.
		From string
		// ReplyTo is the email address to use for the reply-to header.
		ReplyTo string
		// TrackOpens is a boolean indicating if you want to track opens.
		TrackOpens bool
		// TrackLinks is a boolean indicating if you want to track links.
		TrackLinks bool
	}
)

// Validate the config
func (c Config) Validate() error {
	if c.From == "" {
		return ErrMissingFromEmail
	}
	return nil
}

// New func is a factory function,
// returns a new instance of the Client interface implementation
func New(serverToken, accountToken string, conf Config) (*Client, error) {
	if err := conf.Validate(); err != nil {
		return nil, errors.Join(ErrFailedToCreateClient, err)
	}
	if serverToken == "" || accountToken == "" {
		return nil, errors.Join(ErrFailedToCreateClient, ErrMissingCredentials)
	}
	if conf.ReplyTo == "" {
		conf.ReplyTo = conf.From
	}
	return &Client{
		client: postmark.NewClient(serverToken, accountToken),
		config: conf,
	}, nil
}

// SendEmail sends an email.
// It returns an error if sending fails.
// Parameters:
// - ctx: context.Context
// - emailAddr: email address of the recipient
// - subject: subject of the email
// - htmlBody: HTML body of the email
func (c *Client) SendEmail(ctx context.Context, emailAddr, subject, htmlBody string, attachments ...mailer.Attachment) error {
	// Create the email
	e := postmark.Email{
		TrackOpens: true,
		ReplyTo:    c.config.ReplyTo,
		TrackLinks: "HtmlOnly",
		From:       c.config.From,
		To:         emailAddr,
		Subject:    subject,
		HTMLBody:   htmlBody,
		Attachments: func() []postmark.Attachment {
			a := make([]postmark.Attachment, 0, len(attachments))
			for _, attachment := range attachments {
				a = append(a, postmark.Attachment{
					Name:        attachment.Name,
					Content:     attachment.Content,
					ContentType: attachment.ContentType,
					ContentID:   attachment.ContentID,
				})
			}
			return a
		}(),
	}

	// Apply the config
	if !c.config.TrackOpens {
		e.TrackOpens = false
	}
	if !c.config.TrackLinks {
		e.TrackLinks = "None"
	}

	// Send the email
	resp, err := c.client.SendEmail(ctx, e)
	if err != nil {
		return errors.Join(ErrFailedToSendEmail, err)
	}
	if resp.ErrorCode > 0 {
		return errors.Join(ErrFailedToSendEmail, fmt.Errorf("%v %s", resp.ErrorCode, resp.Message))
	}

	return nil
}
