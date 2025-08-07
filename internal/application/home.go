package application

import (
	"net/http"
	"starterA/internal/ui"
)

func (app *Application) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		colors := ui.DefaultColorScheme()

		homeData := ui.HomeData{
			Title:                 "Dark Mode App",
			Description:           "A sleek, minimal dark-themed web application",
			BackgroundInformation: "This lightweight application showcases a sophisticated dark design with high contrast and subtle elements.",
			Colors:                colors,
		}

		pageData := ui.PageData{
			Title:       homeData.Title,
			Description: homeData.Description,
		}

		page := ui.Page(pageData, ui.MainContent(homeData))

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = page.Render(w)
	}
}