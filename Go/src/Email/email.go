package Email

import (
	"SystemConfig"
	"Reports"
	"fmt"
	"os/exec"
)

const PHP_BINARY string = "/usr/bin/php"
const EMAIL_PHP_LOC string = "src/Email/email.php"
const TITLE_FILE string = "src/Email/components/title.txt"
const BODY_FILE string = "src/Email/components/body.txt"

// persists the data on the server
func EmailUsers(users []Reports.User) {

	for _, user := range users {

		cmd := exec.Command(PHP_BINARY, fmt.Sprintf(SystemConfig.DOC_ROOT,EMAIL_PHP_LOC), user.ReportLoc, fmt.Sprintf(SystemConfig.DOC_ROOT,TITLE_FILE), fmt.Sprintf(SystemConfig.DOC_ROOT,BODY_FILE), user.EmailAddress)

		_ = cmd
		// err := cmd.Run()
		// if err != nil {
		// 	fmt.Println("failed to execute command in email.go")
		// 	log.Fatal(err)
		// }

	}
}
