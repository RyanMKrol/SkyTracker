package Reports

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func GeneratePrettyReport(flights []Flight) (reportLoc string) {

	By(b_SourceCity).Sort(flights)

	var orderedFlights [][]Flight
	var reportEntries  [][]ReportEntry

	var maxFrom, maxTo, maxLeaving, maxReturning, maxTrip, maxCost = 0, 0, 0, 0, 0, 0

	var current string
	var counter int = -1

	// creating slice of report entries
	for _ , flight := range flights {

		if flight.sourceAirport != current {
			current = flight.sourceAirport
			counter++
			orderedFlights = append(orderedFlights, []Flight{flight})
		} else {
			orderedFlights[counter] = append(orderedFlights[counter], flight)
		}

	}

	for i := 0; i < len(orderedFlights); i++ {
		By(b_TripPrice).Sort(orderedFlights[i])
	}

	// getting the padding for each column
	for _ , flight := range flights {
		var from, to, leaving, returning, tripLength, cost string

		from = fmt.Sprintf("%s, %s", flight.sourceCity, flight.sourceAirport)
		if len(from) > maxFrom {
			maxFrom = len(from)
		}

		to = fmt.Sprintf("%s, %s, %s", flight.destinationCountry, flight.destinationCity, flight.destinationAirport)
		if len(to) > maxTo {
			maxTo = len(to)
		}

		leaving = fmt.Sprintf("%s", flight.departureDate)
		if len(leaving) > maxLeaving {
			maxLeaving = len(leaving)
		}

		returning = fmt.Sprintf("%s", flight.returnDate)
		if len(returning) > maxReturning {
			maxReturning = len(returning)
		}

		tripLength = fmt.Sprintf("%d", flight.tripLength)
		if len(tripLength) > maxTrip {
			maxTrip = len(tripLength)
		}

		cost = fmt.Sprintf("%d", flight.price)
		if len(cost) > maxCost {
			maxCost = len(cost)
		}
	}

	counter = 0
	current = ""
	first := true

	// creating slice of report entries
	for _ , block := range orderedFlights {
		for _, flight := range block {

			var from, to, leaving, returning, tripLength, cost string

			from = fmt.Sprintf("%s, %s", flight.sourceCity, flight.sourceAirport)
			to = fmt.Sprintf("%s, %s, %s", flight.destinationCountry, flight.destinationCity, flight.destinationAirport)
			leaving = fmt.Sprintf("%s", flight.departureDate)
			returning = fmt.Sprintf("%s", flight.returnDate)
			tripLength = fmt.Sprintf("%d", flight.tripLength)
			cost = fmt.Sprintf("%d", flight.price)

			from       = padString(from, " ", maxFrom - utf8.RuneCountInString(from))
			to         = padString(to, " ", maxTo - utf8.RuneCountInString(to))
			leaving    = padString(leaving, " ", maxLeaving - utf8.RuneCountInString(leaving))
			returning  = padString(returning, " ", maxReturning - utf8.RuneCountInString(returning))
			tripLength = padString(tripLength, " ", maxTrip - utf8.RuneCountInString(tripLength))
			cost       = padString(cost, " ", maxCost - utf8.RuneCountInString(cost))

			entry := ReportEntry{from, to, leaving, returning, tripLength, cost}

			if first {
				reportEntries = append(reportEntries, []ReportEntry{entry})
				first = false
			} else {
				reportEntries[counter] = append(reportEntries[counter], entry)
			}
		}
		counter++
		first = true
	}

	// creating slice of report entries
	for _ , block := range reportEntries {
		for _, flight := range block {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n", flight.from ,flight.to ,flight.leaving ,flight.returning ,flight.lenth ,flight.cost )
		}
		fmt.Println()
	}

	return "things"
}


// returns a string padded with spaces on the right
func padString(original, padString string, num int) string {

	return (original + strings.Repeat(padString, num))

}
