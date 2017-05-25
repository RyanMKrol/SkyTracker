// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// http response codes
const EXCESSIVE_CALLS int = 429
const GOOD_RESPONSE int = 200
const SLEEP_TIME time.Duration = 60000

// collects data from a url and returns it, along with the response code
func Collect(url string) (body []byte) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to get http collect.go")
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// reading data from the response
	var responseCode int = resp.StatusCode

	// if the response is 429, we're sending too many requests, so wait a minute and try again
	if responseCode == EXCESSIVE_CALLS {
		time.Sleep(SLEEP_TIME * time.Millisecond)
		return Collect(url)
	} else if responseCode != GOOD_RESPONSE {
		fmt.Println("just not a good response")
		fmt.Println(responseCode)
		fmt.Println(url)
		fmt.Println(resp.Status)
		log.Fatal(err)
	}

	// parses all data into a byte slice
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read all collect.go")
		log.Fatal(err)
	}

	return
}
