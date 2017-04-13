// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

// this is purely used to compare the type of something to map[string]interface{}
var typeComparator map[string]interface{}

const DATES_ARRAY_ID string = "Dates"
const MIN_PRICE_ID string = "MinPrice"
const QUERY_FORMAT string = "INSERT INTO %s_%s (SourcePort, DestPort, DepartDate, ReturnDate, Price) VALUES ('%s','%s','%s-%02d', '%s-%02d', %d);\n"
const DATE_FORMAT string = "2006-01"

// decodes the data into .sql files
func Decode(data []byte, src, dest string, departDate, returnDate time.Time) {

	var outboundDay = 0
	var inboundDay = 0
	var f interface{}

	departFormatted := departDate.Format(DATE_FORMAT)
	returnFormatted := returnDate.Format(DATE_FORMAT)

	// opening the file for writing, with append flag
	file, err := os.OpenFile(fmt.Sprintf(FILE_LOC, src, dest), os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("failed to open file for writing decode.go")
		log.Fatal(err)
	}
	defer file.Close()

	// parsing the JSON into a blank interface
	err = json.Unmarshal(data, &f)
	if err != nil {
		fmt.Println("failed to get json decode.go")
		log.Fatal(err)
	}

	// using type assertions to access the underlying data
	allData := f.(map[string]interface{})
	datesData := allData[DATES_ARRAY_ID].([]interface{})

	for _, dates := range datesData {

		//accessing the arrays of data
		dateQuotes := dates.([]interface{})

		for _, quotes := range dateQuotes {

			//if we see a non-null entry with the type matching that of a price quote, add it to the file
			if reflect.TypeOf(quotes) == reflect.TypeOf(typeComparator) {

				specificQuote := quotes.(map[string]interface{})
				if specificQuote[MIN_PRICE_ID] != nil {

					price := int(specificQuote[MIN_PRICE_ID].(float64))

					_, err := file.WriteString(fmt.Sprintf(QUERY_FORMAT, src, dest, src, dest, departFormatted, outboundDay, returnFormatted, inboundDay, price))
					if err != nil {
						fmt.Println("failed to write string in file decode.go")
						log.Fatal(err)
					}
				}
			}

			outboundDay++
		}

		inboundDay++
		outboundDay = 0
	}

	return
}
