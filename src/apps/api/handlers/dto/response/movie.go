package response

import (
	"moviedle/src/core/domain/movie"

	"github.com/google/uuid"
)

type Movie struct {
	ID 	        *uuid.UUID `json:"id"`
	Title       string `json:"title"`
	CoverURL    string `json:"cover_url"`
	Genres      []Genre `json:"genres"`
	Directors   []Director `json:"directors"`
	Year        uint16 `json:"year"`
	Slogan      string `json:"slogan"`
	Resume      string `json:"resume"`
	Oscars      []Oscar `json:"oscars,omitempty"`
}

type movieBuilder struct{}

func NewMovieBuilder() *movieBuilder {
    return &movieBuilder{}
}

func (*movieBuilder) BuildFromDomain(data movie.Movie) Movie {
	genreData := data.Genres()
	directorData := data.Directors()
	oscarData := data.Oscars()
	return Movie {
        data.ID(),
        data.Title(),
        data.CoverURL(),
        NewGenreBuilder().BuildFromDomainList(genreData),
        NewDirectorBuilder().BuildFromDomainList(directorData),
        data.Year(),
        data.Slogan(),
        data.Resume(),
        NewOscarBuilder().BuildFromDomainList(oscarData),
    }
}