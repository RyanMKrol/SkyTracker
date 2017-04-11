package Email

import (
	"Credentials"
	"fmt"
	"log"
	"os/exec"
	"time"
)

const DATABASE_USING_SOURCE = "mysql -u %s -p\"%s\" -f \"%s\" < ./../../sql/all.sql"
const SEND_EMAIL_CMD = "echo \"Please find attached the report of cheap flights for Europe!\" | mail -A %s -s \"Your Daily Cheap Flights Report!\" %s"
const SELECT_EMAIL_ADDRESSES = "SELECT UserEmailAddress FROM Users;"
const CONNECTION_FORMAT = "%s:%s@tcp(%s)/%s"
const DATABASE_DRIVER = "mysql"

func Email(attachmentLocation string) {

	user, password, ip, database := Credentials.User(), Credentials.Password(), Credentials.IPAddress(), Credentials.DatabaseName()
	conn := fmt.Sprintf(CONNECTION_FORMAT, user, password, ip, database)

	db, err := sql.Open(DATABASE_DRIVER, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	emailAddresses, err := db.Query(SELECT_EMAIL_ADDRESSES)
	if err != nil {
		log.Fatal(err)
	}
	defer srcAirports.Close()


	// going through and sending report to users
	for emailAddresses.Next() {

		var address string
		if err := srcAirports.Scan(&address); err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command(fmt.Sprintf(SEND_EMAIL_CMD, attachmentLocation, address))

		err = cmd.Run()
		if err != nil{
			log.Fatal(err)
		}
	}

}
