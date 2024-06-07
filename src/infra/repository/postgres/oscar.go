package postgres

import (
	"fmt"
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/oscar"
	"moviedle/src/core/ports/secondary"
	"moviedle/src/infra/repository"
	"moviedle/src/infra/repository/postgres/query"
	"github.com/google/uuid"
)

type oscarRepository struct{}

func NewOscarRepository() secondary.OscarPort {
	return &oscarRepository{}
}

func (g *oscarRepository) ListByMovieID(movieID *uuid.UUID) ([]oscar.Oscar, errors.Error) {
	rows, err := repository.Queryx(query.NewOscarQuery().ListByMovieID(), movieID.String())
	if err != nil {
		return nil, errors.NewUnexpected()
	}

	var oscars = make([]oscar.Oscar, 0)
	for rows.Next() {
		var serializedOscars = map[string]interface{}{}
		if err := rows.MapScan(serializedOscars); err!= nil {
            return nil, errors.NewUnexpected()
        }
		oscar, err := newOscarFromMapRows(serializedOscars)
		if err != nil {
			return nil, errors.NewUnexpected()
		}
		oscars = append(oscars, oscar)
	}
	return oscars, nil
}

func newOscarFromMapRows(data map[string]interface{}) (oscar.Oscar, errors.Error) {
	var id uuid.UUID
	if parsedID, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	
	name := fmt.Sprint(data["name"])
	code := fmt.Sprint(data["code"])

	return oscar.NewBuilder().
			WithID(id).
			WithCode(code).
			WithName(name).
			Build()
}