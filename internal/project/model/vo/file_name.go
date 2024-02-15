package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"strings"
)

type FileName string

func (r FileName) Validate() error {
	if strings.TrimSpace(r.String()) == "" {
		return errorenum.FileNameIsRequired
	}

	return nil
}

func (r FileName) String() string {
	return string(r)
}
