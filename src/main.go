package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type linkbook struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// markdown file is displayed as content.
//
//go:embed documents/entry1.md
var entry1Content string

func (h *linkbook) Render() app.UI {
	return newPage().
		Title("Entry 1").
		Index(
			app.Div().Class("separator"),
			app.Text("A start\n"),
			app.Div().Class("separator"),
			app.A().Href("https://www.google.com").Text("The beginning"),
		).
		Icon(manFaceSVG).
		Button("Score", newPage().onButtonClicked).
		Content(
			newMarkdownDoc().MD(entry1Content), // Use embedded content directly
			app.Div().Class("table"),
		).
		Footnote(fmt.Sprintf("Score: %d", globalScore.buttonScore))
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the components with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", func() app.Composer { return &linkbook{} })
	app.Route("/intro", func() app.Composer { return &intro{} })

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Add this check - if we're in browser, block forever
	if app.IsClient {
		select {} // Block forever - prevent Go runtime from exiting
	}

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Home",
		Description: "Home Page",
		Resources:   app.LocalDir("."),
		Styles: []string{
			"/app.css",
			"/web/css/prism.css",
			"/web/css/docs.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
