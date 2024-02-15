package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"strings"
)

type FileID string

func NewFileID(id string) (FileID, error) {

	if strings.TrimSpace(id) == "" {
		return "", errorenum.ObjectIdCanNotBeEmpty
	}

	var obj = FileID(id)

	return obj, nil
}

func (r FileID) String() string {
	return string(r)
}
