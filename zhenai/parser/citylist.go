package parser

import (
	"github.com/antchfx/htmlquery"
	"github.com/cs686/easy-spider/engine"
	"golang.org/x/net/html"
)

func ParseCityList(doc *html.Node) engine.ParserResult {
	list, err := htmlquery.QueryAll(doc, "//dl/dd/a")
	if err != nil {
		panic(err)
	}
	result := engine.ParserResult{}
	for _, n := range list {
		a := htmlquery.FindOne(n, "//a")
		result.Items = append(result.Items, htmlquery.InnerText(a))
		result.Result = append(result.Result, engine.Request{
			Url:        htmlquery.SelectAttr(a, "href"),
			ParserFunc: ParseCity,
		})
	}
	return result
}