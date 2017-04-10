// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// collects data from a url and returns it, along with the response code
func Collect(url string) (body []byte) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// if the response is 429, we're sending too many requests, so wait a minute and try again
	if responseCode == 429 {
		time.Sleep(60000 * time.Millisecond)
		return Collect(url)
	} else if responseCode != 200 {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("We're about to fail!\n")
		log.Fatal(err)
	}

	return
}
