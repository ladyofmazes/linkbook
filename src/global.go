package main

type ScoreBoard struct {
	buttonScore int
}

var globalScore ScoreBoard

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
