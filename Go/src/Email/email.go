package Email

import (
	"log"
	"os/exec"
	"fmt"
	"io/ioutil"
)

const PHP_BINARY string = "/usr/bin/php"
const EMAIL_PHP_LOC string = "/var/www/html/skytracker.co/Go/src/Email/email.php"
const TITLE_FILE string = "./../components/title.txt"
const BODY_FILE string = "./../components/body.txt"

// persists the data on the server
func Email(attachmentLocation string) {

	title,err := ioutil.ReadFile(TITLE_FILE)
	if err != nil {
		fmt.Println("failed to open file for title text email.go")
		log.Fatal(err)
	}

	titleString := string(title)

	body,err := ioutil.ReadFile(BODY_FILE)
	if err != nil {
		fmt.Println("failed to open file for body text email.go")
		log.Fatal(err)
	}

	bodyString := string(body)

	cmd := exec.Command(PHP_BINARY, EMAIL_PHP_LOC, attachmentLocation, titleString, bodyString)

	err = cmd.Run()
	if err != nil {
		fmt.Println("failed to execute command in email.go")
		log.Fatal(err)
	}

	return
}
