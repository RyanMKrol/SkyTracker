package Reports

import (
	"database/sql"
	"errors"
)

type Report interface {
	GenerateReport(*sql.DB)
	GetBasicReport() string
	GetFormattedReport() string
}

const (
	STANDARD_REPORT = iota
	BUDGET_REPORT
	DATE_REPORT
	BUDGET_DATE_REPORT
)

// function to create reports
func CreateReportBuilder(t int)(Report, error) {

	switch t {

	case STANDARD_REPORT:
		return new(StandardReport),nil

	case BUDGET_REPORT:
		return new(BudgetReport),nil

	case DATE_REPORT:
		return new(DateReport),nil

	case BUDGET_DATE_REPORT:
		return new(DateBudgetReport),nil

	default:
		//if type is invalid, return an error
		return nil, errors.New("Invalid Appliance Type")
	}
}
