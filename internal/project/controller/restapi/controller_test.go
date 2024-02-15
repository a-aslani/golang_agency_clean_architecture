package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/usecase/calendly_discovery_session_on_scheduled"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/usecase/discovery_session_request"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCalendlyDiscoverySessionOnScheduledHandler(t *testing.T) {

	path := "/project/v1/discovery-session/calendly"

	router := primaryDriver.Router

	router.GET(path, primaryDriver.calendlyDiscoverySessionOnScheduledHandler())

	var params bytes.Buffer
	err := json.NewEncoder(&params).Encode(calendly_discovery_session_on_scheduled.InportRequest{})

	req, err := http.NewRequest(http.MethodGet, path, &params)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.NotEmpty(t, rr.Body)

}

func TestDiscoverySessionRequestHandler(t *testing.T) {

	path := "/project/v1/discovery-session"

	router := primaryDriver.Router

	router.POST(path, primaryDriver.discoverySessionRequestHandler())

	timeParamLayout := "2006-01-02T15:04:05.000Z"

	type testcase struct {
		Name       string
		Params     discoverySessionRequestHandlerRequest
		Success    bool
		StatusCode int
		SuccessMsg string
		ErrorMsg   string
	}

	testcases := []testcase{
		{
			Name: "success",
			Params: discoverySessionRequestHandlerRequest{
				Name:           faker.Name(),
				Email:          faker.Email(),
				Date:           time.Now().Add(48 * time.Hour).Format(timeParamLayout),
				ProjectDetails: faker.Sentence(),
				Files: []string{
					faker.URL(),
					faker.URL(),
					faker.URL(),
				},
			},
			Success:    true,
			SuccessMsg: discovery_session_request.SuccessResponseMessage,
			StatusCode: http.StatusOK,
		},
		{
			Name: "bad request name",
			Params: discoverySessionRequestHandlerRequest{
				Name:           "",
				Email:          faker.Email(),
				Date:           time.Now().Add(48 * time.Hour).Format(timeParamLayout),
				ProjectDetails: faker.Sentence(),
				Files:          []string{},
			},
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   errorenum.NameIsRequired.Error(),
		},
		{
			Name: "bad request email",
			Params: discoverySessionRequestHandlerRequest{
				Name:           faker.Name(),
				Email:          "",
				Date:           time.Now().Add(48 * time.Hour).Format(timeParamLayout),
				ProjectDetails: faker.Sentence(),
				Files:          []string{},
			},
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   errorenum.EmailIsRequired.Error(),
		},
		{
			Name: "bad request detail",
			Params: discoverySessionRequestHandlerRequest{
				Name:           faker.Name(),
				Email:          faker.Email(),
				Date:           time.Now().Add(48 * time.Hour).Format(timeParamLayout),
				ProjectDetails: "",
				Files:          []string{},
			},
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   errorenum.ProjectDetailsIsRequired.Error(),
		},
		{
			Name: "bad request datetime",
			Params: discoverySessionRequestHandlerRequest{
				Name:           faker.Name(),
				Email:          faker.Email(),
				Date:           time.Now().Format(timeParamLayout),
				ProjectDetails: faker.Sentence(),
				Files:          []string{},
			},
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   errorenum.InvalidDiscoverySessionDate.Error(),
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
				Params: discoverySessionRequestHandlerRequest{
					Name:           faker.Name(),
					Email:          email,
					Date:           time.Now().Add(48 * time.Hour).Format(timeParamLayout),
					ProjectDetails: faker.Sentence(),
					Files:          []string{},
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
