package response

import (
	"moviedle/src/core/domain/genre"

	"github.com/google/uuid"
)

type Genre struct {
	ID		*uuid.UUID `json:"id"`
	Name	string `json:"name"`
	Code	string `json:"code"`
}

type genreBuilder struct{}

func NewGenreBuilder() *genreBuilder {
    return &genreBuilder{}
}

func (*genreBuilder) BuildFromDomain(data genre.Genre) Genre {
	return Genre {
		data.ID(),
        data.Name(),
        data.Code(),
	}
}

func (*genreBuilder) BuildFromDomainList(data []genre.Genre) []Genre {
	var genres []Genre
    for _, genreData := range data {
        genres = append(genres, NewGenreBuilder().BuildFromDomain(genreData))
    }
    return genres
}