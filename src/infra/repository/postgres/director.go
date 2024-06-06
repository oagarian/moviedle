package postgres

import (
	"fmt"
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/director"
	"moviedle/src/core/ports/secondary"
	"moviedle/src/infra/repository"
	"moviedle/src/infra/repository/postgres/query"
	"github.com/google/uuid"
)

type directorRepository struct{}

func NewDirectorRepository() secondary.DirectorPort {
	return &directorRepository{}
}

func (g *directorRepository) ListByMovieID(movieID *uuid.UUID) ([]director.Director, errors.Error) {
	rows, err := repository.Queryx(query.NewDirectorQuery().ListByMovieID(), movieID.String())
	if err != nil {
		return nil, errors.NewUnexpected()
	}

	var directors = make([]director.Director, 0)
	for rows.Next() {
		var serializedDirectors = map[string]interface{}{}
		if err := rows.MapScan(serializedDirectors); err!= nil {
            return nil, errors.NewUnexpected()
        }
		director, err := newDirectorFromMapRows(serializedDirectors)
		if err != nil {
			return nil, errors.NewUnexpected()
		}
		directors = append(directors, director)
	}
	return directors, nil
}

func newDirectorFromMapRows(data map[string]interface{}) (director.Director, errors.Error) {
	var id uuid.UUID
	if parsedID, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	
	name := fmt.Sprint(data["name"])

	return director.NewBuilder().
			WithID(id).
			WithName(name).
			Build()
}