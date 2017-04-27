// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"SystemConfig"
	"fmt"
	"log"
	"os"
)

// sets up the sql file for further input
func SetupSQL(src, dest string) {

	output, err := os.Create(fmt.Sprintf(fmt.Sprintf(SystemConfig.DOC_ROOT, FILE_LOC), src, dest))
	if err != nil {
		fmt.Println("failed to create file setupsql.go")
		log.Fatal(err)
	}
	defer output.Close()
}
