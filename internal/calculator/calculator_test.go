package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateFinalBalance(t *testing.T) {
	tests := []struct {
		name           string
		startDeposit   float64
		interestRate   float64
		investmentTerm float64
		interestPaid   string
		expected       float64
		expectError    bool
	}{
		{
			name:           "At Maturity",
			startDeposit:   1000,
			interestRate:   5,
			investmentTerm: 2,
			interestPaid:   "at-maturity",
			expected:       1100,
		},
		{
			name:           "Monthly Compounding",
			startDeposit:   1000,
			interestRate:   12,
			investmentTerm: 1,
			interestPaid:   "monthly",
			expected:       1127,
		},
		{
			name:           "Quarterly Compounding",
			startDeposit:   2000,
			interestRate:   8,
			investmentTerm: 3,
			interestPaid:   "quarterly",
			expected:       2536,
		},
		{
			name:           "Annually Compounding",
			startDeposit:   5000,
			interestRate:   6,
			investmentTerm: 5,
			interestPaid:   "annually",
			expected:       6691,
		},
		{
			name:           "Invalid Interest Frequency",
			startDeposit:   1000,
			interestRate:   5,
			investmentTerm: 1,
			interestPaid:   "weekly",
			expectError:    true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			termDeposit := &TermDeposit{
				StartDeposit:   testCase.startDeposit,
				InterestRate:   testCase.interestRate,
				InvestmentTerm: testCase.investmentTerm,
				InterestPaid:   testCase.interestPaid,
			}

			result, err := termDeposit.CalculateFinalBalance()

			if testCase.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, result)
			}
		})
	}
}
