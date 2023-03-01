package common

import "errors"

var (
	// repository errors
	ErrRecordNotFound  = errors.New("record not found")
	ErrDuplicateEntry  = errors.New("duplicate entry")
	ErrQueryRepository = errors.New("unable to query repository")

	// request parameter errors
	ErrInvalidServiceIdValue = errors.New("invalid service_id value")
	ErrInvalidLocaleValue    = errors.New("invalid locale value")
	ErrInvalidKeyValue       = errors.New("invalid key value")
	ErrInvalidValue          = errors.New("invalid value")

	ErrBadRouting       = errors.New("bad routing")
	ErrRequestParamBody = errors.New("request param and body values should be same")
)

// Error codes
var (
	DB_UNIQUE_CONSTRAINT_VIOLATION = "23505"
)
