package genre

import (
	"moviedle/src/core/domain/errors"
	"github.com/google/uuid"
)

type Genre interface {
	ID() *uuid.UUID
	SetID(uuid.UUID) errors.Error

	Name() string
	SetName(string) errors.Error

	Code() string
	SetCode(string) errors.Error
}