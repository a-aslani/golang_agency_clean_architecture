package upload_file

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/contract"
)

type Outport interface {
	contract.Repository
}
