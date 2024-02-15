package subscribe

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
)

type Inport = framework.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.SubscriberCreateRequest
}

type InportResponse struct {
	ID      vo.SubscriberID `json:"id"`
	Message string          `json:"message"`
}
