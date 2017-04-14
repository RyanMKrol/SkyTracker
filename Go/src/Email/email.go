package Email

import (
	"SystemConfig"
	"os/exec"
	"fmt"
	"log"
)

const PHP_BINARY string = "/usr/bin/php"
const EMAIL_PHP_LOC string = "src/Email/email.php"
const TITLE_FILE string = "src/Email/components/title.txt"
const BODY_FILE string = "src/Email/components/body.txt"

// persists the data on the server
func Email(attachmentLocation string) {

	cmd := exec.Command(PHP_BINARY, fmt.Sprintf(SystemConfig.DOC_ROOT,EMAIL_PHP_LOC), attachmentLocation, fmt.Sprintf(SystemConfig.DOC_ROOT,TITLE_FILE), fmt.Sprintf(SystemConfig.DOC_ROOT,BODY_FILE))

	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to execute command in email.go")
		log.Fatal(err)
	}

	return
}
