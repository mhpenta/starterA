package ui

// ColorScheme defines a minimal set of colors
type ColorScheme struct {
	Primary    string
	Background string
	Surface    string
	Text       string
	TextLight  string
}

// DefaultColorScheme returns a simple dark mode color scheme
func DefaultColorScheme() ColorScheme {
	return ColorScheme{
		Primary:    "#60A5FA", // Bright blue
		Background: "#111827", // Very dark blue/gray
		Surface:    "#1F2937", // Dark blue/gray
		Text:       "#F9FAFB", // Nearly white
		TextLight:  "#9CA3AF", // Light gray
	}
}
