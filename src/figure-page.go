package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
)

type figurePage struct {
	app.Compo

	Iclass    string
	Ifigure   string
	Icaption  string
	Icaptions []string
}

func (fp *figurePage) Figure(v string) *figurePage {
	fp.Ifigure = v
	return fp
}

func (fp *figurePage) Caption(v ...string) *figurePage {
	fp.Icaptions = v
	fp.Icaption = fp.Icaptions[figIndex]
	return fp
}

func (fp *figurePage) onFigureClicked(ctx app.Context, e app.Event) {
	figIndex = figIndex + 1
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
					app.Main().Body(
						app.Figure().OnClick(fp.onFigureClicked).Class("scalable-figure", "center").Body(
							app.FigCaption().Text(fp.Icaption).Class("text-center").Hidden(false),
							app.Img().Src(fp.Ifigure),
						),
					),
				),
		)

}
