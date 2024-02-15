package application

import (
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	_ "github.com/a-aslani/golang_agency_clean_architecture/docs"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type swagger struct{}

func NewSwagger() framework.Runner {
	return &swagger{}
}

// @title           Swagger Example API
// @version         1.0
// @description     API Documentation.
// @termsOfService  https://github.com/a-aslani

// @contact.name   API Support
// @contact.url    https://github.com/a-aslani
// @contact.email  a.aslani.dev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8888

// Run
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func (swagger) Run(cfg *configs.Config) error {
	r := gin.Default()
	r.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("swagger service running at :%d", cfg.SwaggerPort)
	err := r.Run(fmt.Sprintf(":%d", cfg.SwaggerPort))
	if err != nil {
		return err
	}
	return nil
}
