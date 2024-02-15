package contract

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/entity"
)

//go:generate go run go.uber.org/mock/mockgen -destination mocks/repository_mock.go -package mocks github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/contract Repository
type Repository interface {
	SaveSubscriber(ctx context.Context, obj *entity.Subscriber) error
	FindOneSubscriberByID(ctx context.Context, subscriberID string) (*entity.Subscriber, error)
	FindOneSubscriberByEmail(ctx context.Context, email string) (*entity.Subscriber, error)
}
