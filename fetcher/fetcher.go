package fetcher

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func Fetch(url string) (*html.Node, error)  {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
