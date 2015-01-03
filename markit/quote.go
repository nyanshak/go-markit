package markit

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type Quote struct {
	Name				string		`json:"Name"`
	Symbol				string		`json:"Symbol"`
	LastPrice			float64		`json:"LastPrice"`
	Change				float64		`json:"Change"`
	ChangePercent		float64		`json:"ChangePercent"`
	Timestamp			string		`json:"Timestamp"`			// TODO: find an appropriate date representation
	MSDate				float64		`json:"MSDate"`				// TODO: find an appropriate date representation
	MarketCap			int			`json:"MarketCap"`
	Volume				int			`json:"Volume"`
	ChangeYTD			float64		`json:"ChangeYTD"`
	ChangePercentYTD	float64		`json:"ChangePercentYTD"`
	High				float64		`json:"High"`
	Low					float64		`json:"Low"`
	Open				float64		`json:"Open"`
	Message				string		`json:"Message"`
	Status				string		`json:"Status"`
}

func GetQuote(symbol string) (*Quote, error) {
	resp, err := http.Get(BaseUrl + "/Quote/json?symbol=" + symbol)

	if err != nil {
		return new(Quote), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return new(Quote), err
	}

	var q Quote
	err = json.Unmarshal(body, &q)

	if err != nil {
		return new(Quote), err
	} else if q.Message != "" {
		return new(Quote), errors.New(q.Message)
	} else if q.Status != "SUCCESS" {
		return new(Quote), errors.New(q.Status)
	}

	return &q, nil
}

