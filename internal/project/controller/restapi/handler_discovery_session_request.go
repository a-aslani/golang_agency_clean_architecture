package restapi

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/util"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/usecase/discovery_session_request"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/gin-gonic/gin"
)

type discoverySessionRequestHandlerRequest struct {
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Date           string   `json:"date"`
	ProjectDetails string   `json:"project_details"`
	RecaptchaToken string   `json:"recaptcha_token"`
	Files          []string `json:"files"`
}

// DiscoverySessionRequest godoc
// @Summary send discovery session
// @Schemes
// @Description sending discovery session date
// @Tags DiscoverySessionRequest
// @Accept       json
// @Produce      json
// @Param        request body discoverySessionRequestHandlerRequest true "body params"
// @Success      200  {object}  discovery_session_request.InportResponse
// @Failure      400  {object}  payload.Response
// @Failure      500  {object}  payload.Response
// @Router       /project/v1/discovery-session [post]
func (r *controller) discoverySessionRequestHandler() gin.HandlerFunc {

	type InportRequest = discovery_session_request.InportRequest
	type InportResponse = discovery_session_request.InportResponse

	inport := framework.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type response struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq discoverySessionRequestHandlerRequest
		err := c.BindJSON(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		layout := "2006-01-02T15:04:05.000Z"
		date, err := time.Parse(layout, jsonReq.Date)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.UUID = uuid.New().String()
		req.Now = time.Now()
		req.Name = vo.DiscoverySessionName(jsonReq.Name)
		req.Email = vo.DiscoverySessionEmail(jsonReq.Email)
		req.Date = vo.DiscoverySessionDate(date)
		req.ProjectDetails = vo.DiscoverySessionProjectDetails(jsonReq.ProjectDetails)
		req.RecaptchaToken = jsonReq.RecaptchaToken
		req.Secret = r.cfg.RecaptchaSecretKey
		req.Files = jsonReq.Files
		req.APIUrl = r.cfg.APIUrl
		req.TestMode = r.cfg.TestMode

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.ID = res.ID
		jsonRes.Message = res.Message

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
