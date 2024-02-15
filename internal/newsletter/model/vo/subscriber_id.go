package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/errorenum"
	"strings"
)

type SubscriberID string

func NewSubscriberID(id string) (SubscriberID, error) {

	if strings.TrimSpace(id) == "" {
		return "", errorenum.ObjectIDCanNotBeEmpty
	}

	var obj = SubscriberID(id)

	return obj, nil
}

func (r SubscriberID) String() string {
	return string(r)
}
