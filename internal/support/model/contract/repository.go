package contract

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/entity"
)

//go:generate go run go.uber.org/mock/mockgen -destination mocks/repository_mock.go -package mocks github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/contract Repository
type Repository interface {
	SaveContactForm(ctx context.Context, obj *entity.ContactForm) error
	FindFilesByIDs(ctx context.Context, ids []string) ([]*entity.File, error)
	SaveFilePath(ctx context.Context, obj *entity.File) error
	FindRolesByCodes(ctx context.Context, codes []string) ([]*entity.Role, error)
	FindChatIdsByRoles(ctx context.Context, roles []*entity.Role) ([]int64, error)
	SaveRole(ctx context.Context, obj *entity.Role) error
	SaveTelegramChatID(ctx context.Context, obj *entity.TelegramChatID) error
}
