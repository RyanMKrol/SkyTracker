package Reports

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func GeneratePrettyReport(flights []Flight) (reportLoc string) {

	By(srcCity).Sort(flights)

	var maxFrom, maxTo, maxLeaving, maxReturning, maxTrip, maxCost = 0, 0, 0, 0, 0, 0

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

	for _ , flight := range flights {
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


		fmt.Printf("%s   %s   %s   %s   %s   %s\n", from, to, leaving, returning, tripLength, cost)
	}

	return "things"

}

// returns a string padded with spaces on the right
func padString(original, padString string, num int) string {

	return (original + strings.Repeat(padString, num))

}
