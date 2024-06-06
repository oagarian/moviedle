package director

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/messages"
	"moviedle/src/utils/validator"

	"github.com/google/uuid"
)

type director struct {
	id 		*uuid.UUID
	name 	string
}

func (d *director) ID() *uuid.UUID {
    return d.id
}

func (d *director) SetID(id uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(id) {
		return errors.NewValidationFromString(messages.DirectorIDInvalidIDMessage)
	}
    d.id = &id
    return nil
}

func (d *director) Name() string {
    return d.name
}

func (d *director) SetName(name string) errors.Error {
	if validator.IsTextEmpty(name) {
        return errors.NewValidationFromString(messages.DirectorNameCannotBeEmptyMessage)
    }
    d.name = name
    return nil
}