package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiscoverySessionProjectDetails(t *testing.T) {

	t.Run("should showing error when sending empty value", func(t *testing.T) {

		v := DiscoverySessionProjectDetails("")
		err := v.Validate()
		require.EqualError(t, err, errorenum.ProjectDetailsIsRequired.Error())
	})

	t.Run("should create project details value object without any error", func(t *testing.T) {

		details := faker.Sentence()
		v := DiscoverySessionProjectDetails(details)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, details, v.String())
	})

	t.Run("should showing max length error when sending long text", func(t *testing.T) {

		txt := ""
		for i := 0; i <= projectDetailsMaxLen; i++ {
			txt += "a"
		}

		v := DiscoverySessionProjectDetails(txt)
		err := v.Validate()
		require.EqualError(t, err, errorenum.MaxLenErr.Var("details", projectDetailsMaxLen, len(txt)).Error())
	})
}
