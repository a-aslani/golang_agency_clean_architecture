package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContactFormMessage(t *testing.T) {

	t.Run("should showing error when ending empty message", func(t *testing.T) {
		v := ContactFormMessage("")
		err := v.Validate()
		require.EqualError(t, err, errorenum.ContactFormMessageIsRequired.Error())
	})

	t.Run("should create without error", func(t *testing.T) {
		msg := faker.Sentence()
		v := ContactFormMessage(msg)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, v.String(), msg)
	})
}
