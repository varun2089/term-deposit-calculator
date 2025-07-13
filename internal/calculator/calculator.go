package calculator

import "math"

type TermDeposit struct {
	StartDeposit  float64
	InterestRate  float64
	InvestmentTerm int
	InterestPaid  string
}

func (td *TermDeposit) CalculateFinalBalance() float64 {
	n := td.getCompoundingPeriodsPerYear()
	// Convert interest rate from percentage to decimal
	r := td.InterestRate / 100
	// Convert investment term to float
	t := float64(td.InvestmentTerm)

	// Handle "at maturity" case
	if n == 0 {
		return td.StartDeposit * (1 + r*t)
	}

	// A = P(1 + r/n)^(nt)
	return td.StartDeposit * math.Pow(1+r/n, n*t)
}

func (td *TermDeposit) getCompoundingPeriodsPerYear() float64 {
	switch td.InterestPaid {
	case "monthly":
		return 12
	case "quarterly":
		return 4
	case "annually":
		return 1
	case "at-maturity":
		return 0
	default:
		return 0
	}
}