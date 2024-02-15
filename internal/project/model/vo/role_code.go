package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"strings"
)

type RoleCode string

func (r RoleCode) Validate() error {

	if strings.TrimSpace(r.String()) == "" {
		return errorenum.RoleCodeIsRequired
	}

	return nil
}

func (r RoleCode) String() string {
	return string(r)
}
