package discovery_session_request

import (
	"context"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/enum"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
	_ "time/tzdata"
)

type discoverySessionRequestInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &discoverySessionRequestInteractor{
		outport: outputPort,
	}
}

const (
	LocalDate              = "Asia/Tehran"
	SuccessResponseMessage = "Your session successfully sent"
)

func (r *discoverySessionRequestInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	if !req.TestMode {
		err := r.outport.SiteVerify(ctx, req.Secret, req.RecaptchaToken)
		if err != nil {
			return nil, err
		}
	}

	fileObjs, err := r.outport.FindFilesByIDs(ctx, req.Files)
	if err != nil {
		return nil, err
	}

	discoverySessionObj, err := entity.NewDiscoverySession(entity.DiscoverySessionCreateRequest{
		ID:             req.UUID,
		Now:            req.Now,
		Name:           req.Name,
		Email:          req.Email,
		Date:           req.Date,
		ProjectDetails: req.ProjectDetails,
		Budget:         req.Budget,
		Files:          fileObjs,
	})
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveDiscoverySession(ctx, discoverySessionObj)
	if err != nil {
		return nil, err
	}

	roles, err := r.outport.FindRolesByCodes(ctx, []string{enum.CEO, enum.CTO})
	if err != nil {
		return nil, err
	}

	chatIds, err := r.outport.FindChatIdsByRoles(ctx, roles)
	if err != nil {
		return nil, err
	}

	msg, err := generateMessage(discoverySessionObj, fileObjs, req.APIUrl)
	if err != nil {
		return nil, err
	}

	r.sendNotificationByTelegram(ctx, chatIds, msg)

	res.ID = discoverySessionObj.ID.String()
	res.Message = SuccessResponseMessage

	return res, nil
}

func (r *discoverySessionRequestInteractor) sendNotificationByTelegram(ctx context.Context, chatIds []int64, msg string) {
	for _, chatId := range chatIds {
		_ = r.outport.SendMessage(ctx, chatId, msg, tgbotapi.ModeMarkdown)
	}
}

func generateMessage(discoverySessionObj *entity.DiscoverySession, files []*entity.File, apiUrl string) (string, error) {

	msgFiles := ""

	for _, file := range files {
		msgFiles += fmt.Sprintf("\n[%s](%s%s)", file.Name, apiUrl, file.Path)
	}

	loc, err := time.LoadLocation(LocalDate)
	if err != nil {
		return "", err
	}

	t := discoverySessionObj.Date.Time().In(loc)

	msg := fmt.Sprintf(
		"*New discovery session!*\n\n*%s*\n\n*Name:* %s\n*E-Mail:* %s\n*Date:* %s (%s)\n*Budget:* %s\n*Project Details:*\n%s\n",
		discoverySessionObj.ID,
		discoverySessionObj.Name,
		discoverySessionObj.Email,
		fmt.Sprintf("%d %v %d - %d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()),
		LocalDate,
		discoverySessionObj.Budget,
		discoverySessionObj.ProjectDetails,
	)
	if len(files) > 0 {
		msg += fmt.Sprintf("\n*Files*:%v", msgFiles)
	}

	return msg, nil
}
