package Email

import (
	"os/exec"
	"fmt"
	"log"
)

const PHP_BINARY string = "/usr/bin/php"
const EMAIL_PHP_LOC string = "/var/www/html/skytracker.co/Go/src/Email/email.php"
const TITLE_FILE string = "/var/www/html/skytracker.co/Go/src/Email/components/title.txt"
const BODY_FILE string = "/var/www/html/skytracker.co/Go/src/Email/components/body.txt"

// persists the data on the server
func Email(attachmentLocation string) {

	cmd := exec.Command(PHP_BINARY, EMAIL_PHP_LOC, attachmentLocation, TITLE_FILE, BODY_FILE)

	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to execute command in email.go")
		log.Fatal(err)
	}

	return
}
