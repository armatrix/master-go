package main

//go:generate go run main.go

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/wcharczuk/go-chart/v2"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	log := logger.Sugar()

	url := "https://api.github.com/search/repositories?q=language:go&sort=stars"
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("failed to get, err: %v.", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("failed to read from response body, err: %v.", err)
		os.Exit(1)
	}

	var summary Summary
	err = json.Unmarshal(b, &summary)
	if err != nil {
		log.Errorf("failed to unmarshal, err: %v.", err)
		os.Exit(1)
	}
	var charts []chart.Value
	for _, item := range summary.Items {
		var chart chart.Value
		chart.Label = item.Name
		chart.Value = float64(item.StargazersCount)
		charts = append(charts, chart)
	}

	graph := chart.BarChart{
		XAxis: chart.Style{
			TextWrap: 1,
		},
		Title: "go",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 40,
		Bars:     charts,
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
