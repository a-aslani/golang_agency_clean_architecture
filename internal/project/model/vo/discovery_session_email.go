package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"regexp"
	"strings"
)

type DiscoverySessionEmail string

const emailPatternRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`

var emailRe = regexp.MustCompile(emailPatternRegex)

const emailMaxLen = 60

func (r DiscoverySessionEmail) Validate() error {

	email := strings.TrimSpace(r.String())

	if email == "" {
		return errorenum.EmailIsRequired
	}

	if len(email) > emailMaxLen {
		return errorenum.MaxLenErr.Var("email", emailMaxLen, len(email))
	}

	if !emailRe.MatchString(email) {
		return errorenum.InvalidEmailAddress.Var(email)
	}

	return nil
}

func (r DiscoverySessionEmail) String() string {
	return string(r)
}
