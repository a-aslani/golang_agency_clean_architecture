package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiscoverySessionName(t *testing.T) {

	t.Run("should showing error when sending empty value", func(t *testing.T) {

		v := DiscoverySessionName("")
		err := v.Validate()
		require.EqualError(t, err, errorenum.NameIsRequired.Error())
	})

	t.Run("should create name value object without any error", func(t *testing.T) {

		name := faker.Name()
		v := DiscoverySessionName(name)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, name, v.String())
	})

	t.Run("should showing max length error when sending long text", func(t *testing.T) {

		txt := ""
		for i := 0; i <= nameMaxLen; i++ {
			txt += "a"
		}

		v := DiscoverySessionName(txt)
		err := v.Validate()
		require.EqualError(t, err, errorenum.MaxLenErr.Var("name", nameMaxLen, len(txt)).Error())
	})
}
