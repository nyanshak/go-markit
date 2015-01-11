package markit

import (
	"time"
)

type InteractiveChartData struct {
	Labels		*Labels			`json:"Labels"`
	Positions	[]float64		`json:"Positions"`
	Dates		[]time.Time		`json:"Dates"`
	Elements	[]ElementData	`json:"Elements"`
}

// Note: Labels currently incomplete, need further testing
type Labels struct {
	Dates		[]time.Time     `json:"dates"`
	Pos			[]string		`json:"pos"`
	Priorities	[]string		`json:"priorities"`
	Text		[]string		`json:"text"`
	utcDates	[]time.Time		`json:"utcDates"`
}

type ElementData struct {
	Currency	string		`json:"Currency"`
	TimeStamp	time.Time	`json:"TimeStamp"`
	Symbol		string		`json:"Symbol"`
	Type		string		`json:"Price"`
	DataSeries	DataSeries	`json:"DataSeries"`
}

type DataSeries struct {
	Open	*DataSeriesObject	`json:"open"`
	Close	*DataSeriesObject	`json:"close"`
	Volume	*DataSeriesObject	`json:"volume"`
}

type DataSeriesObject struct {
	Min		float64		`json:"min"`
	Max		float64		`json:"max"`
	Values	[]float64	`json:"values"`
}
