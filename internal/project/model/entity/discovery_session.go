package entity

import (
	"time"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/vo"
)

type DiscoverySession struct {
	ID             vo.DiscoverySessionID             `bson:"_id" json:"id"`
	Name           vo.DiscoverySessionName           `bson:"name" json:"name"`
	Email          vo.DiscoverySessionEmail          `bson:"email" json:"email"`
	Date           vo.DiscoverySessionDate           `bson:"date" json:"date"`
	ProjectDetails vo.DiscoverySessionProjectDetails `bson:"projectDetails" json:"projectDetails"`
	Files          []*File                           `json:"files"`
	Created        time.Time                         `bson:"created" json:"created"`
	Updated        time.Time                         `bson:"updated" json:"updated"`
}

type DiscoverySessionCreateRequest struct {
	ID             string                            `json:"-"`
	Now            time.Time                         `json:"-"`
	Name           vo.DiscoverySessionName           `json:"name"`
	Email          vo.DiscoverySessionEmail          `json:"email"`
	Date           vo.DiscoverySessionDate           `json:"date"`
	ProjectDetails vo.DiscoverySessionProjectDetails `json:"projectDetails"`
	Files          []*File                           `json:"files"`
}

func (r DiscoverySessionCreateRequest) Validate() error {

	if err := r.Name.Validate(); err != nil {
		return err
	}

	if err := r.Email.Validate(); err != nil {
		return err
	}

	if err := r.Date.Validate(); err != nil {
		return err
	}

	if err := r.ProjectDetails.Validate(); err != nil {
		return err
	}

	return nil
}

func NewDiscoverySession(req DiscoverySessionCreateRequest) (*DiscoverySession, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewDiscoverySessionID(req.ID)
	if err != nil {
		return nil, err
	}

	var obj DiscoverySession
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now
	obj.Name = req.Name
	obj.Email = req.Email
	obj.Date = req.Date
	obj.ProjectDetails = req.ProjectDetails
	obj.Files = req.Files

	return &obj, nil
}
