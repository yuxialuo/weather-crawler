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

	result := ParseRegionList(contents)

	const resultSize = 14
	expectedUrls := []string{
		"shenyang/index.shtml", "dalian/index.shtml", "anshan/index.shtml",
	}
	expectedRegions := []string{
		"辽宁沈阳", "辽宁大连", "辽宁鞍山",
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

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d",
			resultSize, len(result.Items))
	}
	for i, region := range expectedRegions {
		if result.Items[i].(string) != region {
			t.Errorf("expected url #%d: %s; but was %s",
				i, region, result.Items[i].(string))
		}
	}
}
