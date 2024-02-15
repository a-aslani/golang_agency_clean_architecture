package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRoleID(t *testing.T) {

	t.Run("should create role id without any errors", func(t *testing.T) {
		randStr := uuid.New().String()
		id, err := NewRoleID(randStr)
		require.NoError(t, err)
		require.Equal(t, randStr, id.String())
	})

	t.Run("should showing error when sending empty value", func(t *testing.T) {
		randStr := ""
		id, err := NewRoleID(randStr)
		require.EqualError(t, err, errorenum.ObjectIDCanNotBeEmpty.Error())
		require.Equal(t, "", id.String())
	})
}
