package postgres

import (
	"fmt"
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/domain/movie"
	"moviedle/src/core/messages"
	"moviedle/src/core/ports/secondary"
	"moviedle/src/infra/repository"
	"moviedle/src/infra/repository/postgres/query"
	"strconv"

	"github.com/google/uuid"
)

type movieRepository struct{}

func NewMovieRepository() secondary.MoviePort {
	return &movieRepository{}
}

func (*movieRepository) Get(movieID *uuid.UUID) (movie.Movie, errors.Error) {
	rows, err := repository.Queryx(query.NewMovieQuery().Get(), movieID.String())
	if err != nil {
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()

	if !rows.Next() {
		err := errors.NewFromString(messages.MovieNotFoundErrorMessage)
		return nil, err
	}
	
	var serializedMovie = map[string]interface{}{}
	nativeErr := rows.MapScan(serializedMovie)
	if nativeErr	!= nil {
        return nil, errors.NewUnexpected()
    }
	return newMovieFromMapRows(serializedMovie)
}

func (*movieRepository) All() ([]movie.Movie, errors.Error) {
	rows, err := repository.Queryx(query.NewMovieQuery().All())
    if err!= nil {
        return nil, errors.NewUnexpected()
    }
    defer rows.Close()

    var movies []movie.Movie
    for rows.Next() {
        var serializedMovie = map[string]interface{}{}
        nativeErr := rows.MapScan(serializedMovie)
        if nativeErr    != nil {
            return nil, errors.NewUnexpected()
        }
        movie, err := newMovieFromMapRows(serializedMovie)
        if err!= nil {
            return nil, errors.NewUnexpected()
        }
        movies = append(movies, movie)
    }
    return movies, nil
}

func newMovieFromMapRows(data map[string]interface{}) (movie.Movie, errors.Error) {
	var id uuid.UUID
	if parsedID, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	
	title := fmt.Sprint(data["title"])
	year := fmt.Sprint(data["year"])
	parsedYear, err := strconv.Atoi(year)
	if err != nil {
		return nil, errors.NewUnexpected()
	}
	slogan := fmt.Sprint(data["slogan"])
	resume := fmt.Sprint(data["resume"])
	coverURL := fmt.Sprint(data["cover_url"])

	_directors, dirErr := NewDirectorRepository().ListByMovieID(&id)
	if dirErr != nil {
        return nil, errors.NewUnexpected()
    }

	_genres, genErr := NewGenreRepository().ListByMovieID(&id)
	if genErr != nil {
		return nil, errors.NewUnexpected()
	}

	_oscars, oscErr := NewOscarRepository().ListByMovieID(&id)
	if oscErr != nil {
        return nil, errors.NewUnexpected()
    }
	return movie.NewBuilder().
		    WithID(id).
            WithTitle(title).
            WithYear(uint16(parsedYear)).
            WithSlogan(slogan).
            WithResume(resume).
            WithCoverURL(coverURL).
			WithDirectors(_directors).
			WithGenres(_genres).
			WithOscars(_oscars).
            Build()
}

