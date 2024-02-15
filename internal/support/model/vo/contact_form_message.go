package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"strings"
)

type ContactFormMessage string

func (r ContactFormMessage) Validate() error {
	if strings.TrimSpace(r.String()) == "" {
		return errorenum.ContactFormMessageIsRequired
	}

	return nil
}

func (r ContactFormMessage) String() string {
	return string(r)
}
