package movie

import (
	"moviedle/src/core/domain/director"
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/genre"
	"moviedle/src/core/domain/oscar"
	"moviedle/src/core/messages"
	"moviedle/src/utils/validator"
	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	movie         *movie
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		movie:         &movie{},
	}
}

func (b *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		b.fields = append(b.fields, messages.MovieID)
		b.errorMessages = append(b.errorMessages, messages.MovieInvalidIDErrorMessage)
		return b
	}
	b.movie.id = &id
	return b
}

func (b *builder) WithTitle(title string) *builder {
	if validator.IsTextEmpty(title) {
		b.fields = append(b.fields, messages.MovieTitle)
		b.errorMessages = append(b.errorMessages, messages.MovieTitleCannotBeEmptyMessage)
		return b
	}
	b.movie.title = title
	return b
}

func (b *builder) WithCoverURL(coverURL string) *builder {
	if validator.IsTextEmpty(coverURL) {
		b.fields = append(b.fields, messages.MovieCoverURL)
		b.errorMessages = append(b.errorMessages, messages.MovieCoverURLCannotBeEmptyMessage)
		return b
	}
	b.movie.coverURL = coverURL
	return b
}

func (b *builder) WithGenres(genres []genre.Genre) *builder {
	b.movie.genres = genres
	return b
}

func (b *builder) WithDirectors(directors []director.Director) *builder {
	b.movie.directors = directors
	return b
}

func (b *builder) WithYear(year uint16) *builder {
	b.movie.year = year
	return b
}

func (b *builder) WithSlogan(slogan string) *builder {
	if validator.IsTextEmpty(slogan) {
		b.fields = append(b.fields, messages.MovieSlogan)
		b.errorMessages = append(b.errorMessages, messages.MovieSloganCannotBeEmptyMessage)
		return b
	}
	b.movie.slogan = slogan
	return b
}

func (b *builder) WithResume(resume string) *builder {
	if validator.IsTextEmpty(resume) {
		b.fields = append(b.fields, messages.MovieResume)
		b.errorMessages = append(b.errorMessages, messages.MovieResumeCannotBeEmptyMessage)
		return b
	}
	b.movie.resume = resume
	return b
}

func (b *builder) WithOscars(oscars []oscar.Oscar) *builder {
	b.movie.oscars = oscars
	return b
}

func (b *builder) Build() (*movie, errors.Error) {
	if len(b.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(b.errorMessages, map[string]interface{}{
			"fields": b.fields,
		})
	}
	return b.movie, nil
}
