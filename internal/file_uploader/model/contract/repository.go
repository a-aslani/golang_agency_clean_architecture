package contract

import (
	"context"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/entity"
)

//go:generate go run go.uber.org/mock/mockgen -destination mocks/repository_mock.go -package mocks github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/contract Repository
type Repository interface {
	SaveFilePath(ctx context.Context, obj *entity.File) error
}
