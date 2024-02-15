package application

import (
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/controller/restapi"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/gateway/postgres"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/usecase/subscribe"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/token"
)

type newsletter struct{}

func NewNewsletter() framework.Runner {
	return &newsletter{}
}

func (newsletter) Run(cfg *configs.Config) error {

	const appName = "newsletter"

	fmt.Println(cfg.Servers[appName].Address)

	appData := framework.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	postgresRepo, err := postgres.NewGateway(log, appData, cfg)
	if err != nil {
		return err
	}

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)
	primaryDriver.RegisterMetrics(appName)

	primaryDriver.AddUsecase(
		subscribe.NewUsecase(struct {
			contract.Repository
		}{
			postgresRepo,
		}),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
