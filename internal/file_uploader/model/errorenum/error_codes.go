package errorenum

import "github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/apperror"

const (
	FileNameIsRequired    apperror.ErrorType = "ER0001 file name is required"
	FilePathIsRequired    apperror.ErrorType = "ER0002 file path is required"
	ObjectIdCanNotBeEmpty apperror.ErrorType = "ER0003 object id can not be empty"
	InvalidTypeError      apperror.ErrorType = "ER0004 .%s is an invalid type"
)
