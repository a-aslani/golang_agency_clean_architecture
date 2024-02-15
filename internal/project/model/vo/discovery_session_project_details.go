package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"strings"
)

type DiscoverySessionProjectDetails string

const projectDetailsMaxLen = 500

func (r DiscoverySessionProjectDetails) Validate() error {

	details := strings.TrimSpace(r.String())

	if details == "" {
		return errorenum.ProjectDetailsIsRequired
	}

	if len(details) > projectDetailsMaxLen {
		return errorenum.MaxLenErr.Var("details", projectDetailsMaxLen, len(details))
	}

	return nil
}

func (r DiscoverySessionProjectDetails) String() string {
	return string(r)
}
