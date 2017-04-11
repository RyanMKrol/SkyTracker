// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"os/exec"
	"Credentials"
	"fmt"
	"log"
)

const MERGE_SQL_FILES = "cat ./../../sql/raw/* > ./../../sql/all.sql"
const DATABASE_USING_SOURCE = "mysql -u %s -p\"%s\" -f \"%s\" < ./../../sql/all.sql"

// persists the data on the server
func PersistData(source string) {

	cmd := exec.Command(MERGE_SQL_FILES)
	persistCmd := exec.Command(fmt.Sprintf(DATABASE_USING_SOURCE, Credentials.User(), Credentials.Password(), Credentials.DatabaseName))

	err := cmd.Run()
	if err != nil{
		log.Fatal(err)
	}

	err = persistCmd.Run()
	if err != nil{
		log.Fatal(err)
	}

	return
}
