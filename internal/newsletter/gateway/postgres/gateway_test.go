package postgres

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/vo"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomSubscriber(t *testing.T) *entity.Subscriber {
	obj := &entity.Subscriber{
		ID:      vo.SubscriberID(uuid.New().String()),
		Email:   vo.SubscriberEmail(faker.Email()),
		Created: time.Now(),
	}

	err := datasource.SaveSubscriber(context.Background(), obj)
	require.NoError(t, err)
	return obj
}

func TestGateway_SaveSubscriber(t *testing.T) {
	createRandomSubscriber(t)

	t.Run("should showing error when using already registered email address", func(t *testing.T) {

		obj1 := createRandomSubscriber(t)

		obj2 := &entity.Subscriber{
			ID:      vo.SubscriberID(uuid.New().String()),
			Email:   obj1.Email,
			Created: time.Now(),
		}

		err := datasource.SaveSubscriber(context.Background(), obj2)
		require.Error(t, err)

	})
}
