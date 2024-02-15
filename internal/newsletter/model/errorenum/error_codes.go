package errorenum

import "github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/apperror"

const (
	ObjectIDCanNotBeEmpty         apperror.ErrorType = "ER0001 object id can not be empty"
	EmailIsRequired               apperror.ErrorType = "ER0002 email is required"
	InvalidEmailAddress           apperror.ErrorType = "ER0003 '%s' is a invalid email address"
	ThisEmailAddressIsAlreadyUsed apperror.ErrorType = "ER0004 '%s' this email address is already used"
)
