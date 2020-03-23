package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/olivere/elastic"
	"github.com/yuxialuo/weather-crawler/engine"
	"github.com/yuxialuo/weather-crawler/model"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://",
		Type: "china",
		Id:   "10001",
		Payload: model.Profile{
			Weather:         "晴",
			TemperatureHigh: 13,
			TemperatureLow:  2,
			Wind:            "西南风",
			WindForce:       "4-5级",
		},
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "weather_test"
	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(
		actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v expected %v",
			actual, expected)
	}
}
