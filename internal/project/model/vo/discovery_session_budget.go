package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"strings"
)

type DiscoverySessionBudget string

func (d DiscoverySessionBudget) Validate() error {

	if strings.TrimSpace(d.String()) == "" {
		return errorenum.BudgetIsRequired
	}

	return nil
}

func (d DiscoverySessionBudget) String() string {
	return string(d)
}
