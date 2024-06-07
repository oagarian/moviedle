package query

import ()

type oscarQuery struct {}
type OscarQuery interface {
	ListByMovieID() string
}

func NewOscarQuery() OscarQuery {
    return &oscarQuery{}
}

func (g *oscarQuery) ListByMovieID() string {
    return `
		SELECT 
			o.id AS id,
			o.name AS name,
            o.code AS code
		FROM movie_oscar mo
		INNER JOIN oscar_category o
			ON o.id = mo.oscar_id
		WHERE mo.movie_id = $1;
	`
}
