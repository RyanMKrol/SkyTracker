package Reports

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const toWidth int = 255
const leavingWidth int = 120
const returningWidth int = 120
const tripLengthWidth int = 40
const costWidth int = 40
const hrefLink string = "http://partners.api.skyscanner.net/apiservices/referral/v1.0/GB/GBP/en-GB/%s/%s/%s/%s?apiKey=na91261163675973"
const flightWriteError string = "failed to write out one of the flight attributes - httpEmail.go"
const tableHeadingsWriteError string = "failed to write one of the table headers - htmlEmail.go"

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

	groupedFlights = getGroupedFlights(flights)

	for i := 0; i < len(groupedFlights); i++ {
		By(b_TripPrice).Sort(groupedFlights[i])
	}

	writeToFile("<!doctype html>\n<html lang=\"en\">\n<head>\n<meta charset=\"UTF-8\">\n</head>\n<body style = \"font-family: Helvetica;\">\n", "failed to write html template - httpEmail.go", file)

	for i, _ := range groupedFlights {

		writeToFile(fmt.Sprintf("<p style = \"padding: 5px 5px 5px 5px; margin: 0 0 0 0;\"><b>%s, %s</b></p>", groupedFlights[i][0].sourceCity, groupedFlights[i][0].sourceAirport), flightWriteError, file)
		writeTableHeadings(file)

		for _, flight := range groupedFlights[i] {

			writeFlightInfo(file, flight)
		}

		writeToFile("</table>\n", "failed to close table tag - httpEmail.go", file)
		writeToFile("<br>\n", "failed to write <br> tag - httpEmail.go", file)
	}

	writeToFile("</body>\n</html>\n", "failed to close body and html tag - httpEmail.go", file)

	return formattedFileLocation
}

// handles the errors associated with writing to a file - seeing as i do it a lot
func writeToFile(line, errorMessage string, file *os.File) {

	_, err := file.WriteString(line)

	if err != nil {
		fmt.Println(errorMessage)
		log.Fatal(err)
	}

}

func writeFlightInfo(file *os.File, flight Flight) {

	writeToFile("<tr>\n", flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;\"><a href = \"%s\">%s, %s, %s</a></td>\n", fmt.Sprintf(hrefLink,flight.sourceAirport,flight.destinationAirport,flight.departureDate,flight.returnDate),flight.destinationCity, flight.destinationCountry, flight.destinationAirport), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;text-align: center;\">%s</td>\n", flight.departureDate), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;text-align: center;\">%s</td>\n", flight.returnDate), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;text-align: center;\">%d</td>\n", flight.tripLength), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"padding:0 10px 0 10px;text-align: center;\">%d</td>\n", flight.price), flightWriteError, file)
	writeToFile("</tr>\n", flightWriteError, file)

}


func writeTableHeadings(file *os.File) {
	writeToFile("<table style = \"width: 700px\">\n", tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"text-align: left; padding:0 10px 0 10px; min-width = %d;\" >To</th>\n", toWidth,toWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"text-align: center; padding:0 10px 0 10px; min-width = %d;\" >Leaving</th>\n", leavingWidth,leavingWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"text-align: center; padding:0 10px 0 10px; min-width = %d;\" >Returning</th>\n", returningWidth,returningWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"text-align: center; padding:0 10px 0 10px; min-width = %d;\" >Days</th>\n", tripLengthWidth,tripLengthWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th width = %d style = \"text-align: center; padding:0 10px 0 10px; min-width = %d;\" >Cost</th>\n", costWidth,costWidth), tableHeadingsWriteError, file)
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
