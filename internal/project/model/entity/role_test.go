package entity

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/vo"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRole(t *testing.T) {

	t.Run("should create role entity without any error", func(t *testing.T) {

		req := RoleCreateRequest{
			ID:   uuid.New().String(),
			Code: vo.RoleCode(uuid.New().String()),
			Name: vo.RoleName(faker.Name()),
		}

		role, err := NewRole(req)
		require.NoError(t, err)
		require.Equal(t, req.ID, role.ID.String())
		require.Equal(t, req.Code, role.Code)
		require.Equal(t, req.Name, role.Name)

	})

	t.Run("should showing error when sending wrong values", func(t *testing.T) {

		testcases := []struct {
			RoleCreateRequest
			error
		}{
			{
				RoleCreateRequest: RoleCreateRequest{
					ID:   "",
					Code: vo.RoleCode(uuid.New().String()),
					Name: vo.RoleName(faker.Name()),
				},
				error: errorenum.ObjectIDCanNotBeEmpty,
			},
			{
				RoleCreateRequest: RoleCreateRequest{
					ID:   uuid.New().String(),
					Code: "",
					Name: vo.RoleName(faker.Name()),
				},
				error: errorenum.RoleCodeIsRequired,
			},
			{
				RoleCreateRequest: RoleCreateRequest{
					ID:   uuid.New().String(),
					Code: vo.RoleCode(uuid.New().String()),
					Name: "",
				},
				error: errorenum.RoleNameIsRequired,
			},
		}

		for _, tc := range testcases {
			role, err := NewRole(tc.RoleCreateRequest)
			require.EqualError(t, err, tc.error.Error())
			require.Nil(t, role)
		}

	})

}
