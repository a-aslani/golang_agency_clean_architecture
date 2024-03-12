package restapi

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/gateway/postgres"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/usecase/calendly_discovery_session_on_scheduled"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/usecase/discovery_session_request"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/notification"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/recaptcha"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/token"
	"log"
	"os"
	"testing"
)

var primaryDriver *controller

func TestMain(m *testing.M) {

	const appName = "project"

	appData := framework.NewApplicationData(appName)

	var err error

	cfg, err := configs.InitConfig("../../../../config.test.yml")
	if err != nil {
		log.Fatalf("reading config file error: %v", err)
	}

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource, err := postgres.NewGateway(logger.NewSimpleJSONLogger(appData), appData, cfg)
	if err != nil {
		log.Fatalf("connect to the database: %v", err)
	}

	err = datasource.ClearDiscoverySessionFileTable(context.Background())
	if err != nil {
		log.Fatalf("clear database: %v", err)
	}

	err = datasource.ClearDiscoverySessionsTable(context.Background())
	if err != nil {
		log.Fatalf("clear database: %v", err)
	}

	err = datasource.ClearTelegramChatIdsTable(context.Background())
	if err != nil {
		log.Fatalf("clear database: %v", err)
	}

	err = datasource.ClearRolesTable(context.Background())
	if err != nil {
		log.Fatalf("clear database: %v", err)
	}

	primaryDriver = NewController(appData, logger.NewSimpleJSONLogger(appData), cfg, jwtToken).(*controller)

	googleRecaptcha := recaptcha.NewGoogleRecaptcha()

	telegramBot, err := notification.NewTelegramBot(cfg.TelegramBot, true)
	if err != nil {
		log.Fatalf("telegram bot error: %v", err)
	}

	discoverySessionRequestUsecase := discovery_session_request.NewUsecase(struct {
		contract.Repository
		recaptcha.Recaptcha
		notification.TelegramBot
	}{
		datasource,
		googleRecaptcha,
		telegramBot,
	})

	calendlyDiscoverySessionOnScheduledUsecase := calendly_discovery_session_on_scheduled.NewUsecase(struct {
		contract.Repository
		notification.TelegramBot
	}{
		datasource,
		telegramBot,
	})

	primaryDriver.AddUsecase(
		discoverySessionRequestUsecase,
		calendlyDiscoverySessionOnScheduledUsecase,
	)

	os.Exit(m.Run())
}
