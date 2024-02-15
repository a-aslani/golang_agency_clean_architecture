package subscribe

import (
	"context"
	"database/sql"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/contract/mocks"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/vo"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestRunSubscriberCreateInteractor_Execute(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockrepo := mocks.NewMockRepository(ctrl)

	usecase := NewUsecase(struct {
		contract.Repository
	}{
		mockrepo,
	})

	t.Run("should create and save subscriber without any error", func(t *testing.T) {

		id := uuid.New().String()
		now := time.Now()
		email := faker.Email()

		req := InportRequest{SubscriberCreateRequest: entity.SubscriberCreateRequest{
			ID:    id,
			Now:   now,
			Email: vo.SubscriberEmail(email),
		}}

		obj := &entity.Subscriber{
			ID:      vo.SubscriberID(id),
			Email:   vo.SubscriberEmail(email),
			Created: now,
		}

		mockrepo.EXPECT().FindOneSubscriberByEmail(gomock.Any(), gomock.Eq(req.Email.String())).Times(1).Return(nil, sql.ErrNoRows)

		mockrepo.EXPECT().SaveSubscriber(gomock.Any(), gomock.Eq(obj)).Times(1).Return(nil)

		res, err := usecase.Execute(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, res.ID.String(), obj.ID.String())
	})

	t.Run("should showing error when sending invalid input", func(t *testing.T) {

		t.Run("empty object id", func(t *testing.T) {
			id := ""
			now := time.Now()
			email := faker.Email()

			req := InportRequest{SubscriberCreateRequest: entity.SubscriberCreateRequest{
				ID:    id,
				Now:   now,
				Email: vo.SubscriberEmail(email),
			}}

			obj := &entity.Subscriber{
				ID:      vo.SubscriberID(id),
				Email:   vo.SubscriberEmail(email),
				Created: now,
			}

			mockrepo.EXPECT().FindOneSubscriberByEmail(gomock.Any(), gomock.Eq(req.Email.String())).Times(1).Return(nil, sql.ErrNoRows)

			mockrepo.EXPECT().SaveSubscriber(gomock.Any(), gomock.Eq(obj)).Times(0)

			_, err := usecase.Execute(context.Background(), req)
			require.EqualError(t, err, errorenum.ObjectIDCanNotBeEmpty.Error())
		})

		t.Run("empty email address", func(t *testing.T) {
			id := uuid.New().String()
			now := time.Now()
			email := ""

			req := InportRequest{SubscriberCreateRequest: entity.SubscriberCreateRequest{
				ID:    id,
				Now:   now,
				Email: vo.SubscriberEmail(email),
			}}

			obj := &entity.Subscriber{
				ID:      vo.SubscriberID(id),
				Email:   vo.SubscriberEmail(email),
				Created: now,
			}

			mockrepo.EXPECT().FindOneSubscriberByEmail(gomock.Any(), gomock.Eq(req.Email.String())).Times(1).Return(nil, sql.ErrNoRows)

			mockrepo.EXPECT().SaveSubscriber(gomock.Any(), gomock.Eq(obj)).Times(0)

			_, err := usecase.Execute(context.Background(), req)
			require.EqualError(t, err, errorenum.EmailIsRequired.Error())
		})

		t.Run("invalid email address", func(t *testing.T) {
			id := uuid.New().String()
			now := time.Now()
			email := "aaa"

			req := InportRequest{SubscriberCreateRequest: entity.SubscriberCreateRequest{
				ID:    id,
				Now:   now,
				Email: vo.SubscriberEmail(email),
			}}

			obj := &entity.Subscriber{
				ID:      vo.SubscriberID(id),
				Email:   vo.SubscriberEmail(email),
				Created: now,
			}

			mockrepo.EXPECT().FindOneSubscriberByEmail(gomock.Any(), gomock.Eq(req.Email.String())).Times(1).Return(nil, sql.ErrNoRows)

			mockrepo.EXPECT().SaveSubscriber(gomock.Any(), gomock.Eq(obj)).Times(0)

			_, err := usecase.Execute(context.Background(), req)
			require.EqualError(t, err, errorenum.InvalidEmailAddress.Var(email).Error())
		})

		t.Run("email address is already used", func(t *testing.T) {
			id := uuid.New().String()
			now := time.Now()
			email := faker.Email()

			req := InportRequest{SubscriberCreateRequest: entity.SubscriberCreateRequest{
				ID:    id,
				Now:   now,
				Email: vo.SubscriberEmail(email),
			}}

			obj := &entity.Subscriber{
				ID:      vo.SubscriberID(id),
				Email:   vo.SubscriberEmail(email),
				Created: now,
			}

			mockrepo.EXPECT().FindOneSubscriberByEmail(gomock.Any(), gomock.Eq(req.Email.String())).Times(1).Return(obj, errorenum.ThisEmailAddressIsAlreadyUsed)

			mockrepo.EXPECT().SaveSubscriber(gomock.Any(), gomock.Eq(obj)).Times(0)

			_, err := usecase.Execute(context.Background(), req)
			require.EqualError(t, err, errorenum.ThisEmailAddressIsAlreadyUsed.Var(email).Error())
		})

	})
}
