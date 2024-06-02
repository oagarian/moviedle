CREATE TABLE IF NOT EXISTS "genre" (
    id uuid NOT NULL
        CONSTRAINT df_film_gen_id DEFAULT uuid_generate_v4(),
        CONSTRAINT pk_film_gen_id PRIMARY KEY (id),
    name VARCHAR(128) NOT NULL,
        CONSTRAINT uk_film_gen_name UNIQUE (name),
    code VARCHAR(12) NOT NULL,
        CONSTRAINT uk_film_gen_code UNIQUE (name)
);