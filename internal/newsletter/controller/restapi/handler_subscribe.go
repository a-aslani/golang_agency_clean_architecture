package restapi

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/usecase/subscribe"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/util"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type subscribeHandlerRequest struct {
	Email string `json:"email"`
}

// Subscribe godoc
// @Summary send subscribe to the newsletter
// @Schemes
// @Description send subscribe to the newsletter
// @Tags Subscribe
// @Accept       json
// @Produce      json
// @Param        request body subscribeHandlerRequest true "body params"
// @Success      200  {object}  subscribe.InportResponse
// @Failure      400  {object}  payload.Response
// @Failure      500  {object}  payload.Response
// @Router       /newsletter/v1/subscribers [post]
func (r *controller) subscribeHandler() gin.HandlerFunc {

	type InportRequest = subscribe.InportRequest
	type InportResponse = subscribe.InportResponse

	inport := framework.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		*InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.BindJSON(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest

		req.Now = time.Now()
		req.ID = uuid.New().String()
		req.Email = jsonReq.Email

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.InportResponse = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
