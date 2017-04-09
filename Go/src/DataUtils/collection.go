//package contains utility functions for manipulating the data we want
package DataUtils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// collects data from a url and returns it, along with the response code
func collect(url string) (body []byte, responseCode int) {
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

  // reading data from the response
	responseCode = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return
}
