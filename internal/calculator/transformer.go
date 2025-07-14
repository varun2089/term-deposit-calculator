package calculator

import (
	"fmt"
	"strings"
)

func (termDeposit *TermDeposit) getCompoundingPeriodsPerYear() (float64, error) {
	switch strings.ToLower(termDeposit.InterestPaid) {
	case "monthly":
		return 12, nil
	case "quarterly":
		return 4, nil
	case "annually":
		return 1, nil
	case "at-maturity":
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported interest payment frequency: %s", termDeposit.InterestPaid)
	}
}
