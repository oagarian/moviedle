package oscar

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/messages"
	"moviedle/src/utils/validator"
	"github.com/google/uuid"
)

type oscar struct {
	id *uuid.UUID
	name string
	code string
}

func (g *oscar) ID() *uuid.UUID {
    return g.id
}

func (g *oscar) SetID(id uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(id) {
		return errors.NewValidationFromString(messages.OscarInvalidIDErrorMessage)
	}
    g.id = &id
    return nil
}

func (g *oscar) Name() string {
    return g.name
}

func (g *oscar) SetName(name string) errors.Error {
	if validator.IsTextEmpty(name) {
		return errors.NewValidationFromString(messages.OscarNameCannotBeEmptyMessage)
	}
    g.name = name
    return nil
}

func (g *oscar) Code() string {
    return g.code
}

func (g *oscar) SetCode(code string) errors.Error {
	if validator.IsTextEmpty(code) {
		return errors.NewValidationFromString(messages.OscarCodeCannotBeEmptyMessage)
	}
    g.code = code
    return nil
}