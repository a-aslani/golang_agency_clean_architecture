package upload_file

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
)

type Inport = framework.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.FileCreateRequest
}

type InportResponse struct {
	ID      vo.FileID `json:"id"`
	Message string    `json:"message"`
	Path    string    `json:"path"`
}
