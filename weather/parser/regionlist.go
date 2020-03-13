package parser

import (
	"regexp"

	"github.com/yuxialuo/weather-crawler/engine"
)

const (
	provinceDataRe = `(?s)class="navbox"(.*?)</span>`
	provinceNameRe = `/weather.shtml">(.*?)首页`
	regionDataRe   = `(?s)<span>(.*?)</span>`
	regionListRe   = `<a href="/([a-z]+/index.shtml)">([^<]+)`
)

func ParseRegionList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(provinceDataRe)
	data := re.Find(contents)

	re = regexp.MustCompile(provinceNameRe)
	matches := re.FindAllSubmatch(data, -1)
	province := []byte{}
	for _, m := range matches {
		province = m[1]
	}

	re = regexp.MustCompile(regionDataRe)
	data = re.Find(contents)

	re = regexp.MustCompile(regionListRe)
	matches = re.FindAllSubmatch(data, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(append(province, m[2]...)))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result
}
