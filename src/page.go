package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
)

const (
	headerHeight = 100
)

type page struct {
	app.Compo

	Iclass    string
	Iindex    []app.UI
	Iicon     string
	Ititle    string
	Icontent  []app.UI
	Ifootnote string
}

func newPage() *page {
	return &page{}
}

func (p *page) Index(v ...app.UI) *page {
	p.Iindex = app.FilterUIElems(v...)
	return p
}

func (p *page) Icon(v string) *page {
	p.Iicon = v
	return p
}

func (p *page) Title(v string) *page {
	p.Ititle = v
	return p
}

func (p *page) Content(v ...app.UI) *page {
	p.Icontent = app.FilterUIElems(v...)
	return p
}

func (p *page) Footnote(v string) *page {
	p.Ifootnote = v
	return p
}

func (p *page) Render() app.UI {
	shellClass := app.AppendClass("fill", "background")
	return ui.Shell().
		Class(shellClass).
		Index(
			app.If(len(p.Iindex) != 0, func() app.UI {
				return ui.Scroll().
					Class(shellClass).
					HeaderHeight(headerHeight).
					Content(
						app.Nav().
							Class("index").
							Body(
								app.Div().Class("separator"),
								app.Header().
									Class("h2").
									Text("Index"),
								app.Div().Class("separator"),
								app.Range(p.Iindex).Slice(func(i int) app.UI {
									return p.Iindex[i]
								}),
								app.Div().Class("separator"),
							),
					)
			}),
		).
		Content(
			ui.Scroll().
				Class(shellClass).
				HeaderHeight(headerHeight).
				Content(
					app.Main().Body(
						app.Article().Body(
							app.Header().
								ID("page-top").
								Class("page-title").
								Class("center").
								Body(
									ui.Stack().
										Center().
										Middle().
										Content(
											ui.Icon().
												Class("icon-left").
												Class("unselectable").
												Size(90).
												Src(p.Iicon),
											app.H1().Text(p.Ititle),
										),
								),
							app.Div().Class("separator"),
							app.Range(p.Icontent).Slice(func(i int) app.UI {
								return p.Icontent[i]
							}),
							app.Div().Class("separator"),
							app.If(len(p.Ifootnote) != 0, func() app.UI {
								return app.Aside().Body(
									app.Header().
										ID("footnote").
										Class("h2").
										Text("Footnote"),
									app.P().Body(
										app.Text(p.Ifootnote),
									),
								)
							}),
							app.Div().Class("separator"),
							app.Aside().Body(
								app.Header().
									ID("report-an-issue").
									Class("h2").
									Text(""),
								app.P().Body(
									app.Text("For more fun with me or to report an issue: "),
									app.A().
										Href("https://github.com/ladyofmazes/linkbook").
										Text("🚀 Find me on Github!"),
								),
							),
						),
					),
				),
		)
}
