package Email

import (
	"log"
	"os/exec"
)

const PHP_BINARY string = "/usr/bin/php"
const EMAIL_PHP_LOC string = "/var/www/html/skytracker.co/go/src/Email/email.php"

// persists the data on the server
func Email(attachmentLocation string) {

	cmd := exec.Command(PHP_BINARY, EMAIL_PHP_LOC, attachmentLocation)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}
