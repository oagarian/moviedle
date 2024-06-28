package response

import (
	"moviedle/src/core/domain/director"

	"github.com/google/uuid"
)

type Director struct {
	id *uuid.UUID `json:"id"`
	name string `json:"name"`
}

type directorBuilder struct{}

func NewDirectorBuilder() *directorBuilder {
    return &directorBuilder{}
}

func (*directorBuilder) BuildFromDomain(data director.Director) Director {
	return Director {
		data.ID(),
        data.Name(),
	}
}

func (*directorBuilder) BuildFromDomainList(data []director.Director) []Director {
	var directors []Director
    for _, directorData := range data {
        directors = append(directors, NewDirectorBuilder().BuildFromDomain(directorData))
    }
    return directors
}