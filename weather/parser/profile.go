package parser

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/yuxialuo/weather-crawler/engine"
	"github.com/yuxialuo/weather-crawler/model"
)

/*
<ul class="t clearfix">
<li class="sky skyid lv2 on">
<h1>16日（今天）</h1>
<big class="png40 d01"></big>
<big class="png40 n01"></big>
<p title="多云" class="wea">多云</p>
<p class="tem">
<span>7</span>/<i>-4℃</i>
</p>
<p class="win">
<em>
<span title="西南风" class="SW"></span>
<span title="西南风" class="SW"></span>
</em>
<i>5-6级转4-5级</i>
</p>
<div class="slid"></div>
</li>
<li class="sky skyid lv1">
*/

const (
	parseProfileRe = `(?s)class="sky skyid lv2 on"(.*?)class="slid"`
	titleRe        = `<p title="([^"]+)" class="wea">`
	tempratureRe   = `<span>([^<]+)</span>/<i>([^℃]+)℃</i>`
	windRe         = `<span title="([^"]+)" class="SW"></span>`
	windForceRe    = `<i>([^<]+)</i>`
)

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(parseProfileRe)
	data := re.Find(contents)
	fmt.Println(string(data))
	profile := model.Profile{}

	re = regexp.MustCompile(titleRe)
	match := re.FindSubmatch(data)
	if match != nil {
		profile.Weather = string(match[1])
	}

	re = regexp.MustCompile(tempratureRe)
	match = re.FindSubmatch(data)
	if match != nil {
		high, err := strconv.Atoi(string(match[1]))
		if err == nil {
			profile.TemperatureHigh = high
		}
		low, err := strconv.Atoi(string(match[2]))
		if err == nil {
			profile.TemperatureLow = low
		}
	}

	re = regexp.MustCompile(windRe)
	match = re.FindSubmatch(data)
	if match != nil {
		profile.Wind = string(match[1])
	}

	re = regexp.MustCompile(windForceRe)
	match = re.FindSubmatch(data)
	if match != nil {
		profile.WindForce = string(match[1])
	}

	/*	re = regexp.MustCompile(regionRe)
		matches := re.FindAllSubmatch(data, -1)
		result := engine.ParseResult{}
		for _, m := range matches {
			result.Items = append(result.Items, string(m[2]))
			result.Requests = append(
				result.Requests, engine.Request{
					Url:        string(m[1]),
					ParserFunc: engine.NilParser,
				})
		}*/
	return engine.ParseResult{}
}
