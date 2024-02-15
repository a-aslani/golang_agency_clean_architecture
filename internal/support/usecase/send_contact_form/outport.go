package send_contact_form

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/notification"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/recaptcha"
)

type Outport interface {
	contract.Repository
	recaptcha.Recaptcha
	notification.TelegramBot
}
