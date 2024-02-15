package recaptcha

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	MissingInputSecret   = "google reCAPTCHA: the secret parameter is missing. refresh the page and try again"
	InvalidInputSecret   = "google reCAPTCHA: the secret parameter is invalid or malformed. refresh the page and try again"
	MissingInputResponse = "google reCAPTCHA: the response parameter is missing. refresh the page and try again"
	InvalidInputResponse = "google reCAPTCHA: the response parameter is invalid or malformed. refresh the page and try again"
	BadRequest           = "google reCAPTCHA: the request is invalid or malformed. refresh the page and try again"
	TimeoutOrDuplicate   = "google reCAPTCHA: the response is no longer valid: either is too old or has been used previously. refresh the page and try again"
)

type SiteVerifyResponse struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes"`
}

type GoogleRecaptcha struct{}

func NewGoogleRecaptcha() *GoogleRecaptcha {
	return &GoogleRecaptcha{}
}

func (r *GoogleRecaptcha) SiteVerify(ctx context.Context, secret, token string) error {

	requestURL := "https://www.google.com/recaptcha/api/siteverify"

	form := url.Values{}
	form.Add("secret", secret)
	form.Add("response", token)

	req, err := http.NewRequest(http.MethodPost, requestURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var resBody SiteVerifyResponse

	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return err
	}

	if !resBody.Success {

		endMsg := "refresh the page and try again."

		if len(resBody.ErrorCodes) > 0 {
			switch resBody.ErrorCodes[0] {
			case "missing-input-secret":
				return errors.New(MissingInputSecret)
			case "invalid-input-secret":
				return errors.New(InvalidInputSecret)
			case "missing-input-response":
				return errors.New(MissingInputResponse)
			case "invalid-input-response":
				return errors.New(InvalidInputResponse)
			case "bad-request":
				return errors.New(BadRequest)
			case "timeout-or-duplicate":
				return errors.New(TimeoutOrDuplicate)
			default:
				return fmt.Errorf("the reCAPTCHA error: %s. %s", resBody.ErrorCodes[0], endMsg)
			}
		}

		return fmt.Errorf("google reCAPTCHA unknown error. %s", endMsg)
	}

	return nil
}
