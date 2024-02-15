package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/errorenum"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilePath(t *testing.T) {

	t.Run("should create file path without any error", func(t *testing.T) {

		p := faker.URL()
		path := FilePath(p)

		err := path.Validate()
		require.NoError(t, err)
		require.Equal(t, p, path.String())
	})

	t.Run("should showing error when sending empty string value for path parameter", func(t *testing.T) {

		path := FilePath("")

		err := path.Validate()
		require.EqualError(t, err, errorenum.FilePathIsRequired.Error())
	})
}
