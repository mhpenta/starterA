package ui

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// PageData contains the data needed for every page
type PageData struct {
	Title       string
	Description string
	Colors      ColorScheme
}

// Page renders a complete HTML page with the given content
func Page(pageData PageData, content g.Node) g.Node {
	return HTML(
		g.Attr("lang", "en"),
		Head(
			Meta(g.Attr("charset", "UTF-8")),
			Meta(g.Attr("name", "viewport", "content", "width=device-width, initial-scale=1.0")),
			Title(pageData.Title),
			Meta(g.Attr("name", "description", "content", pageData.Description)),
			// Include Tailwind CSS via CDN
			Link(g.Attr("href", "https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css", "rel", "stylesheet")),
		),
		Body(
			Class("antialiased"),
			content,
			// Optional: Add JavaScript for interactivity
			Script(g.Raw(`
				// Custom JavaScript can be added here
				document.addEventListener('DOMContentLoaded', () => {
					console.log('YouTube Monitor loaded');
				});
			`)),
		),
	)
}
