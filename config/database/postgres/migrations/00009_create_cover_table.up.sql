CREATE TABLE IF NOT EXISTS "cover" (
    id uuid NOT NULL
        CONSTRAINT df_mov_cv_id DEFAULT uuid_generate_v4(),
        CONSTRAINT pk_mov_cv_id PRIMARY KEY (id),
    cover_url TEXT NOT NULL,
    movie_id uuid NOT NULL,
        CONSTRAINT fk_mov_cv_movie_id FOREIGN KEY (movie_id) REFERENCES movie(id)
);

COPY cover(id, cover_url, movie_id) 
    FROM '/fixtures/00009/cover.csv'
    DELIMITER ';'
    CSV HEADER;