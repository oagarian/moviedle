package movie

import (
	"moviedle/src/core/domain/director"
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/genre"
	"moviedle/src/core/domain/oscar"

	"github.com/google/uuid"
)

type Movie interface {
	ID() *uuid.UUID
	SetID(*uuid.UUID) errors.Error

	Title() string
	SetTitle(string) errors.Error

	CoverURL() string
	SetCoverURL(string) errors.Error

	Genres() []genre.Genre
	SetGenres([]genre.Genre) errors.Error

	Directors() []director.Director
	SetDirectors([]director.Director) errors.Error

	Year() uint16 
	SetYear(uint16) errors.Error

	Slogan() string
	SetSlogan(string) errors.Error

	Resume() string
	SetResume(string) errors.Error

	Oscars() []oscar.Oscar
	SetOscars([]oscar.Oscar) errors.Error
}