package upload_file

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/contract/mocks"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/vo"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestRunFileCreateInteractor_Execute(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockrepo := mocks.NewMockRepository(ctrl)

	idStr := uuid.New().String()
	id, err := vo.NewFileID(idStr)
	require.NoError(t, err)
	require.Equal(t, idStr, id.String())

	req := InportRequest{
		entity.FileCreateRequest{
			ID:   idStr,
			Name: vo.FileName(faker.Name()),
			Path: vo.FilePath(faker.URL()),
			Now:  time.Now(),
		},
	}

	obj := entity.File{
		ID:      id,
		Name:    req.Name,
		Path:    req.Path,
		Created: req.Now,
	}

	res := InportResponse{
		ID:      id,
		Message: "File has been uploaded",
	}

	mockrepo.EXPECT().SaveFilePath(gomock.Any(), gomock.Eq(&obj)).Times(1).Return(nil)

	usecase := NewUsecase(mockrepo)

	r, err := usecase.Execute(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, res.ID, r.ID)
	require.Equal(t, res.Message, r.Message)
}
