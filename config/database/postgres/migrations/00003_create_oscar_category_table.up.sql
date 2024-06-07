CREATE TABLE IF NOT EXISTS "oscar_category" (
    id uuid NOT NULL
        CONSTRAINT df_osc_gen_id DEFAULT uuid_generate_v4(),
        CONSTRAINT pk_osc_gen_id PRIMARY KEY (id),
    name VARCHAR(128) NOT NULL,
        CONSTRAINT uk_osc_gen_name UNIQUE (name),
    code VARCHAR(16) NOT NULL,
        CONSTRAINT uk_osc_gen_code UNIQUE (code)
);

COPY oscar_category(id, name, code)
    FROM '/fixtures/00003/oscar_category.csv'
    DELIMITER ';' 
    CSV HEADER;