package Reports

import (
	"fmt"
	"log"
	"os"
)

// html sizes
const toWidth int = 275
const leavingWidth int = 120
const returningWidth int = 120
const tripLengthWidth int = 40
const costWidth int = 40
const tableWidth int = 700

// skyscanner redirect link
const hrefLink string = "http://partners.api.skyscanner.net/apiservices/referral/v1.0/GB/GBP/en-GB/%s/%s/%s/%s?apiKey=na91261163675973"

// file write errors
const flightWriteError string = "failed to write out one of the flight attributes - httpEmail.go"
const tableHeadingsWriteError string = "failed to write one of the table headers - htmlEmail.go"
const htmlHeadersWriteErrors string = "failed to write html template header tabs - httpEmail.go"
const closeTableWriteError string = "failed to close table tag - httpEmail.go"
const brTagWriteError string = "failed to write <br> tag - httpEmail.go"
const closeBodyHtmlTagWriteError string = "failed to close body and html tag - httpEmail.go"
const closeDivWriteError string = "failed to close the div tag - httpEmail.go"
const openDivWriteError string = "failed to open the div tag - httpEmail.go"

// write formats
const paragraphHeaderFormat string = "<p style = \"padding: 5px 5px 5px 5px; margin: 0 0 0 0;\"><b>%s, %s</b></p>"

// styles
const paddingStyle string = "padding:0 10px 0 10px;"
const centreAlignStyle string = "text-align: center;"

const baseLink string = "http://www.skytracker.co?"
const unsubLink string = "http://www.skytracker.co?unsubscribe&"

// main function for generating the report itself
func generatePrettyReport(flights []Flight, file *os.File, user *User) {

	By(b_SourceCity).Sort(flights)

	var groupedFlights [][]Flight

	groupedFlights = getGroupedFlights(flights)

	if len(groupedFlights) > 0 {
		(*user).HasReport = true
	}

	for i := 0; i < len(groupedFlights); i++ {
		By(b_TripPrice).Sort(groupedFlights[i])
	}

	writeHTMLHeadings(file)

	for i, _ := range groupedFlights {

		writeToFile("<div style = \"width: 700px; margin: auto;\">\n", openDivWriteError, file)
		writeToFile(fmt.Sprintf(paragraphHeaderFormat, groupedFlights[i][0].sourceCity, groupedFlights[i][0].sourceAirport), flightWriteError, file)
		writeTableHeadings(file)

		for _, flight := range groupedFlights[i] {

			writeFlightInfo(file, flight)
		}

		writeToFile("</table>\n", closeTableWriteError, file)
		writeToFile("</div>\n", closeDivWriteError, file)
		writeToFile("<br>\n", brTagWriteError, file)
	}
	writeToFile("<div style = \"width:700px; margin: auto;\"><br>- the SkyTracker team.</div>\n", brTagWriteError, file)

	writeEndStatements(file, *user)



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

// write any statements at the end of the report
func writeEndStatements(file *os.File, user User) {

	var link = fmt.Sprintf("%stoken=%s&email=%s&tripMin=%d&tripMax=%d&budget=%d&frequency=%d", baseLink, user.salt, user.EmailAddress, user.tripMin, user.tripMax, user.budget,user.ReportFrequency)

	for _, month := range user.months {
		link += fmt.Sprintf("&month=%d", month)
	}

	for _, airport := range user.sources {
		link += fmt.Sprintf("&source=%s", airport)
	}

	var anchorTag string = fmt.Sprintf("<a style = \"color:#5C596B;font-weight: normal;\" href = '%s'>here</a>", link)

	writeToFile("<br><div style = \"width: 700px; margin: auto;\">\n", "failed to finish", file)
	writeToFile("<table style = \"width: 600px; margin: auto; border-bottom-style: solid; border-bottom-color: white; border-bottom-width: 2px;\">\n", "failed to finish", file)
	writeToFile("</table><br>\n", "failed to finish", file)
	// writeToFile("<a href = \"https://github.com/RyanMKrol/SkyTracker\"><img style = \"display: block; margin: 0 auto;\" src = \"http://skytracker.co/Images/GitHub-Mark-32px.png\"></img></a><br>\n", "failed to finish", file)
	writeToFile("<div style = \"font-size: 9pt; width: 500px; margin: auto;text-align:Center;\">\n", "failed to finish", file)
	writeToFile("<a style = \"margin: 0 auto;\" href = \"https://github.com/RyanMKrol/SkyTracker\"><img src = \"http://skytracker.co/Images/GitHub-Mark-32px.png\"></a><br><br>\n", "failed to finish", file)
	writeToFile("Our mailing address is: <br><a style = \"color:#5C596B;font-weight: normal;\" href = \"mailto:root@skytracker.co\">root@skytracker.co</a><br><br>\n", "failed to finish", file)
	writeToFile("Want to change how you receive these emails?<br>\n", "failed to finish", file)
	writeToFile(fmt.Sprintf("You can update your preferences %s or <a style = \"color:#5C596B;font-weight: normal;\"href = '%semail=%s&token=%s'>unsubscribe</a> from this list\n", anchorTag,unsubLink, user.EmailAddress, user.salt), "failed to finish", file)
	writeToFile("</div></div>\n", "failed to finish", file)
}

// writes headers to the html report
func writeHTMLHeadings(file *os.File) {
	writeToFile("<!doctype html>\n", htmlHeadersWriteErrors, file)
	writeToFile("<html lang=\"en\">\n", htmlHeadersWriteErrors, file)
	writeToFile("<head>\n", htmlHeadersWriteErrors, file)
	writeToFile("<meta charset=\"UTF-8\">\n", htmlHeadersWriteErrors, file)
	writeToFile("</head>\n", htmlHeadersWriteErrors, file)
	writeToFile("<body style = \"font-family: Georgia; background-color: #eee; color:#111111\">\n", htmlHeadersWriteErrors, file)

	writeToFile("<div style = \"width: 700px; margin: auto;\">\n", htmlHeadersWriteErrors, file)
	writeToFile("<h1 style = \"text-align: center;font-weight:lighter;\">We scour the internet for great flight deals, so you don't have to -</h1>\n", htmlHeadersWriteErrors, file)
	writeToFile("<h3>Take a look at what we found for you today</h3>\n", htmlHeadersWriteErrors, file)
	writeToFile("<table style = \"width: 600px; margin: auto; border-bottom-style: solid; border-bottom-color: white; border-bottom-width: 2px;\">\n", htmlHeadersWriteErrors, file)
	writeToFile("</table><br>\n", htmlHeadersWriteErrors, file)
	writeToFile("</div>\n", htmlHeadersWriteErrors, file)
}

// writes individual flight info to the reports
func writeFlightInfo(file *os.File, flight Flight) {

	writeToFile("<tr>\n", flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s\"><a style = \"color:#5C596B;font-weight: normal;\" href = \"%s\">%s, %s, %s</a></td>\n", paddingStyle, fmt.Sprintf(hrefLink, flight.sourceAirport, flight.destinationAirport, flight.departureDate, flight.returnDate), flight.destinationCity, flight.destinationCountry, flight.destinationAirport), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s%s\">%s</td>\n", paddingStyle, centreAlignStyle, flight.departureDate), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s%s\">%s</td>\n", paddingStyle, centreAlignStyle, flight.returnDate), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s%s\">%d</td>\n", paddingStyle, centreAlignStyle, flight.tripLength), flightWriteError, file)
	writeToFile(fmt.Sprintf("<td style=\"%s%s\">£%d</td>\n", paddingStyle, centreAlignStyle, flight.price), flightWriteError, file)
	writeToFile("</tr>\n", flightWriteError, file)

}

// write headings for each table
func writeTableHeadings(file *os.File) {
	writeToFile(fmt.Sprintf("<table style = \"width: %dpx\">\n", tableWidth), tableHeadingsWriteError, file)
	writeToFile("<tr>\n", tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th style = \"%s text-align: left; min-width = %d;\" >To</th>\n", paddingStyle, toWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th style = \"%s%s min-width = %d;\" >Leaving</th>\n", paddingStyle, centreAlignStyle, leavingWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th style = \"%s%s min-width = %d;\" >Returning</th>\n", paddingStyle, centreAlignStyle, returningWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th style = \"%s%s min-width = %d;\" >Days</th>\n", paddingStyle, centreAlignStyle, tripLengthWidth), tableHeadingsWriteError, file)
	writeToFile(fmt.Sprintf("<th style = \"%s%s min-width = %d;\" >Cost</th>\n", paddingStyle, centreAlignStyle, costWidth), tableHeadingsWriteError, file)
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
