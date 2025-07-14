package calculator

import (
	"math"
)

func (termDeposit *TermDeposit) CalculateFinalBalance() (float64, error) {
	compoundingPeriodsPerYear, err := termDeposit.getCompoundingPeriodsPerYear()
	if err != nil {
		return 0, err
	}

	if compoundingPeriodsPerYear == 0 {
		finalBalance := termDeposit.StartDeposit * (1 + (termDeposit.InterestRate/100)*termDeposit.InvestmentTerm)
		return math.Round(finalBalance), nil
	}

	periodicRate := (termDeposit.InterestRate / 100) / compoundingPeriodsPerYear
	totalPeriods := termDeposit.InvestmentTerm * compoundingPeriodsPerYear
	finalBalance := termDeposit.StartDeposit * math.Pow(1+periodicRate, totalPeriods)

	finalBalance = math.Round(finalBalance)
	return finalBalance, nil
}
