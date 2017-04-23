package Reports

import (
	"fmt"
	"log"
	"os"
)

const toWidth int = 275
const leavingWidth int = 120
const returningWidth int = 120
const tripLengthWidth int = 40
const costWidth int = 40
const tableWidth int = 700
const hrefLink string = "http://partners.api.skyscanner.net/apiservices/referral/v1.0/GB/GBP/en-GB/%s/%s/%s/%s?apiKey=na91261163675973"
const flightWriteError string = "failed to write out one of the flight attributes - httpEmail.go"
const tableHeadingsWriteError string = "failed to write one of the table headers - htmlEmail.go"
const htmlHeadersWriteErrors string = "failed to write html template header tabs - httpEmail.go"
const closeTableWriteError string = "failed to close table tag - httpEmail.go"
const brTagWriteError string = "failed to write <br> tag - httpEmail.go"
const closeBodyHtmlTagWriteError string = "failed to close body and html tag - httpEmail.go"
const paragraphHeaderFormat string = "<p style = \"padding: 5px 5px 5px 5px; margin: 0 0 0 0;\"><b>%s, %s</b></p>"
const paddingStyle string = "padding:0 10px 0 10px;"
const centreAlignStyle string = "text-align: center;"

func generatePrettyReport(flights []Flight, file *os.File, salt string) {

	By(b_SourceCity).Sort(flights)

	var groupedFlights [][]Flight

	groupedFlights = getGroupedFlights(flights)

	for i := 0; i < len(groupedFlights); i++ {
		By(b_TripPrice).Sort(groupedFlights[i])
	}

	writeHTMLHeadings(file)

	for i, _ := range groupedFlights {

		writeToFile(fmt.Sprintf(paragraphHeaderFormat, groupedFlights[i][0].sourceCity, groupedFlights[i][0].sourceAirport), flightWriteError, file)
		writeTableHeadings(file)

		for _, flight := range groupedFlights[i] {

			writeFlightInfo(file, flight)
		}

		writeToFile("</table>\n", closeTableWriteError, file)
		writeToFile("<br>\n", brTagWriteError, file)
	}

	writeEndStatements(file, salt);

	writeToFile("</body>\n</html>\n", closeBodyHtmlTagWriteError, file)
}

// handles the errors associated with writing to a file - seeing as i do it a lot
func writeToFile(line, errorMessage string, file *os.File) {

	_, err := file.WriteString(line)

	if err != nil {
		fmt.Println(errorMessage)
		log.Fatal(err)
	}

}

func writeEndStatements(file *os.File, salt string){

	writeToFile(fmt.Sprintf("<p>To update your preferences, click <a href = 'http://www.skytracker.co/index2.html?token=%s'>here</a></p>\n",salt),"failed to finish", file)

}

func writeHTMLHeadings(file *os.File){
		writeToFile("<!doctype html>\n", htmlHeadersWriteErrors, file)
		writeToFile("<html lang=\"en\">\n", htmlHeadersWriteErrors, file)
		writeToFile("<head>\n", htmlHeadersWriteErrors, file)
		writeToFile("<meta charset=\"UTF-8\">\n", htmlHeadersWriteErrors, file)
		writeToFile("</head>\n", htmlHeadersWriteErrors, file)
		writeToFile("<body style = \"font-family: Helvetica;\">\n", htmlHeadersWriteErrors, file)
}

func writeFlightInfo(file *os.File, flight Flight) {

	writeToFile("<tr>\n", flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s\"><a href = \"%s\">%s, %s, %s</a></td>\n", paddingStyle, fmt.Sprintf(hrefLink,flight.sourceAirport,flight.destinationAirport,flight.departureDate,flight.returnDate),flight.destinationCity, flight.destinationCountry, flight.destinationAirport), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s%s\">%s</td>\n", paddingStyle, centreAlignStyle, flight.departureDate), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s%s\">%s</td>\n", paddingStyle, centreAlignStyle, flight.returnDate), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s%s\">%d</td>\n", paddingStyle, centreAlignStyle, flight.tripLength), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s%s\">£%d</td>\n", paddingStyle, centreAlignStyle, flight.price), flightWriteError, file)
	writeToFile("</tr>\n", flightWriteError, file)

}

func writeTableHeadings(file *os.File) {
	writeToFile(fmt.Sprintf("<table style = \"width: %dpx\">\n", tableWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"%s text-align: left; min-width = %d;\" >To</th>\n", toWidth,paddingStyle,toWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"%s%s min-width = %d;\" >Leaving</th>\n",paddingStyle, centreAlignStyle, leavingWidth,leavingWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"%s%s min-width = %d;\" >Returning</th>\n",paddingStyle, centreAlignStyle, returningWidth,returningWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"%s%s min-width = %d;\" >Days</th>\n",paddingStyle, centreAlignStyle, tripLengthWidth,tripLengthWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"%s%s min-width = %d;\" >Cost</th>\n",paddingStyle, centreAlignStyle, costWidth,costWidth), tableHeadingsWriteError, file)
	writeToFile("</tr>\n", tableHeadingsWriteError, file)
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
