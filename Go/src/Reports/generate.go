package Reports

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type Flight struct {
	sourceCity         string
	destinationCity    string
	sourceCountry      string
	destinationCountry string
	sourceAirport      string
	destinationAirport string
	departureDate      string
	returnDate         string
	price              int
	tripLength         int
}

const SELECT_SOURCES string = "SELECT * FROM SourceAirports;"
const SELECT_DESTINATIONS string = "SELECT * FROM DestinationAirports;"
const MIN_QUERY string = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM %s_%s Where Price = (SELECT Min(Price) FROM %s_%s WHERE DATEDIFF(ReturnDate, DepartDate) > 2) AND DATEDIFF(ReturnDate, DepartDate) > 2 limit 1;"
const REPORT_LOC string = "./../../reports/Bargains:%s.csv"
const CSV_LINE_FORMAT string = "\"%s, %s\",\"%s, %s, %s\",%s,%s,%d,%d\n"
const CSV_HEADERS string = "From,To,Leaving,Returning,Trip Length,Cost\n"
const DATE_FORMAT string = "2006-01-02"

func GenerateReport(db *sql.DB) (reportLoc string) {

	currentDate := time.Now()
	var minFlight Flight
	var minFlights []Flight
	var potentialMin Flight

	// file storing the report
	file, err := os.Create(fmt.Sprintf(REPORT_LOC, currentDate.Format(DATE_FORMAT)))
	if err != nil {
		fmt.Println("failed to open report item generate.go")
		log.Fatal(err)
	}
	defer file.Close()

	// headers in csv file
	_, err = file.WriteString(CSV_HEADERS)
	if err != nil {
		fmt.Println("failed to write headers generate.go")
		log.Fatal(err)
	}

	// getting source airports from database
	srcAirports, err := db.Query(SELECT_SOURCES)
	if err != nil {
		fmt.Println("failed to get sources generate.go")
		panic(err.Error())
	}
	defer srcAirports.Close()

	// getting destination airports from database
	destAirports, err := db.Query(SELECT_DESTINATIONS)
	if err != nil {
		fmt.Println("failed to get destinations generate.go")
		panic(err.Error())
	}
	defer destAirports.Close()


	fmt.Println("going through airports")
	// getting the cheapest flights to each destination
	for destAirports.Next() {

		// have to refresh the min flights so we don't have hangovers from previous destination
		potentialMin = Flight{}
		minFlight = Flight{}

		for srcAirports.Next() {

			var dummy string

			if err := srcAirports.Scan(&dummy, &dummy, &potentialMin.sourceAirport, &potentialMin.sourceCountry, &potentialMin.sourceCity); err != nil {
				fmt.Println("failed to scan source airports generate.go")
				panic(err.Error())
			}

			if err := destAirports.Scan(&dummy, &dummy, &potentialMin.destinationAirport, &potentialMin.destinationCountry, &potentialMin.destinationCity); err != nil {
				fmt.Println("failed to scan destinations airports generate.go")
				panic(err.Error())
			}

			err = db.QueryRow(fmt.Sprintf(MIN_QUERY, potentialMin.sourceAirport, potentialMin.destinationAirport, potentialMin.sourceAirport, potentialMin.destinationAirport)).Scan(&dummy, &dummy, &dummy, &potentialMin.departureDate, &potentialMin.returnDate, &potentialMin.price, &potentialMin.tripLength)
			if err == nil {
				// updating the local cheapest flight
				if minFlight == (Flight{}) {
					minFlight = potentialMin
				} else if potentialMin.price < minFlight.price {
					minFlight = potentialMin
				}
			}
		}

		minFlights = append(minFlights, minFlight)

		// have to reload the result set into destAirports because .Next()
		srcAirports, err = db.Query(SELECT_SOURCES)
		if err != nil {
			fmt.Println("failed to reload generate.go")
			panic(err.Error())
		}
	}

	fmt.Println("going through flights")
	// writing out the flight data to the report file
	for _, flight := range minFlights {

		_, err = file.WriteString(fmt.Sprintf(CSV_LINE_FORMAT, flight.sourceCity, flight.sourceAirport, flight.destinationCountry, flight.destinationCity, flight.destinationAirport, flight.departureDate, flight.returnDate, flight.tripLength, flight.price))

		if err != nil {
			fmt.Println("failed to write out flights to report generate.go")
			log.Fatal(err)
		}
	}
	fmt.Println("done")
	return fmt.Sprintf(REPORT_LOC, currentDate.Format(DATE_FORMAT))

}
