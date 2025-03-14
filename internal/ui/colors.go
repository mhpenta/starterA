package ui

// ColorScheme defines a set of colors to be used throughout the application
type ColorScheme struct {
	Primary      string // Main brand color
	Secondary    string // Secondary brand color
	Background   string // Page background
	Surface      string // Surface elements like cards
	Text         string // Primary text color
	TextLight    string // Secondary/lighter text color
	Border       string // Border color
	Error        string // Error messages
	Success      string // Success messages
	Warning      string // Warning messages
	Info         string // Informational messages
	Accent       string // Accent color for highlights
}

// DefaultColorScheme returns a modern, accessible color scheme
func DefaultColorScheme() ColorScheme {
	return ColorScheme{
		Primary:    "#3B82F6", // Blue-500
		Secondary:  "#10B981", // Emerald-500
		Background: "#F9FAFB", // Gray-50
		Surface:    "#FFFFFF", // White
		Text:       "#1F2937", // Gray-800
		TextLight:  "#6B7280", // Gray-500
		Border:     "#E5E7EB", // Gray-200
		Error:      "#EF4444", // Red-500
		Success:    "#10B981", // Emerald-500
		Warning:    "#F59E0B", // Amber-500
		Info:       "#3B82F6", // Blue-500
		Accent:     "#8B5CF6", // Violet-500
	}
}

// DarkColorScheme returns a dark mode color scheme
func DarkColorScheme() ColorScheme {
	return ColorScheme{
		Primary:    "#3B82F6", // Blue-500
		Secondary:  "#10B981", // Emerald-500
		Background: "#111827", // Gray-900
		Surface:    "#1F2937", // Gray-800
		Text:       "#F9FAFB", // Gray-50
		TextLight:  "#9CA3AF", // Gray-400
		Border:     "#374151", // Gray-700
		Error:      "#EF4444", // Red-500
		Success:    "#10B981", // Emerald-500
		Warning:    "#F59E0B", // Amber-500
		Info:       "#3B82F6", // Blue-500
		Accent:     "#8B5CF6", // Violet-500
	}
}