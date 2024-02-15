package discovery_session_request

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/notification"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/recaptcha"
)

type Outport interface {
	contract.Repository
	recaptcha.Recaptcha
	notification.TelegramBot
}
