package main

import (
	"Credentials"
	"DataUtils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

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
			//adding another thread to wait for
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

		wg.Wait()

		// have to reload the result set into destAirports because .Next()
		destAirports, err = db.Query("SELECT * FROM DestinationAirports;")
		if err != nil {
			fmt.Printf("Tried to read the Destination Airport data and failed.")
			panic(err.Error())
		}

	}

}

// this function will be for gathering and persisting data with threads
func t_DataProcess(src, dest string) {

	// i need to build up the url here and then
	DataUtils.Collect("http://partners.api.skyscanner.net/apiservices/browsegrid/v1.0/GB/GBP/en-GB/LHR/ORY/2017-08/2017-09?apiKey=na912611636759734898754178423831",count)

	defer wg.Done()
}
