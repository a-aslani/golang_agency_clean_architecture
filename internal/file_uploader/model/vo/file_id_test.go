package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/errorenum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewFileID(t *testing.T) {

	t.Run("should create id without any error", func(t *testing.T) {

		idStr := uuid.New().String()

		id, err := NewFileID(idStr)
		require.NoError(t, err)
		require.NotEmpty(t, id)
		require.Equal(t, idStr, id.String())
	})

	t.Run("should showing error when sending empty string", func(t *testing.T) {

		idStr := ""

		id, err := NewFileID(idStr)
		require.EqualError(t, err, errorenum.ObjectIdCanNotBeEmpty.Error())
		require.Empty(t, id)

	})
}
