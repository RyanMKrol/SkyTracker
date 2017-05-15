package Reports

// struct to parse the data into
type User struct {
	EmailAddress    string
	budget          int
	tripMin         int
	tripMax         int
	salt            string
	months          []int
	sources         []string
	ReportFrequency int
	ReportLoc       string
	HasReport 			bool
}
