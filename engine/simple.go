package engine

import (
	"github.com/cs686/easy-spider/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (s *SimpleEngine)Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)
		if err != nil {
			panic(err)
		}

		requests = append(requests, result.Result...)
		for _, item := range result.Items {
			log.Printf("Got Item %v", item)
		}
	}

}

func worker(r Request) (ParserResult, error) {
	doc, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParserResult{}, nil
	}
	return  r.ParserFunc(doc), nil
}
