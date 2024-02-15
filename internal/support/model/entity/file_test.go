package entity

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/vo"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {

	t.Run("should create file without any error", func(t *testing.T) {

		req := FileCreateRequest{
			ID:   uuid.New().String(),
			Name: vo.FileName(faker.Name()),
			Path: vo.FilePath(faker.URL()),
			Now:  time.Now(),
		}

		file, err := NewFile(req)
		require.NoError(t, err)
		require.Equal(t, file.ID.String(), req.ID)
		require.Equal(t, file.Name.String(), req.Name.String())
		require.Equal(t, file.Path.String(), req.Path.String())
	})

	t.Run("should showing error when sending empty ID", func(t *testing.T) {

		req := FileCreateRequest{
			ID:   "",
			Name: vo.FileName(faker.Name()),
			Path: vo.FilePath(faker.URL()),
			Now:  time.Now(),
		}

		file, err := NewFile(req)
		require.EqualError(t, err, errorenum.ObjectIdCanNotBeEmpty.Error())
		require.Nil(t, file)
	})

	t.Run("should showing error when sending empty name", func(t *testing.T) {

		req := FileCreateRequest{
			ID:   uuid.New().String(),
			Name: vo.FileName(""),
			Path: vo.FilePath(faker.URL()),
			Now:  time.Now(),
		}

		file, err := NewFile(req)
		require.EqualError(t, err, errorenum.FileNameIsRequired.Error())
		require.Nil(t, file)
	})

	t.Run("should showing error when sending empty path", func(t *testing.T) {

		req := FileCreateRequest{
			ID:   uuid.New().String(),
			Name: vo.FileName(faker.Name()),
			Path: vo.FilePath(""),
			Now:  time.Now(),
		}

		file, err := NewFile(req)
		require.EqualError(t, err, errorenum.FilePathIsRequired.Error())
		require.Nil(t, file)
	})

}
