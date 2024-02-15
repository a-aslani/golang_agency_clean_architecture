package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"strings"
)

type ContactFormID string

func NewContactFormID(id string) (ContactFormID, error) {
	if strings.TrimSpace(id) == "" {
		return "", errorenum.ObjectIdCanNotBeEmpty
	}

	var obj = ContactFormID(id)
	return obj, nil
}

func (r ContactFormID) String() string {
	return string(r)
}
