package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoleName(t *testing.T) {

	t.Run("should create role name without any errors", func(t *testing.T) {
		name := faker.Name()
		v := RoleName(name)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, v.String(), name)
	})

	t.Run("should showing error when sending empty value", func(t *testing.T) {
		name := ""
		v := RoleName(name)
		err := v.Validate()
		require.EqualError(t, err, errorenum.RoleNameIsRequired.Error())
		require.Equal(t, v.String(), "")
	})

}
