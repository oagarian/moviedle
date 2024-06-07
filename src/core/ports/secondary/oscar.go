package secondary

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/oscar"

	"github.com/google/uuid"
)

type OscarPort interface {
	ListByMovieID(_ID *uuid.UUID) ([]oscar.Oscar, errors.Error)
}