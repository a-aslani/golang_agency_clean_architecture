package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoleCode(t *testing.T) {

	t.Run("should create role code without any errors", func(t *testing.T) {
		code := uuid.New().String()
		v := RoleCode(code)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, code, v.String())
	})

	t.Run("should showing error when sending empty value", func(t *testing.T) {
		code := ""
		v := RoleCode(code)
		err := v.Validate()
		require.EqualError(t, err, errorenum.RoleCodeIsRequired.Error())
		require.Equal(t, "", v.String())
	})

}
