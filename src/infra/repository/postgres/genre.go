package postgres

import (
	"fmt"
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/genre"
	"moviedle/src/core/ports/secondary"
	"moviedle/src/infra/repository"
	"moviedle/src/infra/repository/postgres/query"

	"github.com/google/uuid"
)

type genreRepository struct{}

func NewGenreRepository() secondary.GenrePort {
	return &genreRepository{}
}

func (g *genreRepository) ListByMovieID(movieID *uuid.UUID) ([]genre.Genre, errors.Error) {
	rows, err := repository.Queryx(query.NewGenreQuery().ListByMovieID(), movieID.String())
	if err != nil {
		return nil, errors.NewUnexpected()
	}

	var genres = make([]genre.Genre, 0)
	for rows.Next() {
		var serializedGenres = map[string]interface{}{}
		if err := rows.MapScan(serializedGenres); err!= nil {
            return nil, errors.NewUnexpected()
        }
		genre, err := newGenreFromMapRows(serializedGenres)
		if err != nil {
			return nil, errors.NewUnexpected()
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func newGenreFromMapRows(data map[string]interface{}) (genre.Genre, errors.Error) {
	var id uuid.UUID
	if parsedID, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	
	name := fmt.Sprint(data["name"])
	code := fmt.Sprint(data["code"])

	_genreBuilder := genre.NewBuilder()
	_genreBuilder.
	    WithID(id).
		WithCode(code).
		WithName(name)

	return _genreBuilder.Build()
}