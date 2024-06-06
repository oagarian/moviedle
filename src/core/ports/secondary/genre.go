package secondary

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/genre"

	"github.com/google/uuid"
)

type GenrePort interface {
	ListByMovieID(_ID *uuid.UUID) ([]genre.Genre, errors.Error)
}