// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"fmt"
	"log"
	"os"
)

// sets up the sql file for further input
func SetupSQL(src, dest string) {

	output, err := os.Create(fmt.Sprintf(FILE_LOC, src, dest))
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	output.WriteString(fmt.Sprintf("DROP TABLE %s_%s;\n\n", src, dest))
	output.WriteString(fmt.Sprintf("CREATE TABLE %s_%s(\n\n", src, dest))
	output.WriteString("TripID int NOT NULL AUTO_INCREMENT,\n")
	output.WriteString("SourcePort varchar(255) NOT NULL,\n")
	output.WriteString("DestPort varchar(255) NOT NULL,\n")
	output.WriteString("DepartDate DATE NOT NULL,\n")
	output.WriteString("ReturnDate DATE NOT NULL,\n")
	output.WriteString("Price int NOT NULL,\n")
	output.WriteString("PRIMARY KEY(TripID),\n")
	output.WriteString("CONSTRAINT uc_date_pair UNIQUE (DepartDate, ReturnDate)\n")
	output.WriteString(");\n\n")
}
