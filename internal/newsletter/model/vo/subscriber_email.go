package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/errorenum"
	"regexp"
	"strings"
)

type SubscriberEmail string

const emailPatternRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`

var emailRe = regexp.MustCompile(emailPatternRegex)

func (r SubscriberEmail) Validate() error {

	email := strings.TrimSpace(r.String())

	if email == "" {
		return errorenum.EmailIsRequired
	}

	if !emailRe.MatchString(email) {
		return errorenum.InvalidEmailAddress.Var(email)
	}

	return nil
}

func (r SubscriberEmail) String() string {
	return string(r)
}
