package restapi

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/util"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/usecase/send_contact_form"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/gin-gonic/gin"
)

type sendContactFormHandlerRequest struct {
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Message        string   `json:"message"`
	RecaptchaToken string   `json:"recaptcha_token"`
	Files          []string `json:"files"`
}

// SendContactForm godoc
// @Summary sending contact form
// @Schemes
// @Description sending contact form data
// @Tags SendContactForm
// @Accept       json
// @Produce      json
// @Param        request body send_contact_form.InportRequest true "body params"
// @Success      200  {object}  send_contact_form.InportResponse
// @Failure      400  {object}  payload.Response
// @Failure      500  {object}  payload.Response
// @Router       /support/v1/contact-us [post]
func (r *controller) sendContactFormHandler() gin.HandlerFunc {

	type InportRequest = send_contact_form.InportRequest
	type InportResponse = send_contact_form.InportResponse

	inport := framework.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type response struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq sendContactFormHandlerRequest
		err := c.BindJSON(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.ID = uuid.New().String()
		req.Now = time.Now()
		req.Name = vo.ContactFormName(jsonReq.Name)
		req.Email = vo.ContactFormEmail(jsonReq.Email)
		req.Message = vo.ContactFormMessage(jsonReq.Message)
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
