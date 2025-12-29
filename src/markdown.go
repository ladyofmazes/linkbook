package main

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
