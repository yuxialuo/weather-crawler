package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("regionlist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseRegionList(contents, "", "")

	const resultSize = 14
	expectedUrls := []string{
		"shenyang/index.shtml", "dalian/index.shtml", "anshan/index.shtml",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d",
			resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				i, url, result.Requests[i].Url)
		}
	}
}
