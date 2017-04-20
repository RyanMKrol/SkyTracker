package Reports

import (
	"SystemConfig"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const SELECT_USERS string = "SELECT * FROM Users;"
const SELECT_SOURCES string = "SELECT * FROM SourceAirports WHERE SrcAirportCode IN (SELECT SourceAirportCode FROM Users NATURAL JOIN UserSourceAirports WHERE UserEmailAddress = \"%s\");"
const SELECT_DESTINATIONS string = "SELECT * FROM DestinationAirports;"
const MIN_QUERY string = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM %s_%s WHERE DATEDIFF(ReturnDate, DepartDate) >= %d AND DATEDIFF(ReturnDate, DepartDate) <= %d AND Price < %d ORDER BY Price ASC limit 1;"
const REPORT_LOC string = "reports/%d_%d_%d_%s.html"
const DATE_FORMAT string = "2006-01-02"
const MAX_NUM int = 2147483647

// this is used to sync up the threads that are doing work before we continue
var wg sync.WaitGroup

func GenerateReports(db *sql.DB) []User {

	currentDate := time.Now()
	users := getUsers(db)

	for i, _ := range users {

		fmt.Println("in the for-loop")

		var filename string = fmt.Sprintf(fmt.Sprintf(SystemConfig.DOC_ROOT, REPORT_LOC), users[i].budget, users[i].tripMin, users[i].tripMax, currentDate.Format(DATE_FORMAT))

		// file doesn't exist so we need to make it ourselves
		if _, err := os.Stat(filename); os.IsNotExist(err) {

			file, err := os.Create(filename)
			if err != nil {
				fmt.Println("failed to create report item generate.go")
				log.Fatal(err)
			}

			wg.Add(1)

			// parallelising the meat of the file
			go func(u User, f *os.File) {
				minFlights := reportForUser(u, db)
				generatePrettyReport(minFlights, f)
				f.Close()
				wg.Done()
			}(users[i], file)
		}

		users[i].NiceReportLoc = filename
	}

	wg.Wait()

	return users
}

// gets the users from the database and parses their information
func getUsers(db *sql.DB) []User {

	var userArr []User

	// getting the users and their preferences from the database
	users, err := db.Query(SELECT_USERS)
	if err != nil {
		fmt.Println("failed to get users generate.go")
		panic(err.Error())
	}
	defer users.Close()

	//these are here until i get actual values

	for users.Next() {

		var dummy string
		var tempUser User = User{}
		var maybeBudget, maybeTripMin, maybeTripMax sql.NullInt64

		if err := users.Scan(&dummy, &tempUser.EmailAddress, &maybeBudget, &maybeTripMin, &maybeTripMax, &dummy); err != nil {
			fmt.Println("failed to scan users generate.go")
			panic(err.Error())
		}

		if maybeBudget.Valid {
			tempUser.budget = int(maybeBudget.Int64)
		} else {
			tempUser.budget = MAX_NUM
		}

		if maybeTripMin.Valid {
			tempUser.tripMin = int(maybeTripMin.Int64)
		} else {
			tempUser.tripMin = 0
		}

		if maybeTripMax.Valid {
			tempUser.tripMax = int(maybeTripMax.Int64)
		} else {
			tempUser.tripMax = MAX_NUM
		}

		userArr = append(userArr, tempUser)
	}

	return userArr
}

// generates a report for a specific user
func reportForUser(user User, db *sql.DB) []Flight {

	var minFlight Flight
	var minFlights []Flight
	var potentialMin Flight

	// getting source airports from database
	srcAirports, err := db.Query(fmt.Sprintf(SELECT_SOURCES,user.EmailAddress))
	if err != nil {
		fmt.Println(user.EmailAddress)
		fmt.Println(fmt.Sprintf(SELECT_SOURCES,user.EmailAddress))
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
		srcAirports, err = db.Query(fmt.Sprintf(SELECT_SOURCES,user.EmailAddress))
		if err != nil {
			fmt.Println("failed to reload generate.go")
			panic(err.Error())
		}
	}

	return minFlights
}
