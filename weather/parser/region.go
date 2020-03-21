package parser

import (
	"regexp"

	"github.com/yuxialuo/weather-crawler/engine"
)

const (
	parseRegionDataRe = `(?s)class="forecastBox"(.*?)frameborder`
	regionRe          = `href="(http://www.weather.com.cn/weather/[0-9]+.shtml)"[^>]+>([^<]+)`
)

func ParseRegion(contents []byte, url, name string) engine.ParseResult {
	re := regexp.MustCompile(parseRegionDataRe)
	data := re.Find(contents)

	re = regexp.MustCompile(regionRe)
	matches := re.FindAllSubmatch(data, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		county := string(m[2])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        url,
				ParserFunc: ProfileParser(name + county),
			})
	}
	return result
}

func RegionParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseRegion(c, url, name)
	}
}
