package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"time"
)

type DiscoverySessionDate time.Time

func (r DiscoverySessionDate) Validate() error {

	if r.Time().Unix() <= time.Now().Add(24*time.Hour).Unix() {
		return errorenum.InvalidDiscoverySessionDate
	}

	return nil
}

func (r DiscoverySessionDate) Time() time.Time {
	return time.Time(r)
}
