package send_contact_form

import (
	"context"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/enum"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	SuccessMessageResponse = "Your message successfully sent"
)

type sendContactFormInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &sendContactFormInteractor{
		outport: outputPort,
	}
}

func (r *sendContactFormInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	if !req.TestMode {
		err := r.outport.SiteVerify(ctx, req.Secret, req.RecaptchaToken)
		if err != nil {
			return nil, err
		}
	}

	res := &InportResponse{}

	fileObjs, err := r.outport.FindFilesByIDs(ctx, req.Files)
	if err != nil {
		return nil, err
	}

	contactFormObj, err := entity.NewContactForm(entity.ContactFormCreateRequest{
		ID:      req.ID,
		Now:     req.Now,
		Name:    req.Name,
		Email:   req.Email,
		Message: req.Message,
		Files:   fileObjs,
	})
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveContactForm(ctx, contactFormObj)
	if err != nil {
		return nil, err
	}

	roles, err := r.outport.FindRolesByCodes(ctx, []string{enum.CEO, enum.SUPPORT})
	if err != nil {
		return nil, err
	}

	chatIds, err := r.outport.FindChatIdsByRoles(ctx, roles)
	if err != nil {
		return nil, err
	}

	msg := generateMessage(contactFormObj, fileObjs, req.APIUrl)

	r.sendNotificationByTelegram(ctx, chatIds, msg)

	res.ID = contactFormObj.ID.String()
	res.Message = SuccessMessageResponse

	return res, nil
}

func (r *sendContactFormInteractor) sendNotificationByTelegram(ctx context.Context, chatIds []int64, msg string) {
	for _, chatId := range chatIds {
		_ = r.outport.SendMessage(ctx, chatId, msg, tgbotapi.ModeMarkdown)
	}
}

func generateMessage(contactFormObj *entity.ContactForm, files []*entity.File, apiUrl string) string {

	msgFiles := ""

	for _, file := range files {
		msgFiles += fmt.Sprintf("\n[%s](%s%s)", file.Name, apiUrl, file.Path)
	}

	msg := fmt.Sprintf(
		"*New contact us message!*\n\n*%s*\n\n*Name:* %s\n*E-Mail:* %s\n*Message:*\n%s\n",
		contactFormObj.ID,
		contactFormObj.Name,
		contactFormObj.Email,
		contactFormObj.Message,
	)

	if len(files) > 0 {
		msg += fmt.Sprintf("\n*Files*:%v", msgFiles)
	}

	return msg
}
