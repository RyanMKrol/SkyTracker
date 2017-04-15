package Reports

import (
	"database/sql"
)

type DateBudgetReport struct {
	basicReportLoc       string
	beautifiedReportloc  string
}

func (r *DateBudgetReport) GetBasicReport() string {
	return r.basicReportLoc
}
func (r *DateBudgetReport) GetFormattedReport() string {
	return r.beautifiedReportloc
}

func (r *DateBudgetReport) GenerateReport(db *sql.DB) {


}
