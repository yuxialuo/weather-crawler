package main

import (
	"fmt"

	"github.com/yuxialuo/weather-crawler/engine"
	"github.com/yuxialuo/weather-crawler/persist"
	"github.com/yuxialuo/weather-crawler/scheduler"
	"github.com/yuxialuo/weather-crawler/weather/parser"
)

var (
	municipalities = []string{
		"bj",
		"tj",
		"sh",
		"cq",
	}
	provinces = []string{
		"hebei",
		"shanxi",
		"ln",
		"hlj",
		"js",
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
		"mo",
	}
)

func main() {
	fmt.Println("weather crawler service v0.1.1 is running......")
	itemChan, err := persist.ItemSaver("weather_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	requests := []engine.Request{}
	/*
		for _, v := range provinces {
			r := engine.Request{
				Url:        fmt.Sprintf("http://%s.weather.com.cn/", v),
				ParserFunc: parser.ParseRegionList,
			}
			requests = append(requests, r)
		}*/
	for _, v := range municipalities {
		r := engine.Request{
			Url:        fmt.Sprintf("http://%s.weather.com.cn/", v),
			ParserFunc: parser.RegionParser(""),
		}
		requests = append(requests, r)
	}

	e.Run(requests...)
}
