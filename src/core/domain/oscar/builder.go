package oscar

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/messages"
	"moviedle/src/utils/validator"
	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	oscar 		  *oscar
}

func NewBuilder() *builder {
	return &builder{
		fields: []string{},
        errorMessages: []string{},
        oscar: &oscar{},
	}
}

func (b *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		b.fields = append(b.fields, messages.OscarID)
        b.errorMessages = append(b.errorMessages, messages.OscarInvalidIDErrorMessage)
		return b
	}
	b.oscar.id = &id
	return b
}

func (b *builder) WithName(name string) *builder {
	if validator.IsTextEmpty(name) {
        b.fields = append(b.fields, messages.OscarName)
        b.errorMessages = append(b.errorMessages, messages.OscarNameCannotBeEmptyMessage)
        return b
    }
    b.oscar.name = name
    return b
}

func (b *builder) WithCode(code string) *builder {
	if validator.IsTextEmpty(code) {
        b.fields = append(b.fields, messages.OscarCode)
        b.errorMessages = append(b.errorMessages, messages.OscarCodeCannotBeEmptyMessage)
        return b
    }
    b.oscar.code = code
    return b
}

func (b *builder) Build() (Oscar, errors.Error) {
	if len(b.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(b.errorMessages, map[string]interface{}{
			"fields": b.fields,
		})
	}
	return b.oscar, nil
}