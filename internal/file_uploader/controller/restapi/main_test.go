package restapi

import (
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/gateway/postgres"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/usecase/upload_file"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
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

	//err = datasource.ClearFilesTable(context.Background())
	//if err != nil {
	//	log.Fatalf("clear database: %v", err)
	//}

	primaryDriver = NewController(appData, logger.NewSimpleJSONLogger(appData), cfg, jwtToken).(*controller)

	primaryDriver.AddUsecase(
		upload_file.NewUsecase(datasource),
	)

	os.Exit(m.Run())
}
