package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"strings"
)

type TelegramChatIDID string

func NewTelegramChatIDID(id string) (TelegramChatIDID, error) {

	if strings.TrimSpace(id) == "" {
		return "", errorenum.ObjectIDCanNotBeEmpty
	}

	return TelegramChatIDID(strings.TrimSpace(id)), nil
}

func (r TelegramChatIDID) String() string {
	return string(r)
}
