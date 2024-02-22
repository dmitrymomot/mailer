package utils

import (
	"bytes"
	"context"
	"fmt"

	"braces.dev/errtrace"
	"github.com/a-h/templ"
)

// RenderToString renders the given templ.Component to a string.
// It returns the rendered template as an HTML string.
func RenderToString(ctx context.Context, t templ.Component) (string, error) {
	// string writter to store the rendered string
	var buf bytes.Buffer
	// render the component to the buffer
	if err := t.Render(ctx, &buf); err != nil {
		return "", errtrace.Wrap(fmt.Errorf("failed to render template: %w", err))
	}
	return buf.String(), nil
}
