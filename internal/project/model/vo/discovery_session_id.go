package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"strings"
)

type DiscoverySessionID string

func NewDiscoverySessionID(id string) (DiscoverySessionID, error) {
	if strings.TrimSpace(id) == "" {
		return "", errorenum.ObjectIDCanNotBeEmpty
	}
	var obj = DiscoverySessionID(id)
	return obj, nil
}

func (r DiscoverySessionID) String() string {
	return string(r)
}
