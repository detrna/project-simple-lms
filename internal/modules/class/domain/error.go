package domain

import (
	"main/internal/shared/errors"
)

var (
	ErrClassNotFound      = errors.New(400, "could not find requested class")
	ErrClassSystemIDTaken = errors.New(400, "systemID already taken by other existing class")
)
