package main

import (
	"DataUtils"
	"Credentials"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

// this is used to sync up the threads that are doing work before we continue
var wg sync.WaitGroup

const MonthsLookahead = 6
const MonthsTripDurationMax = 2

func main() {

	user, password, ip, database := Credentials.User(), Credentials.Password(), Credentials.IPAddress(), Credentials.DatabaseName()
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, ip, database)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// statements to get the source and destination airport pairs

	srcAirports, err := db.Query("SELECT * FROM SourceAirports;")
	if err != nil {
		fmt.Printf("Tried to read the Source Airport data and failed.")
		panic(err.Error())
	}
	defer srcAirports.Close()

	destAirports, err := db.Query("SELECT * FROM DestinationAirports;")
	if err != nil {
		fmt.Printf("Tried to read the Destination Airport data and failed.")
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
				fmt.Printf("Tried to read into source airports and shit the bed.")
				panic(err.Error())
			}

			if err := destAirports.Scan(&dummy, &dummy, &dest, &dummy, &dummy); err != nil {
				fmt.Printf("Tried to read into destination airports and shit the bed.")
				panic(err.Error())
			}

			go t_DataProcess(src, dest)

		}

		// have to reload the result set into destAirports because .Next()
		destAirports, err = db.Query("SELECT * FROM DestinationAirports;")
		if err != nil {
			fmt.Printf("Tried to read the Destination Airport data and failed.")
			panic(err.Error())
		}

		wg.Wait()
	}

}

// this function will be for gathering and persisting data with threads
func t_DataProcess(src, dest string) {

	// the times to depart and return

	departTime := time.Now()
	returnTime := time.Now()

	var departDate string
	var returnDate string

	for i := 1; i <= MonthsLookahead; i++ {

		for j := 0; j < MonthsTripDurationMax; j++ {

			departDate = departTime.Format("2006-01")
			returnDate = returnTime.Format("2006-01")

			url := fmt.Sprintf("http://partners.api.skyscanner.net/apiservices/browsegrid/v1.0/GB/GBP/en-GB/%s/%s/%s/%s?apiKey=%s",src,dest,departDate,returnDate,Credentials.ApiKey())

			response []byte := DataUtils.Collect(url)

			DataUtils.Decode(response)

			returnTime = returnTime.AddDate(0,1,0)
		}

		departTime = departTime.AddDate(0,1,0)

		returnTime = time.Now()
		returnTime = returnTime.AddDate(0,i,0)

	}

	wg.Done()
}
