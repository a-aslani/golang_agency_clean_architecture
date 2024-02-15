package send_contact_form

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"time"
)

type Inport = framework.Inport[InportRequest, InportResponse]

type InportRequest struct {
	ID             string
	Now            time.Time
	Name           vo.ContactFormName
	Email          vo.ContactFormEmail
	Message        vo.ContactFormMessage
	RecaptchaToken string
	Secret         string
	Files          []string
	APIUrl         string
	TestMode       bool
}

type InportResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
