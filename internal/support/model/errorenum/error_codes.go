package errorenum

import "github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/apperror"

const (
	ContactFormNameIsRequired    apperror.ErrorType = "ER0001 name is required"
	ContactFormEmailIsRequired   apperror.ErrorType = "ER0002 email is required"
	ContactFormMessageIsRequired apperror.ErrorType = "ER0003 message is required"
	ObjectIdCanNotBeEmpty        apperror.ErrorType = "ER0004 object id can not be empty"
	FileNameIsRequired           apperror.ErrorType = "ER0005 file name is required"
	FilePathIsRequired           apperror.ErrorType = "ER0006 file path is required"
	RoleCodeIsRequired           apperror.ErrorType = "ER0007 role code is required"
	RoleNameIsRequired           apperror.ErrorType = "ER0008 role name is required"
	MaxLenErr                    apperror.ErrorType = "ER0008 The length of %s must be %d characters or fewer. You entered %d characters."
	InvalidEmailAddress          apperror.ErrorType = "ER0009 %s is invalid email address"
)
