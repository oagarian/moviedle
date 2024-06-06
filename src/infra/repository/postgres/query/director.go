package query

import ()

type directorQuery struct {}
type DirectorQuery interface {
	ListByMovieID() string
}

func NewDirectorQuery() DirectorQuery {
    return &directorQuery{}
}

func (g *directorQuery) ListByMovieID() string {
    return `
		SELECT 
			g.id AS id,
			g.name AS name
		FROM movie_director md
		INNER JOIN director d
			ON d.id = md.director_id
		WHERE md.movie_id = $1;
	`
}