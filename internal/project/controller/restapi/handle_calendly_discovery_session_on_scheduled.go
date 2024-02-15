package restapi

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/usecase/calendly_discovery_session_on_scheduled"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/util"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CalendlyDiscoverySessionOnScheduled godoc
// @Summary notify scheduled Calendly
// @Schemes
// @Description notify messaging when scheduled session with the Calendly
// @Tags CalendlyDiscoverySessionOnScheduled
// @Accept       json
// @Produce      json
// @Success      200  {object}  calendly_discovery_session_on_scheduled.InportResponse
// @Failure      400  {object}  payload.Response
// @Failure      500  {object}  payload.Response
// @Router       /project/v1/discovery-session/calendly [get]
func (r *controller) calendlyDiscoverySessionOnScheduledHandler() gin.HandlerFunc {

	type InportRequest = calendly_discovery_session_on_scheduled.InportRequest
	type InportResponse = calendly_discovery_session_on_scheduled.InportResponse

	inport := framework.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var req InportRequest

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))
	}
}
