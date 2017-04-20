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
// const SELECT_USERS string = "SELECT * FROM UserTravelMonths WHERE UserID = (SELECT UserID FROM Users WHERE UserEmailAddress = %s) ORDER BY TravelMonth ASC;"
const SELECT_TRAVEL_MONTHS string = "SELECT * FROM UserTravelMonths WHERE UserID = (SELECT UserID FROM Users WHERE UserEmailAddress = \"%s\") ORDER BY TravelMonth ASC;"
const SELECT_USERS string = "SELECT * FROM Users;"
const SELECT_SOURCES string = "SELECT * FROM SourceAirports WHERE SrcAirportCode IN (SELECT SourceAirportCode FROM Users NATURAL JOIN UserSourceAirports WHERE UserEmailAddress = \"%s\");"
const SELECT_DESTINATIONS string = "SELECT * FROM DestinationAirports;"
const MIN_QUERY string = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM %s_%s WHERE DATEDIFF(ReturnDate, DepartDate) >= %d AND DATEDIFF(ReturnDate, DepartDate) <= %d AND Price < %d ORDER BY Price ASC limit 1;"
const REPORT_LOC string = "reports/%d_%d_%d_%s.html"
const DATE_FORMAT string = "2006-01-02"
const MAX_NUM int = 2147483647
const MONTHS_IN_YEAR int = 12

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
				intervals := intervalBuilder(u,db)
				minFlights := reportForUser(u, db, intervals)
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

// builds up the intervals to search over that a user will have specified
func intervalBuilder(user User, db *sql.DB) (intervals []Interval) {

	months, err := db.Query(fmt.Sprintf(SELECT_TRAVEL_MONTHS, user.EmailAddress))
	if err != nil {
		fmt.Println("failed to get user months generate.go")
		panic(err.Error())
	}
	defer months.Close()

	var monthArr [MONTHS_IN_YEAR]bool
	for months.Next() {

		var dummy string
		var tempInt int
		if err := months.Scan(&dummy, &dummy, &tempInt); err != nil {
			fmt.Println("failed to scan user travel month generate.go")
			panic(err.Error())
		}
		monthArr[tempInt-1] = true
	}

	var allTrue = true
	var firstFalse int

	// finds the first false if one is present
	for i, m := range monthArr {
		if !m {
			if allTrue {
				firstFalse = i
			}
			allTrue = false
		}
	}

	// if they're all true then we can search the whole year using one interval
	if allTrue {
		intervals = append(intervals, Interval{int(time.Now().Month()),int(time.Now().Month())-1, int(time.Now().Year()), int(time.Now().Year())+1})
		return
	}

	// if not we need to find where the first 'true' after a false is
	var trueStarts int
	for i := firstFalse+1; i < firstFalse + MONTHS_IN_YEAR; i++ {
		if monthArr[i] {
			trueStarts = i%MONTHS_IN_YEAR
			break
		}
	}

	var inInterval bool = true
	var tempInterval Interval
	tempInterval.StartMonth = trueStarts + 1

	// generates the month intervals
	for i := trueStarts; i < trueStarts + MONTHS_IN_YEAR; i++ {
		if monthArr[i%MONTHS_IN_YEAR] && inInterval {

			tempInterval.EndMonth = (i%MONTHS_IN_YEAR)+1

		} else if monthArr[i%MONTHS_IN_YEAR] {

			tempInterval.StartMonth = (i%MONTHS_IN_YEAR)+1
			tempInterval.EndMonth = (i%MONTHS_IN_YEAR)+1
			inInterval = true

		} else if inInterval {

			intervals = append(intervals, tempInterval)
			inInterval = false

		}
	}

	// have to set the year as well as the month
	for i, _ := range intervals {
		intervals[i].StartYear = int(time.Now().Year())
		if intervals[i].StartMonth < intervals[i].EndMonth {
			intervals[i].EndYear = intervals[i].StartYear
		} else {
			intervals[i].EndYear = intervals[i].StartYear + 1
		}
	}

	fmt.Println(intervals)
	return
}

// generates a report for a specific user
func reportForUser(user User, db *sql.DB, intervals []Interval) []Flight {

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
