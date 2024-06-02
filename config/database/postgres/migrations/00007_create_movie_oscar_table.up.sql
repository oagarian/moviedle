CREATE TABLE IF NOT EXISTS "movie_oscar" (
    id uuid NOT NULL
        CONSTRAINT df_mov_osc_id DEFAULT uuid_generate_v4(),
        CONSTRAINT pk_mov_osc_id PRIMARY KEY (id),
    oscar_id uuid NOT NULL,
        CONSTRAINT fk_mov_osc_oscar_id FOREIGN KEY (oscar_id) REFERENCES oscar_category(id),
    movie_id uuid NOT NULL,
        CONSTRAINT fk_mov_osc_movie_id FOREIGN KEY (movie_id) REFERENCES movie(id)
);