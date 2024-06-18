package movie

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"moviedle/src/core/domain/director"
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/genre"
	"moviedle/src/core/domain/oscar"
	"moviedle/src/core/messages"
)

type movieTests struct{}

func TestMovieSetters(t *testing.T) {
	tests := movieTests{}
	t.Run("Test with valid parameters", tests.testWithValidParameters)
	t.Run("Test the ID setters", tests.testIDSetters)
	t.Run("Test the title setters", tests.testTitleSetters)
	t.Run("Test the cover URL setters", tests.testCoverURLSetters)
	t.Run("Test the genres setters", tests.testGenresSetters)
	t.Run("Test the directors setters", tests.testDirectorsSetters)
	t.Run("Test the year setters", tests.testYearSetters)
	t.Run("Test the slogan setters", tests.testSloganSetters)
	t.Run("Test the resume setters", tests.testResumeSetters)
	t.Run("Test the oscars setters", tests.testOscarsSetters)
}

func (*movieTests) testWithValidParameters(t *testing.T) {
	movie := new(movie)
	id := uuid.New()
	_genre, _ := genre.NewBuilder().WithName("Genre").Build()
	_director, _ := director.NewBuilder().WithName("Director").Build()
	_oscar, _ := oscar.NewBuilder().WithName("Best Picture").Build()
	genres := []genre.Genre{_genre}
	directors := []director.Director{_director}
	oscars := []oscar.Oscar{_oscar}

	movie.SetID(id)
	movie.SetTitle("Title")
	movie.SetCoverURL("https://example.com/cover.jpg")
	movie.SetGenres(genres)
	movie.SetDirectors(directors)
	movie.SetYear(2021)
	movie.SetSlogan("Slogan")
	movie.SetResume("Resume")
	movie.SetOscars(oscars)

	assert.NotNil(t, movie)
	assert.Equal(t, &id, movie.ID())
	assert.Equal(t, "Title", movie.Title())
	assert.Equal(t, "https://example.com/cover.jpg", movie.CoverURL())
	assert.Equal(t, genres, movie.Genres())
	assert.Equal(t, directors, movie.Directors())
	assert.Equal(t, uint16(2021), movie.Year())
	assert.Equal(t, "Slogan", movie.Slogan())
	assert.Equal(t, "Resume", movie.Resume())
	assert.Equal(t, oscars, movie.Oscars())
}

func (*movieTests) testIDSetters(t *testing.T) {
	id := uuid.New()
	movie := movie{}
	err := movie.SetID(id)
	assert.Nil(t, err)
	assert.NotNil(t, movie)

	invalidID := uuid.Nil
	err = movie.SetID(invalidID)
	assert.NotNil(t, err)
	assert.Equal(t, errors.NewValidationFromString(messages.MovieInvalidIDErrorMessage), err)
}

func (*movieTests) testTitleSetters(t *testing.T) {
	title := "Test title"
	movie := movie{}
	err := movie.SetTitle(title)
	assert.Nil(t, err)
	assert.NotNil(t, movie)

	invalidTitle := ""
	err = movie.SetTitle(invalidTitle)
	assert.NotNil(t, err)
	assert.Equal(t, errors.NewValidationFromString(messages.MovieTitleCannotBeEmptyMessage), err)
}

func (*movieTests) testCoverURLSetters(t *testing.T) {
	coverURL := "https://example.com/cover.jpg"
	movie := movie{}
	err := movie.SetCoverURL(coverURL)
	assert.Nil(t, err)
	assert.NotNil(t, movie)

	invalidCoverURL := ""
	err = movie.SetCoverURL(invalidCoverURL)
	assert.NotNil(t, err)
	assert.Equal(t, errors.NewValidationFromString(messages.MovieCoverURLCannotBeEmptyMessage), err)
}

func (*movieTests) testGenresSetters(t *testing.T) {
	_genre, _ := genre.NewBuilder().WithName("Genre").Build()
	genres := []genre.Genre{_genre}
	movie := movie{}
	err := movie.SetGenres(genres)
	assert.Nil(t, err)
	assert.NotNil(t, movie)
	assert.Equal(t, genres, movie.Genres())
}

func (*movieTests) testDirectorsSetters(t *testing.T) {
	_director, _ := director.NewBuilder().WithName("Director").Build()
	directors := []director.Director{_director}
	movie := movie{}
	err := movie.SetDirectors(directors)
	assert.Nil(t, err)
	assert.NotNil(t, movie)
	assert.Equal(t, directors, movie.Directors())
}

func (*movieTests) testYearSetters(t *testing.T) {
	year := uint16(2021)
	movie := movie{}
	err := movie.SetYear(year)
	assert.Nil(t, err)
	assert.NotNil(t, movie)
	assert.Equal(t, year, movie.Year())
}

func (*movieTests) testSloganSetters(t *testing.T) {
	slogan := "Test slogan"
	movie := movie{}
	err := movie.SetSlogan(slogan)
	assert.Nil(t, err)
	assert.NotNil(t, movie)
	assert.Equal(t, slogan, movie.Slogan())

	invalidSlogan := ""
	err = movie.SetSlogan(invalidSlogan)
	assert.NotNil(t, err)
	assert.Equal(t, errors.NewValidationFromString(messages.MovieSloganCannotBeEmptyMessage), err)
}

func (*movieTests) testResumeSetters(t *testing.T) {
	resume := "Test resume"
	movie := movie{}
	err := movie.SetResume(resume)
	assert.Nil(t, err)
	assert.NotNil(t, movie)
	assert.Equal(t, resume, movie.Resume())

	invalidResume := ""
	err = movie.SetResume(invalidResume)
	assert.NotNil(t, err)
	assert.Equal(t, errors.NewValidationFromString(messages.MovieResumeCannotBeEmptyMessage), err)
}

func (*movieTests) testOscarsSetters(t *testing.T) {
	_oscar, _ := oscar.NewBuilder().WithName("Best Picture").Build()
    oscars := []oscar.Oscar{_oscar}
    movie := movie{}
    err := movie.SetOscars(oscars)
    assert.Nil(t, err)
    assert.NotNil(t, movie)
    assert.Equal(t, oscars, movie.Oscars())
}
