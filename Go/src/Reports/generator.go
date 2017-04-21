package Reports

import (
	"SystemConfig"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"errors"
)

// for databaser interaction
const SELECT_TRAVEL_MONTHS string = "SELECT * FROM UserTravelMonths WHERE UserID = (SELECT UserID FROM Users WHERE UserEmailAddress = \"%s\") ORDER BY TravelMonth ASC;"
const SELECT_USERS string = "SELECT * FROM Users;"
const SELECT_SOURCES string = "SELECT * FROM SourceAirports WHERE SrcAirportCode IN (SELECT SourceAirportCode FROM Users NATURAL JOIN UserSourceAirports WHERE UserEmailAddress = \"%s\");"
const SELECT_DESTINATIONS string = "SELECT * FROM DestinationAirports;"
const MIN_QUERY string = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM %s_%s WHERE DATEDIFF(ReturnDate, DepartDate) >= %d AND DATEDIFF(ReturnDate, DepartDate) <= %d AND Price < %d %s ORDER BY Price ASC limit 1;"

// for the report writing
const REPORT_LOC string = "reports/%d_%d_%d_%s.html"
const DATE_FORMAT string = "2006-01-02"

// for the query builder
const monthGreater string = "(Month(DepartDate) >= %d AND "
const monthLesser string = "Month(ReturnDate) <= %d  AND "
const yearGreater string = "Year(DepartDate) >= %d  AND "
const yearLesser string = "Year(ReturnDate) <= %d)"
const orConnective string = " OR "
const andConnective string = " AND "

// used throughout
const MAX_NUM int = 2147483647
const MONTHS_IN_YEAR int = 12

// this is used to sync up the threads that are doing work before we continue
var wg sync.WaitGroup

// generates reports for each user
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

		users[i].ReportLoc = filename
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

	// if they're all true we don't need any intervals
	if allTrue(monthArr){
		return
	}

	// have to find the first false so we don't begin the partitioning in the middle of an interval
	firstFalse, err := findFalse(0,monthArr)
	if err != nil {
		panic(err.Error())
	}

	// this will be the start of the first interval
	firstTrueAfterFalse, err := findTrue(firstFalse,monthArr)
	if err != nil {
		panic(err.Error())
	}

	intervals = findIntervals(firstTrueAfterFalse, monthArr)
	intervals = setIntervalYears(intervals)

	return
}

// finds a true element given a starting position and an array
func findTrue(start int, arr [MONTHS_IN_YEAR]bool) (int, error) {
	// finds the first false if one is present
	for i := start; i < start + MONTHS_IN_YEAR; i++ {
		if arr[i%MONTHS_IN_YEAR] {
			return (i%MONTHS_IN_YEAR), nil
		}
	}

	return 0, errors.New("True could not be found")

}

// finds a false element given a starting position and an array
func findFalse(start int, arr [MONTHS_IN_YEAR]bool) (int,error) {

	// finds the first false if one is present
	for i := start; i < MONTHS_IN_YEAR; i++ {
		if !arr[i] {
			return i,nil
		}
	}

	return 0, errors.New("False could not be found")
}

// checks if all items in an array are true
func allTrue(arr [MONTHS_IN_YEAR]bool) bool {
	for _, m := range arr {
		if !m {
			return false
		}
	}
	return true
}

// parses the intervals from an array of months the user wants to search in
func findIntervals(startPos int, months [MONTHS_IN_YEAR]bool) (intervals []Interval){

	var inInterval bool = true
	var tempInterval Interval
	tempInterval.StartMonth = startPos + 1

	// generates the month intervals
	for i := startPos; i < startPos + MONTHS_IN_YEAR; i++ {
		if months[i%MONTHS_IN_YEAR] && inInterval {

			tempInterval.EndMonth = (i%MONTHS_IN_YEAR)+1

		} else if months[i%MONTHS_IN_YEAR] {

			tempInterval.StartMonth = (i%MONTHS_IN_YEAR)+1
			tempInterval.EndMonth = (i%MONTHS_IN_YEAR)+1
			inInterval = true

		} else if inInterval {

			intervals = append(intervals, tempInterval)
			inInterval = false

		}
	}
	return
}

// sets the years of the interval start and end depending on the ordering of the months
func setIntervalYears(intervals []Interval) []Interval {

	for i, _ := range intervals {
		intervals[i].StartYear = int(time.Now().Year())
		if intervals[i].StartMonth < intervals[i].EndMonth {
			intervals[i].EndYear = intervals[i].StartYear
		} else {
			intervals[i].EndYear = intervals[i].StartYear + 1
		}
	}

	return intervals
}

// generates a report for a specific user
func reportForUser(user User, db *sql.DB, intervals []Interval) []Flight {

	var minFlight Flight
	var minFlights []Flight
	var potentialMin Flight

	// getting source airports from database
	srcAirports, err := db.Query(fmt.Sprintf(SELECT_SOURCES,user.EmailAddress))
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

			err = db.QueryRow(fmt.Sprintf(MIN_QUERY, potentialMin.sourceAirport, potentialMin.destinationAirport, user.tripMin, user.tripMax, user.budget,buildDateModifiers(intervals))).Scan(&dummy, &dummy, &dummy, &potentialMin.departureDate, &potentialMin.returnDate, &potentialMin.price, &potentialMin.tripLength)
			if err == nil {
				// updating the local cheapest flight
				if minFlight == (Flight{}) {
					minFlight = potentialMin
				} else if potentialMin.price < minFlight.price {
					minFlight = potentialMin
				}
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

// builds up a string to specify the date range of the query
func buildDateModifiers(intervals []Interval) string {

	var datesModifier string = ""

	var first bool = true

	for _, interval := range intervals {
		if first {
			first = false
			datesModifier += andConnective
		} else {
			datesModifier += orConnective
		}
		datesModifier += fmt.Sprintf(monthGreater,interval.StartMonth)
		datesModifier += fmt.Sprintf(monthLesser,interval.EndMonth)
		datesModifier += fmt.Sprintf(yearGreater,interval.StartYear)
		datesModifier += fmt.Sprintf(yearLesser,interval.EndYear)
	}

	return datesModifier
}
