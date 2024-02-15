package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"strings"
)

type RoleID string

func NewRoleID(id string) (RoleID, error) {

	if strings.TrimSpace(id) == "" {
		return "", errorenum.ObjectIDCanNotBeEmpty
	}

	return RoleID(id), nil
}

func (r RoleID) String() string {
	return string(r)
}
