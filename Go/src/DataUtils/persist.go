// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"SystemConfig"
	"log"
	"os/exec"
	"fmt"
)

const PHP_BINARY string = "/usr/bin/php"
const PERSIST_PHP_LOC string = "src/DataUtils/persist.php"

// persists the data on the server
func PersistData() {

	cmd := exec.Command(PHP_BINARY, fmt.Sprintf(SystemConfig.DOC_ROOT,PERSIST_PHP_LOC))

	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to run command persist.go")
		log.Fatal(err)
	}

	return
}
