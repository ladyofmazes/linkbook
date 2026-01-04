package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
)

type figurePage struct {
	app.Compo

	Iclass  string
	Ifigure string
}

func (fp *figurePage) Figure(v string) *figurePage {
	fp.Ifigure = v
	return fp
}

func newFigurePage() *figurePage {
	return &figurePage{}
}

func (fp *figurePage) Render() app.UI {
	shellClass := app.AppendClass("fill", "background", "center")
	return ui.Shell().
		Class(shellClass).
		Content(
			ui.Shell().
				Content(
					app.Main().Body(app.Figure().Class("scalable-figure", "center").Body(app.Img().Src(fp.Ifigure))),
				),
		)

}
