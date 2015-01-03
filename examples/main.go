package main

import (
	"github.com/nyanshak/go-markit/markit"
	"fmt"
	"log"
)

func main() {
	quote, err := markit.GetQuote("aapl")
	
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(quote)

	companies, err := markit.Lookup("aapl")

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(companies)
}
