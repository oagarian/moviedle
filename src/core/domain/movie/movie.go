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

type movie struct {
	id 	        *uuid.UUID
	title       string
	coverURL    string
	genres      []genre.Genre
	directors   []director.Director
	year        uint16
	slogan      string
	resume      string
	oscars      []oscar.Oscar
}

func (m *movie) ID() *uuid.UUID {
    return m.id
}

func (m *movie) SetID(id *uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(*id) {
        return errors.NewValidationFromString(messages.MovieInvalidIDErrorMessage)
    }
    m.id = id
    return nil
}

func (m *movie) Title() string {
    return m.title
}

func (m *movie) SetTitle(title string) errors.Error {
	if validator.IsTextEmpty(title) {
        return errors.NewValidationFromString(messages.MovieTitleCannotBeEmptyMessage)
    }
    m.title = title
    return nil
}


func (m *movie) CoverURL() string {
    return m.coverURL
}

func (m *movie) SetCoverURL(coverURL string) errors.Error {
	if validator.IsTextEmpty(coverURL) {
        return errors.NewValidationFromString(messages.MovieCoverURLCannotBeEmptyMessage)
    }
	m.coverURL = coverURL
    return nil
}


func (m *movie) Genres() []genre.Genre {
    return m.genres
}

func (m *movie) SetGenres(genres []genre.Genre) errors.Error {
    m.genres = genres
    return nil
}

func (m *movie) Directors() []director.Director {
    return m.directors
}

func (m *movie) SetDirectors(directors []director.Director) errors.Error {
    m.directors = directors
    return nil
}

func (m *movie) Year() uint16 {
    return m.year
}

func (m *movie) SetYear(year uint16) errors.Error {
    m.year = year
    return nil
}

func (m *movie) Slogan() string {
    return m.slogan
}

func (m *movie) SetSlogan(slogan string) errors.Error {
	if validator.IsTextEmpty(slogan) {
        return errors.NewValidationFromString(messages.MovieSloganCannotBeEmptyMessage)
    }
    m.slogan = slogan
    return nil
}

func (m *movie) Resume() string {
    return m.resume
}

func (m *movie) SetResume(resume string) errors.Error {
	if validator.IsTextEmpty(resume) {
        return errors.NewValidationFromString(messages.MovieResumeCannotBeEmptyMessage)
    }
    m.resume = resume
    return nil
}

func (m *movie) Oscars() []oscar.Oscar {
    return m.oscars
}

func (m *movie) SetOscars(oscars []oscar.Oscar) errors.Error {
    m.oscars = oscars
    return nil
}