package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDiscoverySessionID(t *testing.T) {

	t.Run("should showing error when sending empty value", func(t *testing.T) {

		v, err := NewDiscoverySessionID("")
		require.EqualError(t, err, errorenum.ObjectIDCanNotBeEmpty.Error())
		require.Equal(t, "", v.String())
	})

	t.Run("should create ID without any error", func(t *testing.T) {

		id := uuid.New().String()
		v, err := NewDiscoverySessionID(id)
		require.NoError(t, err)
		require.Equal(t, id, v.String())
	})
}
