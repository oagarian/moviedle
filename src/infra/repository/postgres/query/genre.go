package query

import ()

type genreQuery struct {}
type GenreQuery interface {
	ListByMovieID() string
}

func NewGenreQuery() GenreQuery {
    return &genreQuery{}
}

func (g *genreQuery) ListByMovieID() string {
    return `
		SELECT 
			g.id AS id,
			g.name AS name,
            g.code AS code
		FROM movie_genre mg
		INNER JOIN genre g 
			ON g.id = mg.genre_id
		WHERE mg.movie_id = $1;
	`
}
