package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

// PageData contains the data needed to render a complete page
type PageData struct {
	Title       string
	Description string
}

// HomeData contains the data needed to render the home page
type HomeData struct {
	Title                 string
	Description           string
	BackgroundInformation string
	Colors                ColorScheme
}

// Page renders a complete HTML page with the given content
func Page(pageData PageData, content Node) Node {
	return HTML5(HTML5Props{
		Title: pageData.Title,
		Head: []Node{
			Meta(Attr("name", "description"), Attr("content", pageData.Description)),
			Raw(`<script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>`),
		},
		Body: []Node{
			content,
		},
	})
}

// MainContent renders the main content of the home page
func MainContent(homeData HomeData) Node {
	return Div(
		Class("flex flex-col min-h-screen"),
		SimpleHeader(),
		Main(
			Class("flex-grow container mx-auto px-4 py-8"),
			Div(
				Class("max-w-2xl mx-auto"),
				H1(
					Class("text-3xl font-bold mb-4"),
					Text(homeData.Title),
				),
				P(
					Class("mb-8 text-lg"),
					Text(homeData.Description),
				),
				Div(
					Class("app-surface p-6 rounded-lg shadow-md border border-zinc-800"),
					H2(
						Class("text-xl font-semibold mb-3"),
						Text("About"),
					),
					P(
						Class("mb-4"),
						Text(homeData.BackgroundInformation),
					),
					Div(
						Class("mt-4"),
						A(
							Href("https://github.com/mhpenta/starterA"),
							Class("app-link"),
							Text("Documentation →"),
						),
					),
				),
			),
		),
		SimpleFooter(),
	)
}

// SimpleHeader renders a minimal header
func SimpleHeader() Node {
	return Header(
		Class("app-surface p-4 mb-6 border-b border-zinc-800"),
		Div(
			Class("container mx-auto flex justify-between items-center"),
			A(
				Href("/"),
				Class("font-bold text-xl app-primary"),
				Text("Go App"),
			),
			Nav(
				A(
					Href("/docs"),
					Class("app-link ml-6"),
					Text("Docs"),
				),
			),
		),
	)
}

// SimpleFooter renders a minimal footer
func SimpleFooter() Node {
	return Footer(
		Class("app-surface p-4 text-center border-t border-zinc-800 mt-auto"),
		Div(
			Class("container mx-auto"),
			P(
				Class("text-sm text-gray-400"),
				Text("© 2025 Go Web App"),
			),
		),
	)
}
