package ui

import (
	"net/http"
)

// HomeHandler renders the home page
func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		colors := DefaultColorScheme()
		
		// Create data for the home page with generic placeholder content
		homeData := HomeData{
			Title:       "Go Web Application",
			Description: "A modern, scalable web application template",
			BackgroundInformation: "This is a template for building scalable web applications with Go. " +
				"It includes a robust architecture with clean separation of concerns, database integration, " +
				"and a responsive UI built with gomponents.",
			Colors: colors,
		}
		
		// Create page data
		pageData := PageData{
			Title:       homeData.Title,
			Description: homeData.Description,
			Colors:      colors,
		}
		
		// Render the full page with home content
		page := Page(pageData, MainContent(homeData))
		
		// Set content type and render
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = page.Render(w) // Ignoring error for simplicity, in production handle this error
	}
}