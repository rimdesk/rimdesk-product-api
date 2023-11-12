package exceptions

import "errors"

var (
	ErrCompanyHeaderMustBePresent = errors.New("X-Company-Id is required in the header")
)
