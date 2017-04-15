package Reports

import (
	"SystemConfig"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type StandardReport struct {
	basicReportLoc       string
	beautifiedReportloc  string
}

const selectSources string = "SELECT * FROM SourceAirports;"
const selectDestinations string = "SELECT * FROM DestinationAirports;"
const minQuery string = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM %s_%s Where Price = (SELECT Min(Price) FROM %s_%s WHERE DATEDIFF(ReturnDate, DepartDate) > 2) AND DATEDIFF(ReturnDate, DepartDate) > 2 limit 1;"
const reportLoc string = "reports/Bargains:%s.csv"
const csvLineFormat string = "\"%s, %s\",\"%s, %s, %s\",%s,%s,%d,%d\n"
const csvHeaders string = "From,To,Leaving,Returning,Trip Length,Cost\n"
const dateFormat string = "2006-01-02"


func (r *StandardReport) GetBasicReport() string {
	return r.basicReportLoc
}
func (r *StandardReport) GetFormattedReport() string {
	return r.beautifiedReportloc
}

func (r *StandardReport) GenerateReport(db *sql.DB) {

	currentDate := time.Now()
	var minFlight Flight
	var minFlights []Flight
	var potentialMin Flight

	// file storing the report
	file, err := os.Create(fmt.Sprintf(fmt.Sprintf(SystemConfig.DOC_ROOT,reportLoc), currentDate.Format(dateFormat)))
	if err != nil {
		fmt.Println("failed to open report item generate.go")
		log.Fatal(err)
	}
	defer file.Close()

	// headers in csv file
	_, err = file.WriteString(csvHeaders)
	if err != nil {
		fmt.Println("failed to write headers generate.go")
		log.Fatal(err)
	}

	// getting source airports from database
	srcAirports, err := db.Query(selectSources)
	if err != nil {
		fmt.Println("failed to get sources generate.go")
		panic(err.Error())
	}
	defer srcAirports.Close()

	// getting destination airports from database
	destAirports, err := db.Query(selectDestinations)
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

			err = db.QueryRow(fmt.Sprintf(minQuery, potentialMin.sourceAirport, potentialMin.destinationAirport, potentialMin.sourceAirport, potentialMin.destinationAirport)).Scan(&dummy, &dummy, &dummy, &potentialMin.departureDate, &potentialMin.returnDate, &potentialMin.price, &potentialMin.tripLength)
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
		srcAirports, err = db.Query(selectSources)
		if err != nil {
			fmt.Println("failed to reload generate.go")
			panic(err.Error())
		}
	}

	fmt.Println("going through flights")
	// writing out the flight data to the report file
	for _, flight := range minFlights {

		_, err = file.WriteString(fmt.Sprintf(csvLineFormat, flight.sourceCity, flight.sourceAirport, flight.destinationCountry, flight.destinationCity, flight.destinationAirport, flight.departureDate, flight.returnDate, flight.tripLength, flight.price))

		if err != nil {
			fmt.Println("failed to write out flights to report generate.go")
			log.Fatal(err)
		}
	}
	fmt.Println("done")

	GeneratePrettyReport(minFlights)

	r.basicReportLoc = fmt.Sprintf(fmt.Sprintf(SystemConfig.DOC_ROOT,reportLoc), currentDate.Format(dateFormat))
	r.beautifiedReportloc = GeneratePrettyReport(minFlights)
}
