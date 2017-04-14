package Reports

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func GeneratePrettyReport(flights []Flight) (reportLoc string) {

	By(b_SourceCity).Sort(flights)

	var groupedFlights [][]Flight
	var formattedEntries [][]ReportEntry

	groupedFlights = getOrderedFlights(flights)

	for i := 0; i < len(groupedFlights); i++ {
		By(b_TripPrice).Sort(groupedFlights[i])
	}

	formattedEntries = getFormattedEntries(groupedFlights, flights)

	// creating slice of report entries
	for _, block := range formattedEntries {
		for _, flight := range block {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n", flight.from, flight.to, flight.leaving, flight.returning, flight.lenth, flight.cost)
		}
		fmt.Println()
	}

	return "things"
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// returns a string padded with spaces on the right
func padString(original, padString string, num int) string {
	return (original + strings.Repeat(padString, num))

}

// returns flights ordered into slices based on source city
func getOrderedFlights(flights []Flight) [][]Flight {

	var counter int = -1
	var current string

	var flightBlocks [][]Flight

	// creating slice of report entries
	for _, flight := range flights {

		if flight.sourceAirport != current {
			current = flight.sourceAirport
			counter++
			flightBlocks = append(flightBlocks, []Flight{flight})
		} else {
			flightBlocks[counter] = append(flightBlocks[counter], flight)
		}
	}

	return flightBlocks
}

// returns ReportEntries that contain formatted flights
func getFormattedEntries(groupedFlights [][]Flight, flights []Flight) [][]ReportEntry {

	var reports [][]ReportEntry

	var maxFrom, maxTo, maxLeaving, maxReturning, maxTrip, maxCost = 0, 0, 0, 0, 0, 0

	var counter = 0
	var first bool = true

	// gets the padding value for each column
	for _, flight := range flights {
		var from, to, leaving, returning, tripLength, cost string

		from = fmt.Sprintf("%s, %s", flight.sourceCity, flight.sourceAirport)
		if len(from) > maxFrom {
			maxFrom = len(from)
		}

		to = fmt.Sprintf("%s, %s, %s", flight.destinationCountry, flight.destinationCity, flight.destinationAirport)
		maxTo = max(maxTo, len(to))

		leaving = fmt.Sprintf("%s", flight.departureDate)
		maxLeaving = max(maxLeaving, len(leaving))

		returning = fmt.Sprintf("%s", flight.returnDate)
		maxReturning = max(maxReturning, len(returning))

		tripLength = fmt.Sprintf("%d", flight.tripLength)
		maxTrip = max(maxTrip, len(tripLength))

		cost = fmt.Sprintf("%d", flight.price)
		maxCost = max(maxCost, len(cost))
	}

	// does the actual formatting based on max values for each category
	for _, block := range groupedFlights {
		for _, flight := range block {

			var from, to, leaving, returning, tripLength, cost string

			from = fmt.Sprintf("%s, %s", flight.sourceCity, flight.sourceAirport)
			to = fmt.Sprintf("%s, %s, %s", flight.destinationCountry, flight.destinationCity, flight.destinationAirport)
			leaving = fmt.Sprintf("%s", flight.departureDate)
			returning = fmt.Sprintf("%s", flight.returnDate)
			tripLength = fmt.Sprintf("%d", flight.tripLength)
			cost = fmt.Sprintf("%d", flight.price)

			from = padString(from, " ", maxFrom-utf8.RuneCountInString(from))
			to = padString(to, " ", maxTo-utf8.RuneCountInString(to))
			leaving = padString(leaving, " ", maxLeaving-utf8.RuneCountInString(leaving))
			returning = padString(returning, " ", maxReturning-utf8.RuneCountInString(returning))
			tripLength = padString(tripLength, " ", maxTrip-utf8.RuneCountInString(tripLength))
			cost = padString(cost, " ", maxCost-utf8.RuneCountInString(cost))

			entry := ReportEntry{from, to, leaving, returning, tripLength, cost}

			if first {
				reports = append(reports, []ReportEntry{entry})
				first = false
			} else {
				reports[counter] = append(reports[counter], entry)
			}
		}
		counter++
		first = true
	}

	return reports
}
