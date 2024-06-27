package services

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/movie"
	"moviedle/src/core/ports/primary"
	"moviedle/src/core/ports/secondary"

	"github.com/google/uuid"
)

type movieService struct {
	repository secondary.MoviePort
}

func NewMovieService(repository secondary.MoviePort) primary.MoviePort {
	return &movieService{
        repository: repository,
    }
}

func (s *movieService) Get(_ID *uuid.UUID) (movie.Movie, errors.Error) {
    return s.repository.Get(_ID)
}