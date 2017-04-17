package Reports

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const toWidth = 255
const leavingWidth = 90
const returningWidth = 90
const tripLengthWidth = 40
const costWidth = 40
const hrefLink = "http://partners.api.skyscanner.net/apiservices/referral/v1.0/GB/GBP/en-GB/%s/%s/%s/%s?apiKey=na91261163675973"

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

	for i, _ := range groupedFlights {

		_, err = file.WriteString(fmt.Sprintf("<p style = \"padding: 5px 5px 5px 5px; margin: 0 0 0 0;\"><b>%s, %s</b></p>", groupedFlights[i][0].sourceCity, groupedFlights[i][0].sourceAirport))

		errorCheck(err)
		writeTableHeadings(file)

		for _, flight := range groupedFlights[i] {

			_, err = file.WriteString("<tr>\n")
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;\"><a href = \"%s\">%s, %s, %s</a></td>\n", fmt.Sprintf(hrefLink,flight.sourceAirport,flight.destinationAirport,flight.departureDate,flight.returnDate),flight.destinationCity, flight.destinationCountry, flight.destinationAirport))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;text-align: center;\">%s</td>\n", flight.departureDate))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;text-align: center;\">%s</td>\n", flight.returnDate))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;text-align: center;\">%d</td>\n", flight.tripLength))
			errorCheck(err)
			_, err = file.WriteString(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;text-align: center;\">%d</td>\n", flight.price))
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
func errorCheck(err error) {
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
	_, err = file.WriteString(fmt.Sprintf("<th width = %d style = \"text-align: left; padding:0 10px 0 10px;\" >To</th>\n", toWidth))
	errorCheck(err)
	_, err = file.WriteString(fmt.Sprintf("<th width = %d style = \"text-align: center; padding:0 10px 0 10px;\" >Leaving</th>\n", leavingWidth))
	errorCheck(err)
	_, err = file.WriteString(fmt.Sprintf("<th width = %d style = \"text-align: center; padding:0 10px 0 10px;\" >Returning</th>\n", returningWidth))
	errorCheck(err)
	_, err = file.WriteString(fmt.Sprintf("<th width = %d style = \"text-align: center; padding:0 10px 0 10px;\" >Days</th>\n", tripLengthWidth))
	errorCheck(err)
	_, err = file.WriteString(fmt.Sprintf("<th width = %d style = \"text-align: center; padding:0 10px 0 10px;\" >Cost</th>\n", costWidth))
	errorCheck(err)
	_, err = file.WriteString("</tr>\n")
	errorCheck(err)
}

func writeHtmlBase(file *os.File) {

	_, err := file.WriteString("<!doctype html>\n<html lang=\"en\">\n<head>\n<meta charset=\"UTF-8\">\n</head>\n<body style = \"font-family: Helvetica;\">\n")
	if err != nil {
		fmt.Println("failed to write headers htttpEmail.go")
		log.Fatal(err)
	}

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
