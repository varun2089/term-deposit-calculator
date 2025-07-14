package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCompoundingPeriodsPerYear(t *testing.T) {
	tests := []struct {
		interestPaid string
		expected     float64
		expectError  bool
	}{
		{"monthly", 12, false},
		{"quarterly", 4, false},
		{"annually", 1, false},
		{"at-maturity", 0, false},
		{"WEEKLY", 0, true},
	}

	for _, testCase := range tests {
		td := &TermDeposit{InterestPaid: testCase.interestPaid}
		result, err := td.getCompoundingPeriodsPerYear()

		if testCase.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, result)
		}
	}
}
