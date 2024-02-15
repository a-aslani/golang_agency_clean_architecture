package entity

import (
	"time"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/vo"
)

type Subscriber struct {
	ID      vo.SubscriberID    `bson:"_id" json:"id"`
	Email   vo.SubscriberEmail `bson:"email" json:"email"`
	Created time.Time          `bson:"created" json:"created"`
}

type SubscriberCreateRequest struct {
	ID    string             `json:"-"`
	Now   time.Time          `json:"-"`
	Email vo.SubscriberEmail `json:"email"`
}

func (r SubscriberCreateRequest) Validate() error {

	if err := r.Email.Validate(); err != nil {
		return err
	}

	return nil
}

func NewSubscriber(req SubscriberCreateRequest) (*Subscriber, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewSubscriberID(req.ID)
	if err != nil {
		return nil, err
	}

	var obj Subscriber
	obj.ID = id
	obj.Created = req.Now
	obj.Email = req.Email

	return &obj, nil
}
