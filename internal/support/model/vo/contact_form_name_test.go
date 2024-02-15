package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContactFormName(t *testing.T) {

	t.Run("should showing error when sending empty name", func(t *testing.T) {
		v := ContactFormName("")
		err := v.Validate()
		require.EqualError(t, err, errorenum.ContactFormNameIsRequired.Error())
	})

	t.Run("should create without error", func(t *testing.T) {
		name := faker.Name()
		v := ContactFormName(name)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, v.String(), name)
	})
}
