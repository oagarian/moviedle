package query

import ()

type movieQuery struct{}
type MovieQuery interface {
	Get() string
	All() string
}

func NewMovieQuery() MovieQuery {
	return &movieQuery{}
}

func (*movieQuery) Get() string {
	return `
		SELECT
			m.id AS id,
			m.title AS title,
            m.year AS year,
            m.slogan AS slogan,
            m.resume AS resume,
			c.cover_url as cover_url
		FROM movie m
		INNER JOIN cover c
			ON c.movie_id = m.id
		WHERE 
			m.id = $1;
	`
}

func (*movieQuery) All() string {
	return `
        SELECT
            m.id AS id,
            m.title AS title,
            m.year AS year,
            m.slogan AS slogan,
            m.resume AS resume,
            c.cover_url as cover_url
        FROM movie m
        INNER JOIN cover c
            ON c.movie_id = m.id;
    `
}