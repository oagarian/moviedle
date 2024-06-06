CREATE TABLE IF NOT EXISTS "director" (
    id uuid NOT NULL
        CONSTRAINT df_dir_id DEFAULT uuid_generate_v4(),
        CONSTRAINT pk_dir_id PRIMARY KEY (id),
    name varchar(255) NOT NULL,
        CONSTRAINT uk_dir_name UNIQUE (name)
);

COPY director(id, name) 
    FROM '/fixtures/00005/director.csv'
    DELIMITER ';'
    CSV HEADER;