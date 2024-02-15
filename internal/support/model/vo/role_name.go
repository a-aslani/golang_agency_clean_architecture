package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"strings"
)

type RoleName string

func (r RoleName) Validate() error {

	if strings.TrimSpace(r.String()) == "" {
		return errorenum.RoleNameIsRequired
	}

	return nil
}

func (r RoleName) String() string {
	return string(r)
}
