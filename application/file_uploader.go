package application

import (
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/controller/restapi"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/gateway/postgres"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/usecase/upload_file"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/token"
)

type fileUploader struct{}

func NewFileUploader() framework.Runner {
	return &fileUploader{}
}

func (fileUploader) Run(cfg *configs.Config) error {

	const appName = "file_uploader"

	appData := framework.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource, err := postgres.NewGateway(log, appData, cfg)
	if err != nil {
		return err
	}

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)
	primaryDriver.RegisterMetrics(appName)

	primaryDriver.AddUsecase(
		upload_file.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
