CREATE TABLE IF NOT EXISTS "movie_genre" (
    id uuid NOT NULL
        CONSTRAINT df_mov_gen_id DEFAULT uuid_generate_v4(),
        CONSTRAINT pk_mov_gen_id PRIMARY KEY (id),
    genre_id uuid NOT NULL,
        CONSTRAINT fk_mov_gen_genre_id FOREIGN KEY (genre_id) REFERENCES genre(id),
    movie_id uuid NOT NULL,
        CONSTRAINT fk_mov_gen_movie_id FOREIGN KEY (movie_id) REFERENCES movie(id)
);