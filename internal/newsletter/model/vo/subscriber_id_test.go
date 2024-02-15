package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/errorenum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSubscriberID(t *testing.T) {

	t.Run("should be create id without any error", func(t *testing.T) {

		id := uuid.New().String()

		value, err := NewSubscriberID(id)
		require.NoError(t, err)
		require.Equal(t, id, value.String())
	})

	t.Run("should showing id can not be empty error when sending blank id", func(t *testing.T) {

		value, err := NewSubscriberID("")
		require.EqualError(t, err, errorenum.ObjectIDCanNotBeEmpty.Error())
		require.Equal(t, SubscriberID(""), value)
	})
}
