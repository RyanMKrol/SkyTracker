//package contains utility functions for manipulating the data we want
package DataUtils

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

// collects data from a url and returns it, along with the response code
func Collect(url string) (body []byte, responseCode int) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("We're about to fail.\n")
		log.Fatal(err)
	}
	defer resp.Body.Close()

  // reading data from the response
	responseCode = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("We're about to fail further on!\n")
		log.Fatal(err)
	}

	return
}
