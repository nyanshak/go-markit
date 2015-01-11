package markit

type InteractiveChartData struct {
	Labels		Labels			`json:"Labels"`
	Positions	[]float64		`json:"Positions"`
	Dates		[]string		`json:"Dates"`			// TODO: change to time
	Elements	[]ElementData	`json:"Elements"`
}

type Labels struct {
	// TODO: Fill label with proper information
}

type ElementData struct {
	Currency	string		`json:"Currency"`
	TimeStamp	string		`json:"TimeStamp"`			// TODO: change to time
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
