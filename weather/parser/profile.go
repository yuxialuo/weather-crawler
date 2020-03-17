package parser

import (
	"regexp"
	"strconv"

	"github.com/yuxialuo/weather-crawler/engine"
	"github.com/yuxialuo/weather-crawler/model"
)

//`(?s)class="navbox"(.*?)</span>`
var parseProfileRe = regexp.MustCompile(`(?s)（今天）(.*?)（明天）`)
var titleRe = regexp.MustCompile(`class="wea">([^<]+)</p>`)
var tempratureHighRe = regexp.MustCompile(`<span>([^<]+)</span>/<i>[^℃]+℃</i>`)
var tempratureLowRe = regexp.MustCompile(`<span>[^<]+</span>/<i>([^℃]+)℃</i>`)
var windRe = regexp.MustCompile(`<span\stitle="([^"]+)"\sclass`)
var windForceRe = regexp.MustCompile(`</em>[\s]+<i>([^\s]+)</i>`)

func ParseProfile(contents []byte) engine.ParseResult {
	data := parseProfileRe.Find(contents)
	profile := model.Profile{}

	profile.Weather = extractString(data, titleRe)
	profile.TemperatureHigh = extractInt(data, tempratureHighRe)
	profile.TemperatureLow = extractInt(data, tempratureLowRe)
	profile.Wind = extractString(data, windRe)
	profile.WindForce = extractString(data, windForceRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	data := extractString(contents, re)
	n, _ := strconv.Atoi(data)
	return n
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
