package movie

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"moviedle/src/core/domain/director"
	"moviedle/src/core/domain/genre"
	"moviedle/src/core/domain/oscar"
)

type movieBuilderTests struct{}

func TestMovieBuilder(t *testing.T) {
	testcases := movieBuilderTests{}
	t.Run("Test builder with valid params", testcases.testWithValidParams)
	t.Run("Test builder with invalid ID", testcases.testWithInvalidID)
	t.Run("Test builder with invalid title", testcases.testWithInvalidTitle)
	t.Run("Test builder with invalid cover URL", testcases.testWithInvalidCoverURL)
	t.Run("Test builder with invalid slogan", testcases.testWithInvalidSlogan)
	t.Run("Test builder with invalid resume", testcases.testWithInvalidResume)
}

func (bt *movieBuilderTests) testWithValidParams(t *testing.T) {
	_genre, _ := genre.NewBuilder().WithName("Genre").Build()
	_director, _ := director.NewBuilder().WithName("Director").Build()
	_oscar, _ := oscar.NewBuilder().WithName("Best Picture").Build()

	genres := []genre.Genre{_genre}
	directors := []director.Director{_director}
	oscars := []oscar.Oscar{_oscar}

	builder := NewBuilder()
	builder.WithID(uuid.New())
	builder.WithTitle("Title")
	builder.WithCoverURL("https://example.com/cover.jpg")
	builder.WithGenres(genres)
	builder.WithDirectors(directors)
	builder.WithYear(2021)
	builder.WithSlogan("Slogan")
	builder.WithResume("Resume")
	builder.WithOscars(oscars)

	movie, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, movie)
}

func (bt *movieBuilderTests) testWithInvalidID(t *testing.T) {
	_genre, _ := genre.NewBuilder().WithName("Genre").Build()
	_director, _ := director.NewBuilder().WithName("Director").Build()
	_oscar, _ := oscar.NewBuilder().WithName("Best Picture").Build()

	genres := []genre.Genre{_genre}
	directors := []director.Director{_director}
	oscars := []oscar.Oscar{_oscar}

	builder := NewBuilder()
	builder.WithID(uuid.Nil)
	builder.WithTitle("Title")
	builder.WithCoverURL("https://example.com/cover.jpg")
	builder.WithGenres(genres)
	builder.WithDirectors(directors)
	builder.WithYear(2021)
	builder.WithSlogan("Slogan")
	builder.WithResume("Resume")
	builder.WithOscars(oscars)

	_, err := builder.Build()
	assert.NotNil(t, err)
}

func (bt *movieBuilderTests) testWithInvalidTitle(t *testing.T) {
	_genre, _ := genre.NewBuilder().WithName("Genre").Build()
	_director, _ := director.NewBuilder().WithName("Director").Build()
	_oscar, _ := oscar.NewBuilder().WithName("Best Picture").Build()

	genres := []genre.Genre{_genre}
	directors := []director.Director{_director}
	oscars := []oscar.Oscar{_oscar}

	builder := NewBuilder()
	builder.WithID(uuid.New())
	builder.WithTitle("")
	builder.WithCoverURL("https://example.com/cover.jpg")
	builder.WithGenres(genres)
	builder.WithDirectors(directors)
	builder.WithYear(2021)
	builder.WithSlogan("Slogan")
	builder.WithResume("Resume")
	builder.WithOscars(oscars)

	_, err := builder.Build()
	assert.NotNil(t, err)
}

func (bt *movieBuilderTests) testWithInvalidCoverURL(t *testing.T) {
	_genre, _ := genre.NewBuilder().WithName("Genre").Build()
	_director, _ := director.NewBuilder().WithName("Director").Build()
	_oscar, _ := oscar.NewBuilder().WithName("Best Picture").Build()

	genres := []genre.Genre{_genre}
	directors := []director.Director{_director}
	oscars := []oscar.Oscar{_oscar}

	builder := NewBuilder()
	builder.WithID(uuid.New())
	builder.WithTitle("Title")
	builder.WithCoverURL("")
	builder.WithGenres(genres)
	builder.WithDirectors(directors)
	builder.WithYear(2021)
	builder.WithSlogan("Slogan")
	builder.WithResume("Resume")
	builder.WithOscars(oscars)

	_, err := builder.Build()
	assert.NotNil(t, err)
}

func (bt *movieBuilderTests) testWithInvalidSlogan(t *testing.T) {
	_genre, _ := genre.NewBuilder().WithName("Genre").Build()
	_director, _ := director.NewBuilder().WithName("Director").Build()
	_oscar, _ := oscar.NewBuilder().WithName("Best Picture").Build()

	genres := []genre.Genre{_genre}
	directors := []director.Director{_director}
	oscars := []oscar.Oscar{_oscar}

	builder := NewBuilder()
	builder.WithID(uuid.New())
	builder.WithTitle("Title")
	builder.WithCoverURL("https://example.com/cover.jpg")
	builder.WithGenres(genres)
	builder.WithDirectors(directors)
	builder.WithYear(2021)
	builder.WithSlogan("")
	builder.WithResume("Resume")
	builder.WithOscars(oscars)

	_, err := builder.Build()
	assert.NotNil(t, err)
}

func (bt *movieBuilderTests) testWithInvalidResume(t *testing.T) {
	_genre, _ := genre.NewBuilder().WithName("Genre").Build()
	_director, _ := director.NewBuilder().WithName("Director").Build()
	_oscar, _ := oscar.NewBuilder().WithName("Best Picture").Build()

	genres := []genre.Genre{_genre}
	directors := []director.Director{_director}
	oscars := []oscar.Oscar{_oscar}

	builder := NewBuilder()
	builder.WithID(uuid.New())
	builder.WithTitle("Title")
	builder.WithCoverURL("https://example.com/cover.jpg")
	builder.WithGenres(genres)
	builder.WithDirectors(directors)
	builder.WithYear(2021)
	builder.WithSlogan("Slogan")
	builder.WithResume("")
	builder.WithOscars(oscars)

	_, err := builder.Build()
	assert.NotNil(t, err)
}
