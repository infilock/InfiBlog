
-- +migrate Up
ALTER DEFAULT PRIVILEGES FOR ROLE chatgpt GRANT ALL ON TABLES TO PUBLIC;
ALTER DEFAULT PRIVILEGES FOR ROLE chatgpt GRANT ALL ON SEQUENCES TO PUBLIC;

CREATE TABLE articles
(
    id      SERIAL PRIMARY KEY,
    question_id       integer,
    title       TEXT,
    content       TEXT,
    status VARCHAR(30) DEFAULT 'draft',
    created_at TIMESTAMP
);

-- +migrate Down

DROP TABLE articles;