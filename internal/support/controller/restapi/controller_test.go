package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/usecase/send_contact_form"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendContactFormHandler(t *testing.T) {

	path := "/support/v1/contact-us"

	router := primaryDriver.Router

	router.POST(path, primaryDriver.sendContactFormHandler())

	type testcase struct {
		Name       string
		Params     sendContactFormHandlerRequest
		Success    bool
		StatusCode int
		SuccessMsg string
		ErrorMsg   string
	}

	testcases := []testcase{
		{
			Name: "success",
			Params: sendContactFormHandlerRequest{
				Name:    faker.Name(),
				Email:   faker.Email(),
				Message: faker.Sentence(),
				Files: []string{
					faker.URL(),
					faker.URL(),
				},
			},
			Success:    true,
			StatusCode: http.StatusOK,
			SuccessMsg: send_contact_form.SuccessMessageResponse,
		},
		{
			Name: "bad request name",
			Params: sendContactFormHandlerRequest{
				Name:    "",
				Email:   faker.Email(),
				Message: faker.Sentence(),
				Files:   []string{},
			},
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   errorenum.ContactFormNameIsRequired.Error(),
		},
		{
			Name: "bad request email",
			Params: sendContactFormHandlerRequest{
				Name:    faker.Name(),
				Email:   "",
				Message: faker.Sentence(),
				Files:   []string{},
			},
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   errorenum.ContactFormEmailIsRequired.Error(),
		},
		{
			Name: "bad request message",
			Params: sendContactFormHandlerRequest{
				Name:    faker.Name(),
				Email:   faker.Email(),
				Message: "",
				Files:   []string{},
			},
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   errorenum.ContactFormMessageIsRequired.Error(),
		},
	}

	invalidEmails := []string{
		"plainaddress",
		"#@%^%#$@#$@#.com",
		"@example.com",
		"Joe Smith <email@example.com>",
		"email.example.com",
		"email@example@example.com",
		"email@example.com (Joe Smith)",
		"email@example",
	}

	for _, email := range invalidEmails {
		testcases = append(
			testcases,
			testcase{
				Name: fmt.Sprintf("invalid email address %s", email),
				Params: sendContactFormHandlerRequest{
					Name:    faker.Name(),
					Email:   email,
					Message: faker.Sentence(),
					Files:   []string{},
				},
				StatusCode: http.StatusBadRequest,
				ErrorMsg:   errorenum.InvalidEmailAddress.Var(email).Error(),
			},
		)
	}

	for _, tc := range testcases {

		t.Run(tc.Name, func(t *testing.T) {
			var params bytes.Buffer
			err := json.NewEncoder(&params).Encode(tc.Params)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, path, &params)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			require.Equal(t, tc.StatusCode, rr.Code)
			require.NotEmpty(t, rr.Body)

			var res payload.Response

			err = json.Unmarshal(rr.Body.Bytes(), &res)
			require.NoError(t, err)

			require.Equal(t, tc.Success, res.Success)

			if res.Success {
				data := res.Data.(map[string]interface{})
				require.Equal(t, tc.SuccessMsg, data["message"])
			} else {
				require.Equal(t, tc.ErrorMsg, res.ErrorMessage)
			}
		})

	}

}
