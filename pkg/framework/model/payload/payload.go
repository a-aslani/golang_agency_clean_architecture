package payload

import (
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
)

type Payload struct {
	Data      any                       `json:"data"`
	Publisher framework.ApplicationData `json:"publisher"`
	TraceID   string                    `json:"traceId"`
}
