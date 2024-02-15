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

func TestNewContactForm(t *testing.T) {

	t.Run("should create without error", func(t *testing.T) {

		files := make([]*File, 0)

		fileCounts := 3

		for i := 0; i < fileCounts; i++ {

			fileReq := FileCreateRequest{
				ID:   uuid.New().String(),
				Name: vo.FileName(faker.Name()),
				Path: vo.FilePath(faker.URL()),
				Now:  time.Now(),
			}

			file, err := NewFile(fileReq)
			require.NoError(t, err)
			require.Equal(t, file.ID.String(), fileReq.ID)
			require.Equal(t, file.Path, fileReq.Path)
			require.Equal(t, file.Name, fileReq.Name)

			files = append(files, file)
		}

		formReq := ContactFormCreateRequest{
			ID:      uuid.New().String(),
			Now:     time.Now(),
			Name:    vo.ContactFormName(faker.Name()),
			Email:   vo.ContactFormEmail(faker.Email()),
			Message: vo.ContactFormMessage(faker.Sentence()),
			Files:   files,
		}

		form, err := NewContactForm(formReq)
		require.NoError(t, err)
		require.Equal(t, form.ID.String(), formReq.ID)
		require.Equal(t, form.Name, formReq.Name)
		require.Equal(t, form.Email, formReq.Email)
		require.Equal(t, form.Message, formReq.Message)
		require.Len(t, form.Files, fileCounts)
		for k, file := range form.Files {
			require.Equal(t, file.ID, files[k].ID)
			require.Equal(t, file.Name, files[k].Name)
			require.Equal(t, file.Path, files[k].Path)
		}
	})

	t.Run("check files validation", func(t *testing.T) {

		testCases := []struct {
			req    FileCreateRequest
			expect func(req FileCreateRequest)
		}{
			{
				req: FileCreateRequest{
					ID:   uuid.New().String(),
					Name: "",
					Path: vo.FilePath(faker.URL()),
					Now:  time.Now(),
				},
				expect: func(req FileCreateRequest) {
					file, err := NewFile(req)
					require.EqualError(t, err, errorenum.FileNameIsRequired.Error())
					require.Nil(t, file)
				},
			},
			{
				req: FileCreateRequest{
					ID:   "",
					Name: vo.FileName(faker.Name()),
					Path: vo.FilePath(faker.URL()),
					Now:  time.Now(),
				},
				expect: func(req FileCreateRequest) {
					file, err := NewFile(req)
					require.EqualError(t, err, errorenum.ObjectIdCanNotBeEmpty.Error())
					require.Nil(t, file)
				},
			},
			{
				req: FileCreateRequest{
					ID:   uuid.New().String(),
					Name: vo.FileName(faker.Name()),
					Path: "",
					Now:  time.Now(),
				},
				expect: func(req FileCreateRequest) {
					file, err := NewFile(req)
					require.EqualError(t, err, errorenum.FilePathIsRequired.Error())
					require.Nil(t, file)
				},
			},
		}

		for _, tc := range testCases {
			tc.expect(tc.req)
		}

	})

	t.Run("check value objects validations", func(t *testing.T) {

		//files := make([]*File, 0)
		//
		//fileCounts := 3
		//
		//for i := 0; i < fileCounts; i++ {
		//
		//	fileReq := FileCreateRequest{
		//		ID: uuid.New().String(),
		//		Name: vo.FileName(faker.Name()),
		//		Path: vo.FilePath(faker.URL()),
		//		Now:  time.Now(),
		//	}
		//
		//	file, err := NewFile(fileReq)
		//	require.NoError(t, err)
		//	require.Equal(t, file.ID.String(), fileReq.ID)
		//	require.Equal(t, file.Path, fileReq.Path)
		//	require.Equal(t, file.Name, fileReq.Name)
		//
		//	files = append(files, file)
		//}

		fileId, _ := vo.NewFileID(uuid.New().String())
		tm := time.Now()
		uuidStr := uuid.New().String()
		name := vo.ContactFormName(faker.Name())
		email := vo.ContactFormEmail(faker.Email())
		message := vo.ContactFormMessage(faker.Sentence())

		testCases := []struct {
			req    ContactFormCreateRequest
			expect func(req ContactFormCreateRequest)
		}{
			{
				req: ContactFormCreateRequest{
					ID:      "",
					Now:     tm,
					Name:    name,
					Email:   email,
					Message: message,
					Files:   []*File{},
				},
				expect: func(req ContactFormCreateRequest) {
					form, err := NewContactForm(req)
					require.EqualError(t, err, errorenum.ObjectIdCanNotBeEmpty.Error())
					require.Nil(t, form)
				},
			},
			{
				req: ContactFormCreateRequest{
					ID:      uuidStr,
					Now:     tm,
					Name:    "",
					Email:   email,
					Message: message,
					Files:   []*File{},
				},
				expect: func(req ContactFormCreateRequest) {
					form, err := NewContactForm(req)
					require.EqualError(t, err, errorenum.ContactFormNameIsRequired.Error())
					require.Nil(t, form)
				},
			},
			{
				req: ContactFormCreateRequest{
					ID:      uuidStr,
					Now:     tm,
					Name:    name,
					Email:   "",
					Message: message,
					Files:   []*File{},
				},
				expect: func(req ContactFormCreateRequest) {
					form, err := NewContactForm(req)
					require.EqualError(t, err, errorenum.ContactFormEmailIsRequired.Error())
					require.Nil(t, form)
				},
			},
			{
				req: ContactFormCreateRequest{
					ID:      uuidStr,
					Now:     tm,
					Name:    name,
					Email:   email,
					Message: "",
					Files:   []*File{},
				},
				expect: func(req ContactFormCreateRequest) {
					form, err := NewContactForm(req)
					require.EqualError(t, err, errorenum.ContactFormMessageIsRequired.Error())
					require.Nil(t, form)
				},
			},
			{
				req: ContactFormCreateRequest{
					ID:      uuidStr,
					Now:     tm,
					Name:    name,
					Email:   email,
					Message: message,
					Files: []*File{
						{
							ID:      fileId,
							Name:    vo.FileName(faker.Name()),
							Path:    "",
							Created: time.Now(),
						},
					},
				},
				expect: func(req ContactFormCreateRequest) {
					form, err := NewContactForm(req)
					require.EqualError(t, err, errorenum.FilePathIsRequired.Error())
					require.Nil(t, form)
				},
			},
			{
				req: ContactFormCreateRequest{
					ID:      uuidStr,
					Now:     tm,
					Name:    name,
					Email:   email,
					Message: message,
					Files: []*File{
						{
							ID:      fileId,
							Name:    "",
							Path:    vo.FilePath(faker.URL()),
							Created: time.Now(),
						},
					},
				},
				expect: func(req ContactFormCreateRequest) {
					form, err := NewContactForm(req)
					require.EqualError(t, err, errorenum.FileNameIsRequired.Error())
					require.Nil(t, form)
				},
			},
		}

		for _, tc := range testCases {
			tc.expect(tc.req)
		}
	})
}
