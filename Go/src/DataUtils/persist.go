// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"log"
	"os/exec"
)

const PHP_BINARY string = "/usr/bin/php"
const PERSIST_PHP_LOC string = "/var/www/html/skytracker.co/go/src/DataUtils/persist.php"

// persists the data on the server
func PersistData() {

	cmd := exec.Command(PHP_BINARY, PERSIST_PHP_LOC)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}
