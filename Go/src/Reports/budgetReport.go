package Reports

import (
	"database/sql"
)

type BudgetReport struct {
	basicReportLoc       string
	beautifiedReportloc  string
	budget 					     int
}

// const selectSources string = "SELECT * FROM SourceAirports;"
// const selectDestinations string = "SELECT * FROM DestinationAirports;"
// const minQuery string = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM %s_%s Where Price = (SELECT Min(Price) FROM %s_%s WHERE DATEDIFF(ReturnDate, DepartDate) > 2) AND DATEDIFF(ReturnDate, DepartDate) > 2 limit 1;"
// const reportLoc string = "reports/BudgetReport:%s.csv"
// const csvLineFormat string = "\"%s, %s\",\"%s, %s, %s\",%s,%s,%d,%d\n"
// const csvHeaders string = "From,To,Leaving,Returning,Trip Length,Cost\n"
// const dateFormat string = "2006-01-02"

func (r *BudgetReport) GetBasicReport() string {
	return r.basicReportLoc
}
func (r *BudgetReport) GetFormattedReport() string {
	return r.beautifiedReportloc
}

func (r *BudgetReport) GenerateReport(db *sql.DB) {
}
