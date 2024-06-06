CREATE TABLE IF NOT EXISTS "movie_director" (
    id uuid NOT NULL
        CONSTRAINT df_mov_dir_id DEFAULT uuid_generate_v4(),
        CONSTRAINT pk_mov_dir_id PRIMARY KEY (id),
    director_id uuid NOT NULL,
        CONSTRAINT fk_mov_dir_director_id FOREIGN KEY (director_id) REFERENCES director(id),
    movie_id uuid NOT NULL,
        CONSTRAINT fk_mov_dir_movie_id FOREIGN KEY (movie_id) REFERENCES movie(id)
);

COPY movie_director(id, director_id, movie_id) 
    FROM '/fixtures/00006/movie_director.csv'
    DELIMITER ';'
    CSV HEADER;