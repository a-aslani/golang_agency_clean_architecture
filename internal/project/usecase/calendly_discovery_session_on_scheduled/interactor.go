package calendly_discovery_session_on_scheduled

import (
	"context"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/enum"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
	_ "time/tzdata"
)

const LocalDate = "Asia/Tehran"

type calendlyDiscoverySessionOnScheduledInteractor struct {
	outport Outport
}

func NewUsecase(outport Outport) Inport {
	return &calendlyDiscoverySessionOnScheduledInteractor{
		outport: outport,
	}
}

func (c *calendlyDiscoverySessionOnScheduledInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	roles, err := c.outport.FindRolesByCodes(ctx, []string{enum.CEO, enum.CTO})
	if err != nil {
		return nil, err
	}

	chatIds, err := c.outport.FindChatIdsByRoles(ctx, roles)
	if err != nil {
		return nil, err
	}

	msg, err := generateMessage()
	if err != nil {
		return nil, err
	}

	c.sendNotificationByTelegram(ctx, chatIds, msg)

	return res, nil
}

func (c *calendlyDiscoverySessionOnScheduledInteractor) sendNotificationByTelegram(ctx context.Context, chatIds []int64, msg string) {
	for _, chatId := range chatIds {
		_ = c.outport.SendMessage(ctx, chatId, msg, tgbotapi.ModeMarkdown)
	}
}

func generateMessage() (string, error) {

	loc, err := time.LoadLocation(LocalDate)
	if err != nil {
		return "", err
	}

	t := time.Now().In(loc)

	msg := fmt.Sprintf(
		"*New discovery session with Calendly!*\n\n%s (%s)\n",
		fmt.Sprintf("%d %v %d - %d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()),
		LocalDate,
	)
	return msg, nil
}
