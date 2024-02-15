package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiscoverySessionEmail(t *testing.T) {

	t.Run("should showing error when sending empty value", func(t *testing.T) {

		v := DiscoverySessionEmail("")
		err := v.Validate()
		require.EqualError(t, err, errorenum.EmailIsRequired.Error())
	})

	t.Run("should showing max length error when sending long text", func(t *testing.T) {

		txt := ""
		for i := 0; i <= emailMaxLen; i++ {
			txt += "a"
		}

		v := DiscoverySessionEmail(txt)
		err := v.Validate()
		require.EqualError(t, err, errorenum.MaxLenErr.Var("email", emailMaxLen, len(txt)).Error())
	})

	t.Run("should showing invalid email address error when sending wrong email format", func(t *testing.T) {

		testcases := []string{"anything", "anything@gmail", "@anything", "any@thing@gmail.com"}

		for _, tc := range testcases {
			v := DiscoverySessionEmail(tc)
			err := v.Validate()
			require.EqualError(t, err, errorenum.InvalidEmailAddress.Var(tc).Error())
		}

	})

	t.Run("should create email value object without any error", func(t *testing.T) {

		email := faker.Email()
		v := DiscoverySessionEmail(email)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, email, v.String())
	})
}
