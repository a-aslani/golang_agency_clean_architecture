package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"regexp"
	"strings"
)

const emailPatternRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`

var emailRe = regexp.MustCompile(emailPatternRegex)

const emailMaxLen = 60

type ContactFormEmail string

func (r ContactFormEmail) Validate() error {

	email := strings.TrimSpace(r.String())

	if email == "" {
		return errorenum.ContactFormEmailIsRequired
	}

	if len(email) > emailMaxLen {
		return errorenum.MaxLenErr.Var("email", emailMaxLen, len(email))
	}

	if !emailRe.MatchString(email) {
		return errorenum.InvalidEmailAddress.Var(email)
	}

	return nil
}

func (r ContactFormEmail) String() string {
	return string(r)
}
