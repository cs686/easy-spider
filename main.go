package main

import (
	"github.com/cs686/easy-spider/engine"
	"github.com/cs686/easy-spider/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
