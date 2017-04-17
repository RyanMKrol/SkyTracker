package Reports

import (
	"fmt"
	"strings"
	"os"
	"log"
)

func generatePrettyReport(flights []Flight, file_location string) (reportLoc string) {

	By(b_SourceCity).Sort(flights)

	formattedFileLocation := strings.Replace(file_location, ".csv", ".html", 1)

	var groupedFlights [][]Flight

	file, err := os.Create(formattedFileLocation)
	if err != nil {
		fmt.Println("failed to create report item htttpEmail.go")
		log.Fatal(err)
	}
	defer file.Close()

	writeHtmlBase(file)

	groupedFlights = getGroupedFlights(flights)

	for i := 0; i < len(groupedFlights); i++ {
		By(b_TripPrice).Sort(groupedFlights[i])
	}


	for _, block := range groupedFlights {

		writeTableHeadings(file)

		for _, flight := range block {

			_, err = file.WriteString("<tr>\n")
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;\">%s, %s</td>\n", flight.sourceCity, flight.sourceAirport ))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;\">%s, %s, %s</td>\n", flight.destinationCity, flight.destinationCountry, flight.destinationAirport))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;\">%s</td>\n", flight.departureDate))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;\">%s</td>\n", flight.returnDate))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;\">%d</td>\n", flight.tripLength))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;\">%d</td>\n", flight.price))
			errorCheck(err)
			_, err = file.WriteString("</tr>\n")
			errorCheck(err)

		}

		_, err = file.WriteString("</table>\n")
		errorCheck(err)
		_, err = file.WriteString("<br>\n")
		errorCheck(err)
	}

	_, err = file.WriteString("</body>\n</html>\n")
	errorCheck(err)

	return formattedFileLocation
}


// certain functions were getting too meaty with the constant error checking, so
//  this lightweight function will act as a general error catcher
func errorCheck(err error){
	if err != nil {
		fmt.Println("failed the error check")
		log.Fatal(err)
	}
}

func writeTableHeadings(file *os.File) {

	_, err := file.WriteString("<table>\n")
	errorCheck(err)
	_, err = file.WriteString("<tr>\n")
	errorCheck(err)
	_, err = file.WriteString("<th style = \"text-align: left; padding:0 10px 0 10px;\" >From</th>\n")
	errorCheck(err)
	_, err = file.WriteString("<th style = \"text-align: left; padding:0 10px 0 10px;\" >To</th>\n")
	errorCheck(err)
	_, err = file.WriteString("<th style = \"text-align: left; padding:0 10px 0 10px;\" >Leaving</th>\n")
	errorCheck(err)
	_, err = file.WriteString("<th style = \"text-align: left; padding:0 10px 0 10px;\" >Returning</th>\n")
	errorCheck(err)
	_, err = file.WriteString("<th style = \"text-align: left; padding:0 10px 0 10px;\" >Trip Length</th>\n")
	errorCheck(err)
	_, err = file.WriteString("<th style = \"text-align: left; padding:0 10px 0 10px;\" >Cost</th>\n")
	errorCheck(err)
	_, err = file.WriteString("</tr>\n")
	errorCheck(err)
}

func writeHtmlBase(file *os.File) {

	_, err := file.WriteString("<!doctype html>\n<html lang=\"en\">\n<head>\n<meta charset=\"UTF-8\">\n</head>\n<body>\n")
	if err != nil {
		fmt.Println("failed to write headers htttpEmail.go")
		log.Fatal(err)
	}

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
func getGroupedFlights(flights []Flight) [][]Flight {

	var counter int = -1
	var current string = ""

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
