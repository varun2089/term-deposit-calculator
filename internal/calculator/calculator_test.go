package calculator

import (
	"testing"
)

func TestTermDeposit_CalculateFinalBalance(t *testing.T) {
	td := TermDeposit{
		StartDeposit:  10000,
		InterestRate:  1.10,
		InvestmentTerm: 3,
		InterestPaid:  "at-maturity",
	}

	expected := 10330.00
	actual := td.CalculateFinalBalance()

	if actual != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, actual)
	}
}
