package entity

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/vo"
)

type Role struct {
	ID   vo.RoleID   `bson:"_id" json:"id"`
	Code vo.RoleCode `bson:"code" json:"code"`
	Name vo.RoleName `bson:"name" json:"name"`
}

type RoleCreateRequest struct {
	ID   string      `json:"-"`
	Code vo.RoleCode `json:"code"`
	Name vo.RoleName `json:"name"`
}

func (r RoleCreateRequest) Validate() error {

	if err := r.Code.Validate(); err != nil {
		return err
	}

	if err := r.Name.Validate(); err != nil {
		return err
	}

	return nil
}

func NewRole(req RoleCreateRequest) (*Role, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewRoleID(req.ID)
	if err != nil {
		return nil, err
	}

	var obj Role
	obj.ID = id
	obj.Code = req.Code
	obj.Name = req.Name

	return &obj, nil
}
