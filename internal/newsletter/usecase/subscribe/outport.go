package subscribe

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/contract"
)

type Outport interface {
	contract.Repository
}
