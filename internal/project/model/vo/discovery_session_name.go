package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"strings"
)

const nameMaxLen = 30

type DiscoverySessionName string

func (r DiscoverySessionName) Validate() error {

	name := strings.TrimSpace(r.String())

	if name == "" {
		return errorenum.NameIsRequired
	}

	if len(name) > nameMaxLen {
		return errorenum.MaxLenErr.Var("name", nameMaxLen, len(name))
	}

	return nil
}

func (r DiscoverySessionName) String() string {
	return string(r)
}
