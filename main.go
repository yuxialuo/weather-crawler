package main

import (
	"fmt"
	"regexp"

	"github.com/yuxialuo/weather-crawler/engine"
	"github.com/yuxialuo/weather-crawler/weather/parser"
)

var (
	municipalities = []string{
		//		"bj",
		//		"tj",
		"sh",
		//		"cq",
	}
	provinces = []string{
		//"hebei",
		//"shanxi",
		"ln",
		//"hlj",
		/*		"js",
				"zj",
				"ah",
				"fj",
				"jx",
				"sd",
				"henan",
				"hubei",
				"hainan",
				"gd",
				"hunan",
				"sc",
				"gz",
				"yn",
				"shaanxi",
				"gs",
				"qh",
				"tw",
				"nmg",
				"gx",
				"xz",
				"nx",
				"xj",
				"hk",
				"mo",*/
	}
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://ln.weather.com.cn/",
		ParserFunc: parser.ParseRegionList,
	})
}

func printCountyList(contents []byte) {
	re := regexp.MustCompile(`<a href="http://www.weather.com.cn/weather/[0-9]+\.shtml"[^>]*>[^<]+</a>`)
	matches := re.FindAll(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", m)
	}
}
