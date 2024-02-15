package entity

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/vo"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewDiscoverySession(t *testing.T) {

	t.Run("should create discovery session entity without any error", func(t *testing.T) {

		req := DiscoverySessionCreateRequest{
			ID:             uuid.New().String(),
			Now:            time.Now(),
			Name:           vo.DiscoverySessionName(faker.Name()),
			Email:          vo.DiscoverySessionEmail(faker.Email()),
			Date:           vo.DiscoverySessionDate(time.Now().Add(72 * time.Hour)),
			ProjectDetails: vo.DiscoverySessionProjectDetails(faker.Sentence()),
		}

		obj, err := NewDiscoverySession(req)
		require.NoError(t, err)
		require.Equal(t, req.Email, obj.Email)
		require.Equal(t, req.Name, obj.Name)
		require.Equal(t, req.ProjectDetails, obj.ProjectDetails)
	})

	t.Run("should showing error when sending invalid inputs", func(t *testing.T) {

		testcases := []struct {
			req DiscoverySessionCreateRequest
			err error
		}{
			{
				req: DiscoverySessionCreateRequest{
					ID:             uuid.New().String(),
					Now:            time.Now(),
					Name:           vo.DiscoverySessionName(faker.Name()),
					Email:          vo.DiscoverySessionEmail(faker.Email()),
					Date:           vo.DiscoverySessionDate(time.Now()),
					ProjectDetails: vo.DiscoverySessionProjectDetails(faker.Sentence()),
				},
				err: errorenum.InvalidDiscoverySessionDate,
			},
			{
				req: DiscoverySessionCreateRequest{
					ID:             uuid.New().String(),
					Now:            time.Now(),
					Name:           vo.DiscoverySessionName(faker.Name()),
					Email:          vo.DiscoverySessionEmail(""),
					Date:           vo.DiscoverySessionDate(time.Now().Add(72 * time.Hour)),
					ProjectDetails: vo.DiscoverySessionProjectDetails(faker.Sentence()),
				},
				err: errorenum.EmailIsRequired,
			},
			{
				req: DiscoverySessionCreateRequest{
					ID:             uuid.New().String(),
					Now:            time.Now(),
					Name:           vo.DiscoverySessionName(faker.Name()),
					Email:          vo.DiscoverySessionEmail("anything"),
					Date:           vo.DiscoverySessionDate(time.Now().Add(72 * time.Hour)),
					ProjectDetails: vo.DiscoverySessionProjectDetails(faker.Sentence()),
				},
				err: errorenum.InvalidEmailAddress.Var("anything"),
			},
			{
				req: DiscoverySessionCreateRequest{
					ID:             "",
					Now:            time.Now(),
					Name:           vo.DiscoverySessionName(faker.Name()),
					Email:          vo.DiscoverySessionEmail(faker.Email()),
					Date:           vo.DiscoverySessionDate(time.Now().Add(72 * time.Hour)),
					ProjectDetails: vo.DiscoverySessionProjectDetails(faker.Sentence()),
				},
				err: errorenum.ObjectIDCanNotBeEmpty,
			},
			{
				req: DiscoverySessionCreateRequest{
					ID:             uuid.New().String(),
					Now:            time.Now(),
					Name:           vo.DiscoverySessionName(""),
					Email:          vo.DiscoverySessionEmail(faker.Email()),
					Date:           vo.DiscoverySessionDate(time.Now().Add(72 * time.Hour)),
					ProjectDetails: vo.DiscoverySessionProjectDetails(faker.Sentence()),
				},
				err: errorenum.NameIsRequired,
			},
			{
				req: DiscoverySessionCreateRequest{
					ID:             uuid.New().String(),
					Now:            time.Now(),
					Name:           vo.DiscoverySessionName(faker.Name()),
					Email:          vo.DiscoverySessionEmail(faker.Email()),
					Date:           vo.DiscoverySessionDate(time.Now().Add(72 * time.Hour)),
					ProjectDetails: vo.DiscoverySessionProjectDetails(""),
				},
				err: errorenum.ProjectDetailsIsRequired,
			},
		}

		for _, tc := range testcases {

			obj, err := NewDiscoverySession(tc.req)
			require.EqualError(t, err, tc.err.Error())
			require.Nil(t, obj)
		}

	})

}
