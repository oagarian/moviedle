package secondary

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/director"

	"github.com/google/uuid"
)

type DirectorPort interface {
	ListByMovieID(_ID *uuid.UUID) ([]director.Director, errors.Error)
}