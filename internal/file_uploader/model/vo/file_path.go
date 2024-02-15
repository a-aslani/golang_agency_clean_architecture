package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/errorenum"
	"strings"
)

type FilePath string

func (r FilePath) Validate() error {
	if strings.TrimSpace(r.String()) == "" {
		return errorenum.FilePathIsRequired
	}
	return nil
}

func (r FilePath) String() string {
	return string(r)
}
