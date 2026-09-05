package domain

import (
	"main/internal/shared/errors"
)

var (
	ErrClassNotFound      = errors.New(404, "could not find requested class")
	ErrClassSystemIDTaken = errors.New(409, "systemID already taken by other existing class")
)
