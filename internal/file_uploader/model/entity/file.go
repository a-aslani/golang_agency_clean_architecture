package entity

import (
	"time"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/vo"
)

type File struct {
	ID      vo.FileID   `bson:"_id" json:"id"`
	Name    vo.FileName `bson:"name" json:"name"`
	Path    vo.FilePath `bson:"path" json:"path"`
	Created time.Time   `bson:"created" json:"created"`
}

type FileCreateRequest struct {
	ID   string      `json:"-"`
	Name vo.FileName `json:"name"`
	Path vo.FilePath `json:"path"`
	Now  time.Time   `json:"-"`
}

func (r FileCreateRequest) Validate() error {

	if err := r.Name.Validate(); err != nil {
		return err
	}

	if err := r.Path.Validate(); err != nil {
		return err
	}

	return nil
}

func NewFile(req FileCreateRequest) (*File, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewFileID(req.ID)
	if err != nil {
		return nil, err
	}

	var obj File
	obj.ID = id
	obj.Created = req.Now
	obj.Name = req.Name
	obj.Path = req.Path

	return &obj, nil
}
