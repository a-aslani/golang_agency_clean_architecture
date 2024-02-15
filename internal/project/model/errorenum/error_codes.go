package errorenum

import "github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/apperror"

const (
	EmailIsRequired             apperror.ErrorType = "ER0001 email is required"
	InvalidEmailAddress         apperror.ErrorType = "ER0002 %s is invalid email address"
	ObjectIDCanNotBeEmpty       apperror.ErrorType = "ER0003 object id can not be empty"
	NameIsRequired              apperror.ErrorType = "ER0004 name is required"
	ProjectDetailsIsRequired    apperror.ErrorType = "ER0005 project details is required"
	InvalidDiscoverySessionDate apperror.ErrorType = "ER0006 The selected date for the meeting is not valid"
	MaxLenErr                   apperror.ErrorType = "ER0007 The length of %s must be %d characters or fewer. You entered %d characters."
	FileNameIsRequired          apperror.ErrorType = "ER0008 file name is required"
	FilePathIsRequired          apperror.ErrorType = "ER0009 file path is required"
	RoleCodeIsRequired          apperror.ErrorType = "ER0010 role code is required"
	RoleNameIsRequired          apperror.ErrorType = "ER0011 role name is required"
)
