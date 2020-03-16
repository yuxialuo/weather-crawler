package main

import (
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
