package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
)

type figurePage struct {
	app.Compo

	Iclass      string
	Iname       string
	Ifigure     string
	Ipage       []string
	IpageScore  map[string]int
	IpageVisits map[string]int
	Icaption    string
	Icaptions   []string
	Iaudio      string
}

func (fp *figurePage) Name(v string) *figurePage {
	fp.Iname = v
	return fp
}

func (fp *figurePage) Page(v ...string) *figurePage {
	fp.Ipage = v
	return fp
}

func (fp *figurePage) Figure(v string) *figurePage {
	fp.Ifigure = v
	return fp
}

func (fp *figurePage) Audio(v string) *figurePage {
	fp.Iaudio = v
	return fp
}

func (fp *figurePage) Caption(v ...string) *figurePage {
	fp.Icaptions = v
	fp.Icaption = fp.Icaptions[figIndex]
	return fp
}

func (fp *figurePage) onFigureClicked(ctx app.Context, e app.Event) {
	if figIndex+1 < len(fp.Icaptions) {
		figIndex = figIndex + 1
	}
	if len(fp.Iaudio) != 0 {
		var myAudio = app.Window().GetElementByID("my-audio")
		myAudio.Call("play")
	}
	_, ok := globalScore.figureScores[fp.Iname]
	if ok {
		globalScore.figureScores[fp.Iname] = globalScore.figureScores[fp.Iname] + 1
	} else {
		globalScore.figureScores[fp.Iname] = 1
	}
	ctx.SessionStorage().Set(fp.Iname, globalScore.figureScores[fp.Iname])
	ctx.Update()
}

func (fp *figurePage) onFigureDoubleClicked(ctx app.Context, e app.Event) {
	if len(fp.Iaudio) != 0 {
		var myAudio = app.Window().GetElementByID("my-audio")
		myAudio.Call("pause")
	}
}

func newFigurePage() *figurePage {
	return &figurePage{
		IpageScore:  map[string]int{},
		IpageVisits: map[string]int{},
	}
}

func (fp *figurePage) OnMount(ctx app.Context) {
	var visits int
	ctx.SessionStorage().Get(fp.Iname+"Visits", &visits)
	ctx.SessionStorage().Set(fp.Iname+"Visits", visits+1)
}

func (fp *figurePage) Render() app.UI {
	shellClass := app.AppendClass("fill", "background", "center")
	return ui.Shell().
		Class(shellClass).
		Content(
			ui.Shell().
				Content(
					app.Main().Body(
						app.Figure().
							OnClick(fp.onFigureClicked).
							OnDblClick(fp.onFigureDoubleClicked).
							Class("scalable-figure", "center").
							Body(
								app.If(len(fp.Iaudio) != 0, func() app.UI {
									return app.Audio().Loop(true).Style("display", "none").ID("my-audio").Src(fp.Iaudio)
								}),
								app.FigCaption().Text(fp.Icaption).Class("text-center").Hidden(false),
								app.Img().Src(fp.Ifigure),
							),
					),
				),
		)

}
