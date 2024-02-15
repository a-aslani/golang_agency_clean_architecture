package recaptcha

import "context"

//go:generate go run go.uber.org/mock/mockgen -destination mocks/recaptcha_mock.go -package mockrecaptcha github.com/a-aslani/golang_agency_clean_architecture/pkg/recaptcha Recaptcha
type Recaptcha interface {
	SiteVerify(ctx context.Context, secret, token string) error
}
