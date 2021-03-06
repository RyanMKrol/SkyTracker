package Email

import (
	"Reports"
	"SystemConfig"
	"fmt"
	"log"
	"os/exec"
)

// file location constants
const PHP_BINARY string = "/usr/bin/php"
const EMAIL_PHP_LOC string = "src/Email/email.php"
const TITLE_FILE string = "src/Email/components/title.txt"

// persists the data on the server
func EmailUsers(users []Reports.User) {

	for _, user := range users {

		if user.HasReport {
			cmd := exec.Command(PHP_BINARY, fmt.Sprintf(SystemConfig.DOC_ROOT, EMAIL_PHP_LOC), fmt.Sprintf(SystemConfig.DOC_ROOT, TITLE_FILE), user.ReportLoc, user.EmailAddress)

			err := cmd.Run()
			if err != nil {
				fmt.Println("failed to execute command in email.go")
				log.Fatal(err)
			}
		}
	}
}
