package Reports

import (
	"database/sql"
)

type DateReport struct {
	basicReportLoc       string
	beautifiedReportloc  string
}

func (r *DateReport) GetBasicReport() string {
	return r.basicReportLoc
}
func (r *DateReport) GetFormattedReport() string {
	return r.beautifiedReportloc
}

func (r *DateReport) GenerateReport(db *sql.DB) {


}
