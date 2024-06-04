CREATE TABLE IF NOT EXISTS "movie" (
    id uuid NOT NULL
        CONSTRAINT df_movie_id DEFAULT uuid_generate_v4(),
        CONSTRAINT pk_movie_id PRIMARY KEY (id),
    title VARCHAR(255) NOT NULL,
    year integer NOT NULL,
    slogan text NOT NULL,
    resume text NOT NULL
);

COPY movie(id, title, year, slogan, resume) 
    FROM '/fixtures/00004/movie.csv'
    DELIMITER ';'
    CSV HEADER;