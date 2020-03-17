package parser

import (
	"io/ioutil"
	"testing"

	"github.com/yuxialuo/weather-crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.shtml")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents)

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v",
			result.Items...)
	}
	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Weather:         "晴",
		TemperatureHigh: 13,
		TemperatureLow:  2,
		Wind:            "西南风",
		WindForce:       "4-5级",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v",
			expected, profile)
	}
}
