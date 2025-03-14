package ui

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// HomeData contains the data needed to render the home page
type HomeData struct {
	Title                 string
	Description           string
	BackgroundInformation string
	Colors                ColorScheme
}

// MainContent renders the main content of the home page
func MainContent(homeData HomeData) g.Node {
	colors := homeData.Colors

	return Div(
		Class("min-h-screen bg-"+colors.Background),
		PageHeader(homeData),
		Main(
			Class("container mx-auto px-4 py-8"),
			Div(
				Class("max-w-4xl mx-auto"),
				H2(
					Class("text-2xl font-bold mb-6 text-"+colors.Text),
					g.Text("About This Project"),
				),
				P(
					Class("mb-4 text-"+colors.Text),
					g.Text(homeData.BackgroundInformation),
				),
				Div(
					Class("mt-8 p-6 bg-"+colors.Surface+" rounded-lg shadow-md border border-"+colors.Border),
					H3(
						Class("text-xl font-semibold mb-4 text-"+colors.Text),
						g.Text("Getting Started"),
					),
					P(
						Class("text-"+colors.Text),
						g.Text("To begin using this application, simply log in or create an account."),
					),
					Div(
						Class("mt-6"),
						A(
							Href("/login"),
							Class("inline-block px-5 py-3 bg-"+colors.Primary+" text-white font-medium rounded-md hover:bg-blue-600 transition-colors mr-4"),
							g.Text("Log In"),
						),
						A(
							Href("/register"),
							Class("inline-block px-5 py-3 bg-white text-"+colors.Primary+" font-medium rounded-md border border-"+colors.Primary+" hover:bg-gray-50 transition-colors"),
							g.Text("Register"),
						),
					),
				),
			),
		),
		PageFooter(homeData),
	)
}

// PageHeader renders the header section of the page
func PageHeader(homeData HomeData) g.Node {
	colors := homeData.Colors

	return Nav(
		Class("bg-"+colors.Surface+" shadow-sm border-b border-"+colors.Border),
		Div(
			Class("container mx-auto px-4 py-4"),
			Div(
				Class("flex justify-between items-center"),
				Div(
					Class("flex items-center"),
					A(
						Href("/"),
						Class("text-xl font-bold text-"+colors.Primary),
						g.Text("Go Web App"),
					),
				),
				Div(
					Class("hidden md:flex space-x-6"),
					A(
						Href("/features"),
						Class("text-"+colors.Text+" hover:text-"+colors.Primary+" transition-colors"),
						g.Text("Features"),
					),
					A(
						Href("/pricing"),
						Class("text-"+colors.Text+" hover:text-"+colors.Primary+" transition-colors"),
						g.Text("Pricing"),
					),
					A(
						Href("/docs"),
						Class("text-"+colors.Text+" hover:text-"+colors.Primary+" transition-colors"),
						g.Text("Documentation"),
					),
				),
				Div(
					Class("flex items-center space-x-4"),
					A(
						Href("/login"),
						Class("text-"+colors.Text+" hover:text-"+colors.Primary+" transition-colors"),
						g.Text("Log In"),
					),
					A(
						Href("/register"),
						Class("px-4 py-2 bg-"+colors.Primary+" text-white rounded-md hover:bg-blue-600 transition-colors"),
						g.Text("Sign Up"),
					),
				),
			),
		),
	)
}

// Hero renders the hero section of the home page
func Hero(homeData HomeData) g.Node {
	colors := homeData.Colors

	return Div(
		Class("bg-gradient-to-r from-"+colors.Primary+" to-"+colors.Accent+" py-20"),
		Div(
			Class("container mx-auto px-4"),
			Div(
				Class("max-w-3xl mx-auto text-center"),
				H1(
					Class("text-4xl md:text-5xl font-bold text-white mb-6"),
					g.Text(homeData.Title),
				),
				P(
					Class("text-xl text-white opacity-90 mb-8"),
					g.Text(homeData.Description),
				),
				Div(
					Class("flex flex-col sm:flex-row justify-center gap-4"),
					A(
						Href("/register"),
						Class("px-6 py-3 bg-white text-"+colors.Primary+" font-medium rounded-md hover:bg-gray-100 transition-colors"),
						g.Text("Get Started"),
					),
					A(
						Href("/learn-more"),
						Class("px-6 py-3 bg-transparent text-white border border-white font-medium rounded-md hover:bg-white/10 transition-colors"),
						g.Text("Learn More"),
					),
				),
			),
		),
	)
}

// PageFooter renders the footer section of the page
func PageFooter(homeData HomeData) g.Node {
	colors := homeData.Colors

	return Div(
		Class("bg-"+colors.Surface+" border-t border-"+colors.Border+" mt-16"),
		Div(
			Class("container mx-auto px-4 py-12"),
			Div(
				Class("grid grid-cols-1 md:grid-cols-4 gap-8"),
				Div(
					Class("md:col-span-1"),
					Div(
						Class("text-xl font-bold text-"+colors.Primary+" mb-4"),
						g.Text("Go Web App"),
					),
					P(
						Class("text-"+colors.TextLight),
						g.Text("A modern, scalable web application template."),
					),
				),
				Div(
					Class("space-y-2"),
					H3(
						Class("text-sm font-semibold uppercase tracking-wider text-"+colors.TextLight+" mb-3"),
						g.Text("Product"),
					),
					FooterLink("/features", "Features", colors),
					FooterLink("/pricing", "Pricing", colors),
					FooterLink("/roadmap", "Roadmap", colors),
				),
				Div(
					Class("space-y-2"),
					H3(
						Class("text-sm font-semibold uppercase tracking-wider text-"+colors.TextLight+" mb-3"),
						g.Text("Support"),
					),
					FooterLink("/docs", "Documentation", colors),
					FooterLink("/help", "Help Center", colors),
					FooterLink("/contact", "Contact Us", colors),
				),
				Div(
					Class("space-y-2"),
					H3(
						Class("text-sm font-semibold uppercase tracking-wider text-"+colors.TextLight+" mb-3"),
						g.Text("Legal"),
					),
					FooterLink("/privacy", "Privacy Policy", colors),
					FooterLink("/terms", "Terms of Service", colors),
				),
			),
			Div(
				Class("mt-8 pt-8 border-t border-"+colors.Border+" text-"+colors.TextLight+" text-center"),
				g.Text("© 2025 Go Web App. All rights reserved."),
			),
		),
	)
}

// FooterLink renders a link in the footer with consistent styling
func FooterLink(href, text string, colors ColorScheme) g.Node {
	return Div(
		A(
			Href(href),
			Class("text-"+colors.Text+" hover:text-"+colors.Primary+" transition-colors"),
			g.Text(text),
		),
	)
}