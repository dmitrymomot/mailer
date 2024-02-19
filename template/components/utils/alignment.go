package utils

// Alignment represents the text alignment
type Alignment string

// String returns the string representation of the alignment
func (a Alignment) String() string {
	return string(a)
}

// Predefined alignments
const (
	AlignLeft   Alignment = "left"
	AlignCenter Alignment = "center"
	AlignRight  Alignment = "right"
)
