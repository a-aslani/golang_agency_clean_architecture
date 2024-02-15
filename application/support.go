package application

import (
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/controller/restapi"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/gateway/postgres"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/usecase/send_contact_form"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/notification"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/recaptcha"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/token"
)

type support struct{}

func NewSupport() framework.Runner {
	return &support{}
}

func (support) Run(cfg *configs.Config) error {

	const appName = "support"

	appData := framework.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	postgresRepo, err := postgres.NewGateway(log, appData, cfg)
	if err != nil {
		return err
	}

	googleRecaptcha := recaptcha.NewGoogleRecaptcha()

	telegramBot, err := notification.NewTelegramBot(cfg.TelegramBotToken, true)
	if err != nil {
		return err
	}

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)
	primaryDriver.RegisterMetrics(appName)

	primaryDriver.AddUsecase(
		send_contact_form.NewUsecase(struct {
			contract.Repository
			recaptcha.Recaptcha
			notification.TelegramBot
		}{
			postgresRepo,
			googleRecaptcha,
			telegramBot,
		}),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
