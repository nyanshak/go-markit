package markit

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type Company struct {
	Symbol				string		`json:"Symbol"`
	Name				string		`json:"Name"`
	Exchange			string		`json:"Exchange"`
}

func Lookup(queryString string) ([]Company, error) {
	if queryString == "" {
		return []Company{}, errors.New("Missing query string")
	}

	resp, err := http.Get(BaseUrl + "/Lookup/json?input=" + queryString)

	if err != nil {
		return []Company{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Company{}, err
	}

	var companies []Company
	err = json.Unmarshal(body, &companies)

	if err != nil {
		return []Company{}, err
	}

	return companies, nil
}

