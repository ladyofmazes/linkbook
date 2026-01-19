package lbook

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
)

type Page struct {
	app.Compo

	Iclass          string
	Iname           string
	Iindex          []app.UI
	Ibutton         string
	Ibuttonfunction func(ctx app.Context, e app.Event)
	Iicon           string
	Ipage           []string
	IpageScore      map[string]int
	IpageVisits     map[string]int
	Ititle          string
	Icontent        []app.UI
	Ifootnote       string
}

func NewPage() *Page {
	return &Page{
		IpageScore:  map[string]int{},
		IpageVisits: map[string]int{},
	}
}

func (p *Page) Page(v ...string) *Page {
	p.Ipage = v
	return p
}

func (p *Page) Score() *Page {
	GlobalScore.ButtonScore = GlobalScore.ButtonScore + 1
	return p
}

func (p *Page) OnButtonClicked(ctx app.Context, e app.Event) {
	p.Score()
}

func (p *Page) Index(v ...app.UI) *Page {
	p.Iindex = app.FilterUIElems(v...)
	return p
}

func (p *Page) Icon(v string) *Page {
	p.Iicon = v
	return p
}

func (p *Page) Title(v string) *Page {
	p.Ititle = v
	return p
}

func (p *Page) Content(v ...app.UI) *Page {
	p.Icontent = app.FilterUIElems(v...)
	return p
}

func (p *Page) Footnote(v string) *Page {
	p.Ifootnote = v
	return p
}

func (p *Page) Button(v string, buttonFunction func(ctx app.Context, e app.Event)) *Page {
	p.Ibutton = v
	p.Ibuttonfunction = buttonFunction
	return p
}

func (p *Page) OnMount(ctx app.Context) {

	var visits int
	ctx.SessionStorage().Get(p.Iname+"Visits", &visits)
	ctx.SessionStorage().Set(p.Iname+"Visits", visits+1)
}

func (p *Page) Render() app.UI {
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
									Text(fmt.Sprintf("Index %d", p.IpageVisits["cookies"])),
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
							app.If(len(p.Ibutton) != 0, func() app.UI {
								return app.Button().
									Class("button").
									Text(p.Ibutton).
									OnClick(p.Ibuttonfunction)
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
