package Reports

import (
	"SystemConfig"
	"database/sql"
	"errors"
	"strings"
	"strconv"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// for databaser interaction
const SELECT_TRAVEL_MONTHS string = "SELECT * FROM UserTravelMonths WHERE UserEmailAddress = \"%s\" ORDER BY TravelMonth ASC;"
const SELECT_USERS string = "SELECT * FROM Users WHERE UserLastReport IS NULL OR DATEDIFF(NOW(),UserLastReport) = (Select FLDays from FrequencyLookup WHERE FLID = UserReportFrequency);"
const SELECT_SOURCES string = "SELECT * FROM SourceAirports WHERE SrcAirportCode IN (SELECT SourceAirportCode FROM UserSourceAirports WHERE UserEmailAddress = \"%s\");"
const SELECT_DESTINATIONS string = "SELECT * FROM DestinationAirports;"
const MIN_QUERY string = "(select *, DATEDIFF(ReturnDate, DepartDate) from Flights where DestinationAirportCode = '%s' AND DATEDIFF(ReturnDate, DepartDate) >= %d AND DATEDIFF(ReturnDate, DepartDate) <= %d AND Price <= %d AND SourceAirportCode IN (SELECT SourceAirportCode FROM UserSourceAirports WHERE UserEmailAddress = '%s') %s ORDER BY Price ASC limit 1)"
const UPDATE_REPORT string = "UPDATE Users SET UserLastReport = NOW() WHERE UserLastReport IS NULL OR DATEDIFF(NOW(),UserLastReport) = (Select FLDays from FrequencyLookup WHERE FLID = UserReportFrequency);"

// for the report writing
const REPORT_LOC string = "reports/%d_%d_%d_%s_%s_%s.html"
const DATE_FORMAT string = "2006-01-02"

// for the query builder
const monthGreater string = "(Month(DepartDate) >= %d AND "
const monthLesser string = "Month(ReturnDate) <= %d  AND "
const yearGreater string = "Year(DepartDate) >= %d  AND "
const yearLesser string = "Year(ReturnDate) <= %d)"
const orConnective string = " OR "
const andConnective string = " AND "
const openingStatement string = "SELECT * FROM SourceAirports INNER JOIN ("
const endingStatement string = ") AS t1 ON SourceAirports.SrcAirportCode = t1.SourceAirportCode INNER JOIN DestinationAirports ON t1.DestinationAirportCode = DestinationAirports.DestAirportCode;"

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
		users[i].HasReport = false

		var filename string = fmt.Sprintf(fmt.Sprintf(SystemConfig.DOC_ROOT, REPORT_LOC), users[i].budget, users[i].tripMin, users[i].tripMax, getAirportsString(users[i].sources), getMonthsString(users[i].months), currentDate.Format(DATE_FORMAT))

		// file doesn't exist so we need to make it ourselves
		if _, err := os.Stat(filename); os.IsNotExist(err) {

			file, err := os.Create(filename)
			if err != nil {
				fmt.Println("failed to create report item generate.go")
				log.Fatal(err)
			}

			wg.Add(1)

			// parallelising the meat of the file
			go func(u *User, f *os.File) {

				fmt.Println((*u).EmailAddress)
				intervals := intervalBuilder(*u, db)
				fmt.Println(intervals)
				minFlights := reportForUser(*u, db, intervals)
				generatePrettyReport(minFlights, f, u)
				f.Close()
				wg.Done()
			}(&users[i], file)
		}

		users[i].ReportLoc = filename
	}

	wg.Wait()

	// updates the users LastReportDate to today
	_, err := db.Exec(UPDATE_REPORT)
	if err != nil {
		fmt.Println("failed to update user last report dates generate.go")
		panic(err.Error())
	}

	return users
}

//convert an arry of airport strings to a single string
func getAirportsString(arr []string) string {
	return strings.Join(arr, "")
}

//convert an array of months into a single string
func getMonthsString(arr []int) string {
	var intString string = ""

	for _, i := range arr {
		intString += strconv.Itoa(i)
	}

	return intString
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

	//parsing the user informaton
	for users.Next() {

		var tempUser User = User{}

		// getting the other information from the user
		var dummy sql.NullString
		var maybeBudget, maybeTripMin, maybeTripMax sql.NullInt64

		if err := users.Scan(&dummy, &tempUser.EmailAddress, &maybeBudget, &maybeTripMin, &maybeTripMax, &tempUser.ReportFrequency, &dummy, &tempUser.salt); err != nil {
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

		// getting the months that the user wants to travel in
		var monthArr []int
		months, err := db.Query(fmt.Sprintf(SELECT_TRAVEL_MONTHS, tempUser.EmailAddress))
		if err != nil {
			fmt.Println("failed to get users generate.go")
			panic(err.Error())
		}

		//parsing the months into the user
		for months.Next() {

			var dummy string
			var month int

			if err := months.Scan(&dummy, &dummy, &month); err != nil {
				fmt.Println("failed to scan month generate.go")
				panic(err.Error())
			}

			monthArr = append(monthArr, month)
		}

		tempUser.months = monthArr

		// getting the airports that the user wants to fly from
		var airportArr []string
		airports, err := db.Query(fmt.Sprintf(SELECT_SOURCES, tempUser.EmailAddress))
		if err != nil {
			fmt.Println("failed to get users generate.go")
			panic(err.Error())
		}

		//parsing the airport information
		for airports.Next() {

			var dummy string
			var airport string

			if err := airports.Scan(&dummy, &dummy, &airport, &dummy, &dummy); err != nil {
				fmt.Println("failed to scan airport generate.go")
				panic(err.Error())
			}

			airportArr = append(airportArr, airport)
		}
		tempUser.sources = airportArr

		months.Close()
		airports.Close()
		userArr = append(userArr, tempUser)
	}

	return userArr
}

// builds up the intervals to search over that a user will have specified
func intervalBuilder(user User, db *sql.DB) (intervals []Interval) {

	var monthArr [MONTHS_IN_YEAR]bool

	for _, month := range user.months {
		monthArr[month-1] = true
	}

	// if they're all true we don't need any intervals
	if allTrue(monthArr) {
		return
	}

	// have to find the first false so we don't begin the partitioning in the middle of an interval
	firstFalse, err := findFalse(0, monthArr)
	if err != nil {
		panic(err.Error())
	}

	// this will be the start of the first interval
	firstTrueAfterFalse, err := findTrue(firstFalse, monthArr)
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
	for i := start; i < start+MONTHS_IN_YEAR; i++ {
		if arr[i%MONTHS_IN_YEAR] {
			return (i % MONTHS_IN_YEAR), nil
		}
	}

	return 0, errors.New("True could not be found")

}

// finds a false element given a starting position and an array
func findFalse(start int, arr [MONTHS_IN_YEAR]bool) (int, error) {

	// finds the first false if one is present
	for i := start; i < MONTHS_IN_YEAR; i++ {
		if !arr[i] {
			return i, nil
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
func findIntervals(startPos int, months [MONTHS_IN_YEAR]bool) (intervals []Interval) {

	var inInterval bool = true
	var tempInterval Interval
	tempInterval.StartMonth = startPos + 1

	// generates the month intervals
	for i := startPos; i < startPos+MONTHS_IN_YEAR; i++ {
		if months[i%MONTHS_IN_YEAR] && inInterval {

			tempInterval.EndMonth = (i % MONTHS_IN_YEAR) + 1

		} else if months[i%MONTHS_IN_YEAR] {

			tempInterval.StartMonth = (i % MONTHS_IN_YEAR) + 1
			tempInterval.EndMonth = (i % MONTHS_IN_YEAR) + 1
			inInterval = true

		} else if inInterval {

			intervals = append(intervals, tempInterval)
			inInterval = false

		}
	}

	//the following mechanism splits up an interval if it spands over the current month in the years
	// this means april - june when ran in may will become april - may and may - june
	//  later on the first interval will look in the next year, the latter looking in this years
	currentmonth := int(time.Now().Month())
	for i, interval := range intervals {
		if (interval.StartMonth < currentmonth && interval.EndMonth >= currentmonth) || (interval.StartMonth < currentmonth && interval.EndMonth < interval.StartMonth) {
			intervalsTail := intervals[i+1:]
			intervals = intervals[:i]
			intervals = append(intervals, intervalsTail...)
			var intervalTemp Interval = Interval{StartMonth: interval.StartMonth, EndMonth: currentmonth}
			intervals = append(intervals, intervalTemp)
			intervalTemp = Interval{StartMonth: currentmonth, EndMonth: interval.EndMonth}
			intervals = append(intervals, intervalTemp)
			break
		}
	}

	fmt.Println(intervals)
	return
}

// sets the years of the interval start and end depending on the ordering of the months
func setIntervalYears(intervals []Interval) []Interval {

	currentmonth := int(time.Now().Month())
	currentyear := int(time.Now().Year())

	for i, _ := range intervals {

		if intervals[i].StartMonth < currentmonth {

			intervals[i].StartYear = currentyear + 1
			intervals[i].EndYear = currentyear + 1

		} else {
			intervals[i].StartYear = currentyear

			if intervals[i].StartMonth <= intervals[i].EndMonth {
				intervals[i].EndYear = currentyear
			} else {
				intervals[i].EndYear = currentyear + 1
			}
		}
	}

	return intervals
}

// generates a report for a specific user
func reportForUser(user User, db *sql.DB, intervals []Interval) []Flight {

	var minFlights []Flight

	// getting destination airports from database
	destAirports, err := db.Query(SELECT_DESTINATIONS)
	if err != nil {
		fmt.Println("failed to get destinations generate.go")
		panic(err.Error())
	}
	defer destAirports.Close()

	var destinationInfo []Flight
	fmt.Println("getting all of the flights in one area")

	for destAirports.Next() {
		var dummy string
		var tempFlight = Flight{}

		if err := destAirports.Scan(&dummy, &dummy, &tempFlight.destinationAirport, &tempFlight.destinationCountry, &tempFlight.destinationCity); err != nil {
			fmt.Println("failed to scan destinations airports generate.go")
			panic(err.Error())
		}

		destinationInfo = append(destinationInfo, tempFlight)
	}

	var query string = queryBuilder(destinationInfo, intervals, user)

	flights, err := db.Query(query)
	if err != nil {
		fmt.Println("failed to get destinations generate.go")
		panic(err.Error())
	}
	defer flights.Close()

	for flights.Next() {
		var dummy string
		var tempFlight = Flight{}

		if err := flights.Scan(&dummy, &dummy, &dummy, &tempFlight.sourceCountry, &tempFlight.sourceCity, &dummy, &tempFlight.sourceAirport, &dummy, &tempFlight.departureDate, &tempFlight.returnDate, &tempFlight.price, &tempFlight.tripLength, &dummy, &dummy, &tempFlight.destinationAirport, &tempFlight.destinationCountry, &tempFlight.destinationCity); err != nil {
			fmt.Println("failed to scan destinations airports generate.go")
			panic(err.Error())
		}

		minFlights = append(minFlights, tempFlight)
	}

	return minFlights
}

func queryBuilder(destinations []Flight, intervals []Interval, user User) string {

	var first bool = true
	var query string = openingStatement

	for _, destInfo := range destinations {

		if first {
			first = false
		} else {
			query += " UNION \n"
		}

		query += fmt.Sprintf(MIN_QUERY, destInfo.destinationAirport, user.tripMin, user.tripMax, user.budget, user.EmailAddress, buildDateModifiers(intervals))

	}
	query += endingStatement

	return query

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
		datesModifier += fmt.Sprintf(monthGreater, interval.StartMonth)
		datesModifier += fmt.Sprintf(monthLesser, interval.EndMonth)
		datesModifier += fmt.Sprintf(yearGreater, interval.StartYear)
		datesModifier += fmt.Sprintf(yearLesser, interval.EndYear)
	}

	return datesModifier
}
