package Reports

import (
	"database/sql"
)

type BudgetReport struct {
	basicReportLoc       string
	beautifiedReportloc  string
}

func (r *BudgetReport) GetBasicReport() string {
	return r.basicReportLoc
}
func (r *BudgetReport) GetFormattedReport() string {
	return r.beautifiedReportloc
}

func (r *BudgetReport) GenerateReport(db *sql.DB) {


}
