package Reports

import (
	"sort"
)

type User struct {
	emailAddress string
	budget int
	tripMin int
	tripMax int
	reportLoc string
	prettyReportLoc string
}
