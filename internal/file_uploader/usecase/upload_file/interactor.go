package upload_file

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/entity"
)

var AllowedExtensions = []string{"csv", "zip", "txt", "xlx", "xlsx", "xls", "pdf", "doc", "docx", "png", "jpg"}

const (
	Dir                    = "./internal/file_uploader"
	UploadPath             = "/public/upload/files"
	SuccessResponseMessage = "File has been uploaded"
)

type uploadFileInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &uploadFileInteractor{
		outport: outputPort,
	}
}

func (r *uploadFileInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	fileObj, err := entity.NewFile(req.FileCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveFilePath(ctx, fileObj)
	if err != nil {
		return nil, err
	}

	res.ID = fileObj.ID
	res.Message = SuccessResponseMessage
	res.Path = req.Path.String()

	return res, nil
}
