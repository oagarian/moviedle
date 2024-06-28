package response

import (
	"moviedle/src/core/domain/movie"

	"github.com/google/uuid"
)

type Movie struct {
	id 	        *uuid.UUID `json:"id"`
	title       string `json:"title"`
	coverURL    string `json:"cover_url"`
	genres      []Genre `json:"genres"`
	directors   []Director `json:"directors"`
	year        uint16 `json:"year"`
	slogan      string `json:"slogan"`
	resume      string `json:"resume"`
	oscars      []Oscar `json:"oscars"`
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