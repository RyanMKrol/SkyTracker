package Reports

import (
	"SystemConfig"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

const SELECT_USERS string = "SELECT * FROM Users;"
const SELECT_SOURCES string = "SELECT * FROM SourceAirports;"
const SELECT_DESTINATIONS string = "SELECT * FROM DestinationAirports;"
const MIN_QUERY string = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM %s_%s WHERE DATEDIFF(ReturnDate, DepartDate) >= %d AND DATEDIFF(ReturnDate, DepartDate) <= %d AND Price < %d ORDER BY Price ASC limit 1;"
const REPORT_LOC string = "reports/%d_%d_%d_%s.csv"
const CSV_LINE_FORMAT string = "\"%s, %s\",\"%s, %s, %s\",%s,%s,%d,%d\n"
const CSV_HEADERS string = "From,To,Leaving,Returning,Trip Length,Cost\n"
const DATE_FORMAT string = "2006-01-02"
const MAX_PRICE int = 2147483647

func GenerateReport(db *sql.DB) []User {

	currentDate := time.Now()

	users := getUsers(db)
	for _, user := range users {

		var filename string = fmt.Sprintf(fmt.Sprintf(SystemConfig.DOC_ROOT,REPORT_LOC), user.budget, user.tripMin, user.tripMax, currentDate.Format(DATE_FORMAT))

		// file doesn't exist so we need to make it ourselves
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			file, err := os.Create(filename)

			if err != nil {
				fmt.Println("failed to create report item generate.go")
				log.Fatal(err)
			}

			writeHeaders(file)
			reportForUser(user, db, file)
			file.Close()
		}
		user.reportLoc = filename
	}

	return users
}

func getUsers(db *sql.DB) []User{

	var userArr []User

	// getting the users and their preferences from the database
	users, err := db.Query(SELECT_USERS)
	if err != nil {
		fmt.Println("failed to get users generate.go")
		panic(err.Error())
	}
	defer users.Close()

	var num = 100
	var lil = 7
	var big = 14

	for users.Next() {

		var dummy string
		var tempUser User = User{}

		if err := users.Scan(&dummy, &tempUser.emailAddress); err != nil {
			fmt.Println("failed to scan users generate.go")
			panic(err.Error())
		}

		tempUser.budget = num
		tempUser.tripMin = lil
		tempUser.tripMax = big

		num++
		lil++
		big++

		userArr = append(userArr, tempUser)
	}

	return userArr
}

// writes headers to the csv file
func writeHeaders(file *os.File) {

	_, err := file.WriteString(CSV_HEADERS)
	if err != nil {
		fmt.Println("failed to write headers generate.go")
		log.Fatal(err)
	}
}

func reportForUser(user User, db *sql.DB, file *os.File) {

	var minFlight Flight
	var minFlights []Flight
	var potentialMin Flight

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

			err = db.QueryRow(fmt.Sprintf(MIN_QUERY, potentialMin.sourceAirport, potentialMin.destinationAirport, user.tripMin, user.tripMax, user.budget)).Scan(&dummy, &dummy, &dummy, &potentialMin.departureDate, &potentialMin.returnDate, &potentialMin.price, &potentialMin.tripLength)
			if err == nil {
				// updating the local cheapest flight
				if minFlight == (Flight{}) {
					minFlight = potentialMin
				} else if potentialMin.price < minFlight.price {
					minFlight = potentialMin
				}
			} else {
			}
		}
		if minFlight != (Flight{}) {
			minFlights = append(minFlights, minFlight)
		}

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
}
