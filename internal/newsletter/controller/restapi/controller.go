package restapi

import (
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/token"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"

	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type controller struct {
	framework.ControllerStarter
	framework.UsecaseRegisterer
	Router     *gin.Engine
	log        logger.Logger
	cfg        *configs.Config
	jwtToken   token.JWTToken
	reqCounter prometheus.Counter
	reqLatency prometheus.Histogram
}

func NewController(appData framework.ApplicationData, log logger.Logger, cfg *configs.Config, jwtToken token.JWTToken) framework.ControllerRegisterer {

	router := gin.Default()

	// PING API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, appData)
	})

	// used for web ui development
	// contentStatic, _ := fs.Sub(web.StaticFiles, "dist")
	// router.StaticFS("/web", http.FS(contentStatic))

	// CORS
	router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	address := cfg.Servers[appData.AppName].Address

	return &controller{
		UsecaseRegisterer: framework.NewBaseController(),
		ControllerStarter: NewGracefullyShutdown(log, router, address),
		Router:            router,
		log:               log,
		cfg:               cfg,
		jwtToken:          jwtToken,
	}

}
