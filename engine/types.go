package engine

import "golang.org/x/net/html"

type ParserResult struct {
	Result []Request
	Items []interface{}
}

type Request struct {
	Url string
	ParserFunc func(*html.Node) ParserResult
}

func NilParser(*html.Node) ParserResult {
	return ParserResult{}
}
