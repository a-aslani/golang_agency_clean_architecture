package entity

import (
	"time"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/vo"
)

type ContactForm struct {
	ID        vo.ContactFormID      `bson:"_id" json:"id"`
	Name      vo.ContactFormName    `json:"name"`
	Email     vo.ContactFormEmail   `json:"email"`
	Message   vo.ContactFormMessage `json:"message"`
	Files     []*File               `json:"files"`
	CreatedAt time.Time             `bson:"createdAt" json:"createdAt"`
}

type ContactFormCreateRequest struct {
	ID      string                `json:"-"`
	Now     time.Time             `json:"-"`
	Name    vo.ContactFormName    `json:"name"`
	Email   vo.ContactFormEmail   `json:"email"`
	Message vo.ContactFormMessage `json:"message"`
	Files   []*File               `json:"files"`
}

func (r ContactFormCreateRequest) Validate() error {

	if err := r.Name.Validate(); err != nil {
		return err
	}

	if err := r.Email.Validate(); err != nil {
		return err
	}

	if err := r.Message.Validate(); err != nil {
		return err
	}

	for _, file := range r.Files {

		if err := file.Path.Validate(); err != nil {
			return err
		}

		if err := file.Name.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func NewContactForm(req ContactFormCreateRequest) (*ContactForm, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewContactFormID(req.ID)
	if err != nil {
		return nil, err
	}

	var obj ContactForm
	obj.ID = id
	obj.CreatedAt = req.Now
	obj.Name = req.Name
	obj.Email = req.Email
	obj.Message = req.Message
	obj.Files = req.Files

	return &obj, nil
}
