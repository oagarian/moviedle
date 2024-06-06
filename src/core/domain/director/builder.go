package director

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/messages"
	"moviedle/src/utils/validator"
	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	director 	  *director
}

func NewBuilder() *builder {
	return &builder{
		fields: []string{},
        errorMessages: []string{},
        director: &director{},
	}
}

func (b *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		b.fields = append(b.fields, messages.DirectorID)
        b.errorMessages = append(b.errorMessages, messages.DirectorIDInvalidIDMessage)
		return b
	}
	b.director.id = &id
	return b
}

func (b *builder) WithName(name string) *builder {
	if validator.IsTextEmpty(name) {
        b.fields = append(b.fields, messages.DirectorName)
        b.errorMessages = append(b.errorMessages, messages.DirectorNameCannotBeEmptyMessage)
        return b
    }
    b.director.name = name
    return b
}

func (b *builder) Build() (Director, errors.Error) {
	if len(b.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(b.errorMessages, map[string]interface{}{
			"fields": b.fields,
		})
	}
	return b.director, nil
}