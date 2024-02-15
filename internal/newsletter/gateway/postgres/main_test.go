package postgres

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"log"
	"os"
	"testing"
)

var datasource *gateway

func TestMain(m *testing.M) {

	const appName = "newsletter"

	appData := framework.NewApplicationData(appName)

	var err error

	cfg, err := configs.InitConfig("../../../../config.test.yml")
	if err != nil {
		log.Fatalf("reading config file error: %v", err)
	}

	datasource, err = NewGateway(logger.NewSimpleJSONLogger(appData), appData, cfg)
	if err != nil {
		log.Fatalf("connect to the database: %v", err)
	}

	err = datasource.ClearSubscribers(context.Background())
	if err != nil {
		log.Fatalf("clear database: %v", err)
	}

	os.Exit(m.Run())
}
