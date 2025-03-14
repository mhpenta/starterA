package handlers

import (
	"net/http"
	"youtubeMonitor/internal/ui"
)

// HomeHandler renders the home page
func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		colors := ui.DefaultColorScheme()

		// Create minimal page data
		homeData := ui.HomeData{
			Title:                 "Dark Mode App",
			Description:           "A sleek, minimal dark-themed web application",
			BackgroundInformation: "This lightweight application showcases a sophisticated dark design with high contrast and subtle elements.",
			Colors:                colors,
		}

		pageData := ui.PageData{
			Title:       homeData.Title,
			Description: homeData.Description,
			Colors:      colors,
		}

		page := ui.Page(pageData, ui.MainContent(homeData))

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = page.Render(w)
	}
}