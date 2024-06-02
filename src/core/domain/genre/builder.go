package genre

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/messages"
	validator "moviedle/src/utils"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	genre 		  *genre
}

func NewBuilder() *builder {
	return &builder{
		fields: []string{},
        errorMessages: []string{},
        genre: &genre{},
	}
}

func (b *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		b.fields = append(b.fields, messages.GenreID)
        b.errorMessages = append(b.errorMessages, messages.GenreInvalidIDErrorMessage)
		return b
	}
	b.genre.id = &id
	return b
}

func (b *builder) WithName(name string) *builder {
	if validator.IsTextEmpty(name) {
        b.fields = append(b.fields, messages.GenreName)
        b.errorMessages = append(b.errorMessages, messages.GenreNameCannotBeEmptyMessage)
        return b
    }
    b.genre.name = name
    return b
}

func (b *builder) WithCode(code string) *builder {
	if validator.IsTextEmpty(code) {
        b.fields = append(b.fields, messages.GenreCode)
        b.errorMessages = append(b.errorMessages, messages.GenreCodeCannotBeEmptyMessage)
        return b
    }
    b.genre.code = code
    return b
}

func (b *builder) Build() (Genre, errors.Error) {
	if len(b.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(b.errorMessages, map[string]interface{}{
			"fields": b.fields,
		})
	}
	return b.genre, nil
}