package genre

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/messages"
	"moviedle/src/utils/validator"
	"github.com/google/uuid"
)

type genre struct {
	id *uuid.UUID
	name string
	code string
}

func (g *genre) ID() *uuid.UUID {
    return g.id
}

func (g *genre) SetID(id uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(id) {
		return errors.NewValidationFromString(messages.GenreInvalidIDErrorMessage)
	}
    g.id = &id
    return nil
}

func (g *genre) Name() string {
    return g.name
}

func (g *genre) SetName(name string) errors.Error {
	if validator.IsTextEmpty(name) {
		return errors.NewValidationFromString(messages.GenreNameCannotBeEmptyMessage)
	}
    g.name = name
    return nil
}

func (g *genre) Code() string {
    return g.code
}

func (g *genre) SetCode(code string) errors.Error {
	if validator.IsTextEmpty(code) {
		return errors.NewValidationFromString(messages.GenreCodeCannotBeEmptyMessage)
	}
    g.code = code
    return nil
}