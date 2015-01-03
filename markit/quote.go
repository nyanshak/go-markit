package markit

import (
	"net/http"
	"io/ioutil"
)

type Quote struct {
	Name				string		`json:"Name"`
	Symbol				string		`json:"Symbol"`
	LastPrice			float64		`json:"LastPrice"`
	Change				float64		`json:"Change"`
	ChangePercent		float64		`json:"ChangePercent"`
	Timestamp			string		`json:"Timestamp"`			// TODO: find an appropriate date representation
	MSDate				string		`json:"MSDate"`				// TODO: find an appropriate date representation
	MarketCap			int			`json:"MarketCap"`
	Volume				int			`json:"Volume"`
	ChangeYTD			float64		`json:"ChangeYTD"`
	ChangePercentYTD	float64		`json:"ChangePercentYTD"`
	High				float64		`json:"High"`
	Low					float64		`json:"Low"`
	Open				float64		`json:"Open"`
}

func GetQuote(symbol string) (Quote, error) {
	var q Quote
	resp, err := http.Get(BaseUrl + "/Quote/json?symbol=" + symbol)

	if err != nil {
		return q, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return q, err
	}

	
	return q, nil
}

