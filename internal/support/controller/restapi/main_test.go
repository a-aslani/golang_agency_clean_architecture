package restapi

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/gateway/postgres"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/usecase/send_contact_form"
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

	const appName = "support"

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

	err = datasource.ClearContactUsFileTable(context.Background())
	if err != nil {
		log.Fatalf("clear database: %v", err)
	}

	err = datasource.ClearContactUsTable(context.Background())
	if err != nil {
		log.Fatalf("clear database: %v", err)
	}

	err = datasource.ClearTelegramChatIdsTable(context.Background())
	if err != nil {
		log.Fatalf("clear database: %v", err)
	}

	googleRecaptcha := recaptcha.NewGoogleRecaptcha()

	telegramBot, err := notification.NewTelegramBot(cfg.TelegramBotToken, true)
	if err != nil {
		log.Fatalf("telegram bot error: %v", err)
	}

	primaryDriver = NewController(appData, logger.NewSimpleJSONLogger(appData), cfg, jwtToken).(*controller)

	primaryDriver.AddUsecase(
		send_contact_form.NewUsecase(struct {
			contract.Repository
			recaptcha.Recaptcha
			notification.TelegramBot
		}{
			datasource,
			googleRecaptcha,
			telegramBot,
		}),
	)

	os.Exit(m.Run())
}
