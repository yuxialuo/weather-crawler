package engine

import (
	"log"
	"strings"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		for i := 0; i < len(parseResult.Requests); i++ {
			if strings.Contains(parseResult.Requests[i].Url, "http://") == false {
				parseResult.Requests[i].Url = r.Url + parseResult.Requests[i].Url
			}
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
