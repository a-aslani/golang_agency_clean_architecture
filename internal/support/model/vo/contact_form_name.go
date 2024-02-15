package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"strings"
)

type ContactFormName string

func (r ContactFormName) Validate() error {
	if strings.TrimSpace(r.String()) == "" {
		return errorenum.ContactFormNameIsRequired
	}

	return nil
}

func (r ContactFormName) String() string {
	return string(r)
}
