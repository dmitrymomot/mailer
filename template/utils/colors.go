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
	Transparent Color = "transparent"
	Primary     Color = "#6366f1"
	Secondary   Color = "#6b7280"
	Danger      Color = "#ef4444"
	Success     Color = "#22c55e"
	Warning     Color = "#eab308"
	Info        Color = "#3b82f6"
)

// Predefined background colors
const (
	BgWhite   Color = "#ffffff"
	BgGray100 Color = "#f3f4f6"
	BgGray200 Color = "#e5e7eb"
)

// Predefined text colors
const (
	TextHeader    Color = "#1f2937"
	TextPrimary   Color = "#1f2937"
	TextSecondary Color = "#4b5563"
	TextSuccess   Color = "#166534"
	TextDanger    Color = "#991b1b"
	TextWarning   Color = "#854d0e"
	TextInfo      Color = "#1e40af"
	TextLink      Color = "#0284c7"
)
