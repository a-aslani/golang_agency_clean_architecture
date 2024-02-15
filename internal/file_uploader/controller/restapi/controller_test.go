package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/usecase/upload_file"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUploadFileHandler(t *testing.T) {

	path := "/file-uploader/v1/upload"

	router := primaryDriver.Router

	router.POST(path, primaryDriver.uploadFileHandler())

	testFilePath := "../../../../assets/test_files/test"
	files := make([]string, 0)

	for _, e := range upload_file.AllowedExtensions {
		files = append(files, fmt.Sprintf("%s.%s", testFilePath, e))
	}

	for k, e := range upload_file.AllowedExtensions {

		t.Run(fmt.Sprintf("file with %s extenstion", e), func(t *testing.T) {
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("file", files[k])
			if err != nil {
				t.Fatal(err)
			}
			sample, err := os.Open(files[k])
			if err != nil {
				t.Fatal(err)
			}

			_, err = io.Copy(part, sample)
			require.NoError(t, err)
			require.NoError(t, writer.Close())
			req, err := http.NewRequest(http.MethodPost, path, body)
			require.NoError(t, err)

			req.Header.Set("Content-Type", writer.FormDataContentType())

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)

			var res payload.Response

			err = json.Unmarshal(rr.Body.Bytes(), &res)
			require.NoError(t, err)

			data := res.Data.(map[string]interface{})
			require.Equal(t, upload_file.SuccessResponseMessage, data["message"])
		})
	}

	t.Run("error when sending without any file", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		require.NoError(t, writer.Close())
		req, err := http.NewRequest(http.MethodPost, path, body)
		require.NoError(t, err)

		req.Header.Set("Content-Type", writer.FormDataContentType())

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)

		var res payload.Response

		err = json.Unmarshal(rr.Body.Bytes(), &res)
		require.NoError(t, err)

		require.Equal(t, false, res.Success)
		require.Equal(t, errorenum.FilePathIsRequired.Error(), res.ErrorMessage)

	})

	t.Run("error when sending invalid file", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", fmt.Sprintf("%s.invalid", testFilePath))
		if err != nil {
			t.Fatal(err)
		}
		sample, err := os.Open(fmt.Sprintf("%s.invalid", testFilePath))
		if err != nil {
			t.Fatal(err)
		}

		_, err = io.Copy(part, sample)
		require.NoError(t, err)
		require.NoError(t, writer.Close())

		req, err := http.NewRequest(http.MethodPost, path, body)
		require.NoError(t, err)

		req.Header.Set("Content-Type", writer.FormDataContentType())

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)

		var res payload.Response

		err = json.Unmarshal(rr.Body.Bytes(), &res)
		require.NoError(t, err)

		require.Equal(t, false, res.Success)
		require.Equal(t, errorenum.InvalidTypeError.Var("invalid").Error(), res.ErrorMessage)

	})

}
