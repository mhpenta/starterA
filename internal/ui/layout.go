package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// PageData contains the data needed for every page
type PageData struct {
	Title       string
	Description string
	Colors      ColorScheme
}

// Page renders a complete HTML page with the given content
func Page(pageData PageData, content Node) Node {
	colors := pageData.Colors
	styleContent := `
		html, body { 
			background-color: ` + colors.Background + `; 
			color: ` + colors.Text + `; 
			font-family: system-ui, -apple-system, sans-serif;
			height: 100%;
			margin: 0;
		}
		.app-primary { color: ` + colors.Primary + `; }
		.app-bg { background-color: ` + colors.Background + `; }
		.app-surface { background-color: ` + colors.Surface + `; }
		.app-link { 
			color: ` + colors.Primary + `; 
			text-decoration: none; 
		}
		.app-link:hover { 
			text-decoration: underline; 
			opacity: 0.9;
		}
		::selection {
			background-color: ` + colors.Primary + `; 
			color: ` + colors.Background + `; 
		}
		.min-h-screen {
			min-height: 100vh;
		}
		.flex {
			display: flex;
		}
		.flex-col {
			flex-direction: column;
		}
		.flex-grow {
			flex-grow: 1;
		}
		.mt-auto {
			margin-top: auto;
		}
	`
	
	scriptContent := `document.addEventListener('DOMContentLoaded', () => {
		console.log('App loaded');
	});`
	
	return HTML(
		Attr("lang", "en"),
		Head(
			Meta(Attr("charset", "UTF-8")),
			Meta(Attr("name", "viewport"), Attr("content", "width=device-width, initial-scale=1.0")),
			El("title", Text(pageData.Title)),
			Meta(Attr("name", "description"), Attr("content", pageData.Description)),
			Link(Attr("href", "https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css"), Attr("rel", "stylesheet")),
			El("style", Raw(styleContent)),
		),
		Body(
			Class("antialiased app-bg"),
			content,
			El("script", Raw(scriptContent)),
		),
	)
}