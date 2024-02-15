package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContactFormEmail(t *testing.T) {

	t.Run("should showing error when sending empty email", func(t *testing.T) {

		v := ContactFormEmail("")
		err := v.Validate()
		require.EqualError(t, err, errorenum.ContactFormEmailIsRequired.Error())
	})

	t.Run("should create email value object without error", func(t *testing.T) {

		email := faker.Email()

		v := ContactFormEmail(email)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, v.String(), email)
	})
}
