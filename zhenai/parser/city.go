package parser

import (
	"github.com/cs686/easy-spider/engine"
	"github.com/cs686/easy-spider/model"
	"golang.org/x/net/html"
	"github.com/antchfx/htmlquery"

)

func ParseCity(doc *html.Node) engine.ParserResult  {
	all, err := htmlquery.QueryAll(doc, "//div[@class='g-list']/div[@class='list-item']")
	if err != nil {
		panic(err)
	}
	result := engine.ParserResult{}
	profile := model.Profile{}
	for _, content := range all {
		photoHtml := htmlquery.FindOne(content, "//div[@class='photo']//img")
		photoUrl := htmlquery.SelectAttr(photoHtml, "src")
		profile.Photo = photoUrl

		profileInfo := htmlquery.Find(content, "//div[@class='content']//tbody")
		for _, pro := range profileInfo {
			nameHtml := htmlquery.FindOne(pro, "//tr[1]//a")
			name := htmlquery.InnerText(nameHtml)
			profile.Name = name

			gender := htmlquery.FindOne(pro, "//tr[2]/td[1]/text()")
			profile.Gender = htmlquery.InnerText(gender)

			address := htmlquery.FindOne(pro, "//tr[2]/td[2]/text()")
			profile.Address = htmlquery.InnerText(address)

			age := htmlquery.FindOne(pro, "//tr[3]/td[1]/text()")
			profile.Age = htmlquery.InnerText(age)

			education := htmlquery.FindOne(pro, "//tr[3]/td[2]/text()")
			profile.Education = htmlquery.InnerText(education)

			marriage := htmlquery.FindOne(pro, "//tr[4]/td[1]/text()")
			profile.Marriage = htmlquery.InnerText(marriage)

			height := htmlquery.FindOne(pro, "//tr[4]/td[2]/text()")
			profile.Height = htmlquery.InnerText(height)
		}

		result.Items = append(result.Items, profile)
		//result.Result = append(result.Result, engine.Request{
		//	Url:        htmlquery.SelectAttr(text, "href"),
		//	ParserFunc: engine.NilParser,
		//})
	}
	return result
}