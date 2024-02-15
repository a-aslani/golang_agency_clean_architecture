package calendly_discovery_session_on_scheduled

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/notification"
)

type Outport interface {
	contract.Repository
	notification.TelegramBot
}
