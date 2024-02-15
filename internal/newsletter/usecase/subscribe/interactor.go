package subscribe

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/errorenum"
)

const (
	SuccessResponseMessage = "Your email has been successfully registered in the newsletter"
)

type subscribeInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &subscribeInteractor{
		outport: outputPort,
	}
}

func (r *subscribeInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	subscriberObj, err := r.outport.FindOneSubscriberByEmail(ctx, req.Email.String())
	if subscriberObj != nil {
		return nil, errorenum.ThisEmailAddressIsAlreadyUsed.Var(req.Email.String())
	}

	subscriberObj, err = entity.NewSubscriber(req.SubscriberCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveSubscriber(ctx, subscriberObj)
	if err != nil {
		return nil, err
	}

	res.ID = subscriberObj.ID
	res.Message = SuccessResponseMessage

	return res, nil
}
