package application

import (
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/controller/restapi"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/gateway/postgres"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/usecase/calendly_discovery_session_on_scheduled"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/usecase/discovery_session_request"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/notification"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/recaptcha"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/token"
)

type project struct{}

func NewProject() framework.Runner {
	return &project{}
}

func (project) Run(cfg *configs.Config) error {

	const appName = "project"

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

	discoverySessionRequestUsecase := discovery_session_request.NewUsecase(struct {
		contract.Repository
		recaptcha.Recaptcha
		notification.TelegramBot
	}{
		postgresRepo,
		googleRecaptcha,
		telegramBot,
	})

	calendlyDiscoverySessionOnScheduledUsecase := calendly_discovery_session_on_scheduled.NewUsecase(struct {
		contract.Repository
		notification.TelegramBot
	}{
		postgresRepo,
		telegramBot,
	})

	primaryDriver.AddUsecase(
		discoverySessionRequestUsecase,
		calendlyDiscoverySessionOnScheduledUsecase,
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
