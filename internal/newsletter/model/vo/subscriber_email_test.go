package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSubscriberEmail(t *testing.T) {

	t.Run("should showing error when sending empty string", func(t *testing.T) {

		value := SubscriberEmail("")
		err := value.Validate()
		require.EqualError(t, err, errorenum.EmailIsRequired.Error())
		require.Equal(t, "", value.String())

	})

	t.Run("should showing invalid email address error when sending wrong email", func(t *testing.T) {

		testCases := []struct {
			email string
		}{
			{
				email: "aa",
			},
			{
				email: "a@b",
			},
			{
				email: "ab.com",
			},
			{
				email: "@b.com",
			},
		}

		for _, v := range testCases {
			value := SubscriberEmail(v.email)
			err := value.Validate()
			require.EqualError(t, err, errorenum.InvalidEmailAddress.Var(v.email).Error())
			require.Equal(t, v.email, value.String())
		}

	})

	t.Run("should create email address without any error", func(t *testing.T) {

		email := faker.Email()

		value := SubscriberEmail(email)
		err := value.Validate()
		require.NoError(t, err)
		require.Equal(t, email, value.String())

	})
}
