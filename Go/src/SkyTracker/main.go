package main

import (
	"Credentials"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"sync"
	"time"
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

			go threadedDataProcess(src, dest)

		}

		// have to reload the result set into destAirports because .Next()
		destAirports, err = db.Query("SELECT * FROM DestinationAirports;")
		if err != nil {
			fmt.Printf("Tried to read the Destination Airport data and failed.")
			panic(err.Error())
		}

	}

	wg.Wait()

}

func threadedDataProcess(src, dest string) {
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	fmt.Printf("%s %s\n", src, dest)
	defer wg.Done()
}
