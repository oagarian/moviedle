package director

import (
	"moviedle/src/core/domain/errors"
	"github.com/google/uuid"
)

type Director interface {
	ID() *uuid.UUID
	SetID(uuid.UUID) errors.Error

	Name() string
	SetName(string) errors.Error
}