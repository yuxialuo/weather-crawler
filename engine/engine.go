package engine

import (
	"log"
	"strings"

	"github.com/yuxialuo/weather-crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		for i := 0; i < len(parseResult.Requests); i++ {
			if strings.Contains(parseResult.Requests[i].Url, "http://") == false {
				parseResult.Requests[i].Url = r.Url + parseResult.Requests[i].Url
			}
		}
		for i := 0; i < len(parseResult.Items); i++ {
			if r.Data != nil {
				parseResult.Items[i] = r.Data.(string) + parseResult.Items[i].(string)
			}
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
