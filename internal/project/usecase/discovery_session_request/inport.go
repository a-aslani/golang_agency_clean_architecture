package discovery_session_request

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"time"
)

type Inport = framework.Inport[InportRequest, InportResponse]

type InportRequest struct {
	UUID           string
	Now            time.Time
	Name           vo.DiscoverySessionName
	Email          vo.DiscoverySessionEmail
	Date           vo.DiscoverySessionDate
	ProjectDetails vo.DiscoverySessionProjectDetails
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
