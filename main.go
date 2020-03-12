package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	municipalities = []string{
		"bj",
		/*		"tj",
				"sh",
				"cq",*/
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
	for _, v := range municipalities {
		url := fmt.Sprintf("http://%s.weather.com.cn/", v)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error: status code",
				resp.StatusCode)
			continue
		}
		/*all, err := bufio.NewReader(resp.Body).Peek(100)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		} else {

		}*/
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		printCityList(all)
	}
}

func printCityList(contents []byte) {
	//fmt.Println(string(contents))
	//fmt.Println(len(contents))

	re := regexp.MustCompile(`(?s)class="navbox"(.*?)</span>`)
	matches := re.Find(contents)

	fmt.Println(string(matches))

	/*
		re := regexp.MustCompile(`<a href="http://www.weather.com.cn/weather/[0-9]+\.shtml"[^>]*>[^<]+</a>`)
		matches := re.FindAll(contents, -1)
		for _, m := range matches {
			fmt.Printf("%s\n", m)
		}
	*/

}
