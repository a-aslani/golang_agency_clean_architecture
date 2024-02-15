package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewChatIDID(t *testing.T) {
	t.Run("should showing error when sending empty id", func(t *testing.T) {

		_, err := NewTelegramChatIDID("")
		require.EqualError(t, err, errorenum.ObjectIdCanNotBeEmpty.Error())
	})

	t.Run("should create chat ID without any error", func(t *testing.T) {

		uid := uuid.New().String()
		id, err := NewTelegramChatIDID(uid)
		require.NoError(t, err)
		require.Equal(t, id.String(), uid)
	})
}
