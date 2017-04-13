package main

import (
	"DataUtils"
	"Reports"
	"Credentials"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

const MONTHS_LOOKAHEAD = 6
const MONTHS_TRIP_MAX = 2
const URL_FORMAT = "http://partners.api.skyscanner.net/apiservices/browsegrid/v1.0/GB/GBP/en-GB/%s/%s/%s/%s?apiKey=%s"
const COMPRESS_SOURCE = "./../../sql/raw"
const COMPRESS_TARGET = "./../../sql/compressed/archive.zip"
const SELECT_SOURCES = "SELECT * FROM SourceAirports;"
const SELECT_DESTINATIONS = "SELECT * FROM DestinationAirports;"

// this is used to sync up the threads that are doing work before we continue
var wg sync.WaitGroup

func main() {

	user, password, ip, database := Credentials.User(), Credentials.Password(), Credentials.IPAddress(), Credentials.DatabaseName()
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, ip, database)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// statements to get the source and destination airport pairs

	srcAirports, err := db.Query(SELECT_SOURCES)
	if err != nil {
		panic(err.Error())
	}
	defer srcAirports.Close()

	destAirports, err := db.Query(SELECT_DESTINATIONS)
	if err != nil {
		panic(err.Error())
	}
	defer destAirports.Close()

	// sending off each thread

	for srcAirports.Next() {

		for destAirports.Next() {

			// adding another thread to wait for
			wg.Add(1)

			var src, dest, dummy string

			if err := srcAirports.Scan(&dummy, &dummy, &src, &dummy, &dummy); err != nil {
				panic(err.Error())
			}

			if err := destAirports.Scan(&dummy, &dummy, &dest, &dummy, &dummy); err != nil {
				panic(err.Error())
			}

			fmt.Println(src + " " + dest)

			go t_DataProcess(src, dest)

		}

		// have to reload the result set into destAirports because .Next()
		destAirports, err = db.Query(SELECT_DESTINATIONS)
		if err != nil {
			panic(err.Error())
		}

		wg.Wait()
	}

	at this point all of the files will be setup, now I need to persist it with the server

	DataUtils.PersistData()

	Reports.GenerateReport(db)

}

// this function will be for gathering and persisting data with threads
func t_DataProcess(src, dest string) {

	departTime := time.Now()
	returnTime := time.Now()

	// creates the relevant sql file
	DataUtils.SetupSQL(src,dest)

	// goes through each date and collects data
	for i := 1; i <= MONTHS_LOOKAHEAD; i++ {

		for j := 0; j < MONTHS_TRIP_MAX; j++ {

			departDate := departTime.Format("2006-01")
			returnDate := returnTime.Format("2006-01")

			url := fmt.Sprintf(URL_FORMAT,src,dest,departDate,returnDate,Credentials.ApiKey())

			response := DataUtils.Collect(url)

			// decodes the response and stores it in the .sql file
			DataUtils.Decode(response, src, dest, departTime, returnTime)

			returnTime = returnTime.AddDate(0,1,0)
		}

		departTime = departTime.AddDate(0,1,0)

		returnTime = time.Now()
		returnTime = returnTime.AddDate(0,i,0)

	}

	wg.Done()
}
