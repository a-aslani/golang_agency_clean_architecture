package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewContactFormID(t *testing.T) {

	t.Run("should showing error when sending empty id", func(t *testing.T) {

		_, err := NewContactFormID("")
		require.EqualError(t, err, errorenum.ObjectIdCanNotBeEmpty.Error())
	})

	t.Run("should create form ID without any error", func(t *testing.T) {

		uid := uuid.New().String()
		id, err := NewContactFormID(uid)
		require.NoError(t, err)
		require.Equal(t, id.String(), uid)
	})
}
