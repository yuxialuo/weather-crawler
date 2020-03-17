package parser

import (
	"regexp"

	"github.com/yuxialuo/weather-crawler/engine"
)

/*
const regionRe = `<a href="http://www.weather.com.cn/weather/[0-9]+\.shtml"[^>]*>[^<]+</a>`

func ParseRegion(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(regionRe)
	matches = re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2])))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result
}
*/
const (
	parseRegionDataRe = `(?s)class="forecastBox"(.*?)frameborder`
	regionRe          = `href="(http://www.weather.com.cn/weather/[0-9]+.shtml)"[^>]+>([^<]+)`
)

func ParseRegion(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(parseRegionDataRe)
	data := re.Find(contents)

	re = regexp.MustCompile(regionRe)
	matches := re.FindAllSubmatch(data, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseProfile,
			})
	}
	return result
}
