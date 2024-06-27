package primary

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/movie"
	"github.com/google/uuid"
)

type MoviePort interface {
	Get(_ID *uuid.UUID) (movie.Movie, errors.Error)
}