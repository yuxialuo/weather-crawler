package parser

import (
	"io/ioutil"
	"testing"

	"github.com/yuxialuo/weather-crawler/engine"
	"github.com/yuxialuo/weather-crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.shtml")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://", "")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v",
			result.Items)
	}
	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://",
		Type: "china",
		Id:   "",
		Payload: model.Profile{
			Weather:         "晴",
			TemperatureHigh: 13,
			TemperatureLow:  2,
			Wind:            "西北风",
			WindForce:       "4-5级",
		},
	}

	if actual != expected {
		t.Errorf("expected %+v; but was %+v",
			expected, actual)
	}
}
