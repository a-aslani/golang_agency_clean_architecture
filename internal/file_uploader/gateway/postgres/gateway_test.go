package postgres

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/vo"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGateway_SaveFilePath(t *testing.T) {

	id := uuid.New().String()

	obj := entity.File{
		ID:      vo.FileID(id),
		Name:    vo.FileName(faker.Name()),
		Path:    vo.FilePath(faker.URL()),
		Created: time.Now(),
	}

	err := datasource.SaveFilePath(context.Background(), &obj)
	require.NoError(t, err)
}
