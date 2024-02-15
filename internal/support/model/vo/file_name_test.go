package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileName(t *testing.T) {

	t.Run("should showing error when sending empty string value", func(t *testing.T) {

		name := FileName("")
		err := name.Validate()
		require.EqualError(t, err, errorenum.FileNameIsRequired.Error())

	})

	t.Run("should creating name without any error", func(t *testing.T) {

		n := faker.Name()

		name := FileName(n)

		err := name.Validate()
		require.NoError(t, err)
		require.Equal(t, name.String(), n)
	})
}
