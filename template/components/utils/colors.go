package utils

// Color represents a color
type (
	Color string
)

// String returns the string representation of the color
func (c Color) String() string {
	return string(c)
}

// Predefined colors
const (
	Primary   Color = ""
	Secondary Color = ""
	Danger    Color = ""
	Success   Color = ""
	Warning   Color = ""
	Info      Color = ""
)

// Predefined text colors
const (
	TextHeader    Color = ""
	TextPrimary   Color = ""
	TextSecondary Color = ""
	TextSuccess   Color = ""
	TextDanger    Color = ""
	TextWarning   Color = ""
	TextInfo      Color = ""
)
