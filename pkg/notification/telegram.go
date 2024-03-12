package notification

import (
	"context"
	"crypto/tls"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"time"
)

//go:generate go run go.uber.org/mock/mockgen -destination mocks/telegram_mock.go -package mocktelegram github.com/a-aslani/golang_agency_clean_architecture/pkg/notification TelegramBot
type TelegramBot interface {
	SendMessage(ctx context.Context, chatId int64, text string, parseMode string) error
	CommandHandling(ctx context.Context, cmd func(tgbotapi.Update) string) error
}

type telegramBot struct {
	bot      *tgbotapi.BotAPI
	isEnable bool
}

func NewTelegramBot(telegramBotCfg configs.TelegramBot, debug bool) (*telegramBot, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}

	var bot *tgbotapi.BotAPI
	var err error

	if telegramBotCfg.Enable {
		bot, err = tgbotapi.NewBotAPIWithClient(telegramBotCfg.Token, tgbotapi.APIEndpoint, client)
		if err != nil {
			return nil, err
		}

		bot.Debug = debug
	}

	return &telegramBot{
		bot:      bot,
		isEnable: telegramBotCfg.Enable,
	}, nil
}

func (t *telegramBot) CommandHandling(ctx context.Context, cmd func(tgbotapi.Update) string) error {

	if !t.isEnable {
		return nil
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = cmd(update)

		if _, err := t.bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (t *telegramBot) SendMessage(ctx context.Context, chatId int64, text string, parseMode string) error {

	if !t.isEnable {
		return nil
	}

	msg := tgbotapi.NewMessage(chatId, text)

	msg.ParseMode = parseMode

	_, err := t.bot.Send(msg)

	return err
}
