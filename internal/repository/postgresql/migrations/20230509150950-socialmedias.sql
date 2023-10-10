
-- +migrate Up
ALTER DEFAULT PRIVILEGES FOR ROLE chatgpt GRANT ALL ON TABLES TO PUBLIC;
ALTER DEFAULT PRIVILEGES FOR ROLE chatgpt GRANT ALL ON SEQUENCES TO PUBLIC;

CREATE TABLE socialmedias
(
    id      SERIAL PRIMARY KEY,
    question_id       integer,
    type       TEXT,
    content TEXT,
    status SMALLINT DEFAULT 0,
    created_at TIMESTAMP
);

-- +migrate Down

DROP TABLE socialmedias;