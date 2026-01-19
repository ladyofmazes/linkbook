package lbook

type ScoreBoard struct {
	ButtonScore  int
	FigureScores map[string]int
}

var GlobalScore = ScoreBoard{
	FigureScores: map[string]int{},
	ButtonScore:  0}
var figIndex int

const (
	headerHeight = 100
)

const (
	getMarkdown = "/markdown/get"
)

type status int

const (
	neverLoaded status = iota
	loading
	loadingErr
	loaded
)

func markdownState(src string) string {
	return src
}

type markdownContent struct {
	Status status
	Err    error
	Data   string
}
