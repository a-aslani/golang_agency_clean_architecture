package entity

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/vo"
)

type TelegramChatID struct {
	ID     vo.TelegramChatIDID `bson:"_id" json:"id"`
	ChatID string              `bson:"chat_id" json:"chat_id"`
	RoleID string              `bson:"role_id" json:"role_id"`
}

type TelegramChatIDCreateRequest struct {
	ID     string `json:"-"`
	ChatID string `json:"chat_id"`
	RoleID string `json:"role_id"`
}

func (r TelegramChatIDCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewChatID(req TelegramChatIDCreateRequest) (*TelegramChatID, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewTelegramChatIDID(req.ID)
	if err != nil {
		return nil, err
	}

	var obj TelegramChatID
	obj.ID = id
	obj.ChatID = req.ChatID
	obj.RoleID = req.RoleID

	return &obj, nil
}
