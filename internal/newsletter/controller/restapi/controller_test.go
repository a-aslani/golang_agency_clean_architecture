package restapi

import (
	"bytes"
	"encoding/json"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/usecase/subscribe"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubscribeHandler(t *testing.T) {

	p := "/newsletter/v1/subscribers"

	router := primaryDriver.Router

	router.POST(p, primaryDriver.subscribeHandler())

	fakeEmail := faker.Email()

	type testcase struct {
		name       string
		reqParams  subscribeHandlerRequest
		success    bool
		statusCode int
		msg        string
		errMsg     string
	}

	testcases := []testcase{
		{
			name: "successfully request",
			reqParams: subscribeHandlerRequest{
				Email: fakeEmail,
			},
			success:    true,
			statusCode: http.StatusOK,
			msg:        subscribe.SuccessResponseMessage,
		},
		{
			name:       "required email",
			statusCode: http.StatusBadRequest,
			errMsg:     errorenum.EmailIsRequired.Error(),
		},
		{
			name: "already email address",
			reqParams: subscribeHandlerRequest{
				Email: fakeEmail,
			},
			statusCode: http.StatusBadRequest,
			errMsg:     errorenum.ThisEmailAddressIsAlreadyUsed.Var(fakeEmail).Error(),
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
				name: "invalid email address",
				reqParams: subscribeHandlerRequest{
					Email: email,
				},
				statusCode: http.StatusBadRequest,
				errMsg:     errorenum.InvalidEmailAddress.Var(email).Error(),
			},
		)
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(tc.reqParams)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, p, &buf)
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			require.Equal(t, tc.statusCode, rr.Code)
			require.NotEmpty(t, rr.Body)

			var res payload.Response

			err = json.Unmarshal(rr.Body.Bytes(), &res)
			require.NoError(t, err)

			require.Equal(t, tc.success, res.Success)

			if res.Success {
				data := res.Data.(map[string]interface{})
				require.Equal(t, tc.msg, data["message"])
			} else {
				require.Equal(t, tc.errMsg, res.ErrorMessage)
			}

		})
	}

}
