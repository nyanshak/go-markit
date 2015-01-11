package markit

import (
	"time"
)

type TimePeriod string

const (
	Minute	TimePeriod	= "Minute"
	Hour				= "Hour"
	Day					= "Day"
	Week				= "Week"
	Month				= "Month"
	Quarter				= "Quarter"
	Year				= "Year"
)

type ElementType string

const (
	Price	ElementType	= "price"
	Volume				= "volume"
	SimpleMovingAverage	= "sma"
)

type InteractiveChartInput struct {
	Normalized		bool		`json:"Normalized"`		// show in price units (false) or percentages (true)
	StartDate		time.Time	`json:"StartDate"`
	EndDate			time.Time	`json:"EndDate"`
	EndOffsetDays	int32		`json:"EndOffsetDays"`
	NumberOfDays	int32		`json:"NumberOfDays"`
	DataPeriod		TimePeriod	`json:"DataPeriod"`
	DataInterval	int32		`json:"DataInterval"`	// for intraday data, specifies # of periods btwn data points
	LabelPeriod		TimePeriod	`json:"LabelPeriod"`
	LabelInterval	int32		`json:"LabelInterval"`	// how many label periods to skip between labels, 1 is safe default
	Elements		[]Element	`json:"Elements"`		// array of 1 or more elements
}

type Element struct {
	Symbol		string		`json:"Symbol"`
	Type		ElementType	`json:"Type"`
	Params		[]string	`json:"Params"`
}

func (self Element) setParams() {
	switch self.Type {
	case Price:
		self.Params = []string{"ohlc"}
		break
	case Volume:
		break
	case SimpleMovingAverage:
		break
	}
}

type InteractiveChartData struct {
	Labels		*Labels			`json:"Labels"`
	Positions	[]float64		`json:"Positions"`
	Dates		[]time.Time		`json:"Dates"`
	Elements	[]ElementData	`json:"Elements"`
}

// Note: Labels currently incomplete, need further testing
type Labels struct {
	Dates		[]*time.Time    `json:"dates"`
	Pos			[]*string		`json:"pos"`
	Priorities	[]*string		`json:"priorities"`
	Text		[]*string		`json:"text"`
	utcDates	[]*time.Time	`json:"utcDates"`
}

type ElementData struct {
	Currency	string		`json:"Currency"`
	TimeStamp	time.Time	`json:"TimeStamp"`
	Symbol		string		`json:"Symbol"`
	Type		ElementType	`json:"Type"`
	DataSeries	DataSeries	`json:"DataSeries"`
}

type DataSeries struct {
	Open				*DataSeriesObject	`json:"open"`
	Close				*DataSeriesObject	`json:"close"`
	High				*DataSeriesObject	`json:"high"`
	Low					*DataSeriesObject	`json:"low"`
	Volume				*DataSeriesObject	`json:"volume"`
	SimpleMovingAverage	*DataSeriesObject	`json:"sma"`
}

type DataSeriesObject struct {
	Min		float64		`json:"min"`
	Max		float64		`json:"max"`
	Values	[]float64	`json:"values"`
}
