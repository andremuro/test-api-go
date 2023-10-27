-- Table: test_api_go.test_api_1

-- DROP TABLE IF EXISTS test_api_go.test_api_1;

CREATE TABLE IF NOT EXISTS test_api_go.test_api_1
(
    id bigint NOT NULL,
    username character(100)[] COLLATE pg_catalog."default",
    body character(200)[] COLLATE pg_catalog."default",
    created_at daterange NOT NULL,
    CONSTRAINT test_api_1_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS test_api_go.test_api_1
    OWNER to postgres;