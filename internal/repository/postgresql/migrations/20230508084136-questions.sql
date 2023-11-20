
-- +migrate Up
-- ALTER DEFAULT PRIVILEGES FOR ROLE chatgpt GRANT ALL ON TABLES TO PUBLIC;
-- ALTER DEFAULT PRIVILEGES FOR ROLE chatgpt GRANT ALL ON SEQUENCES TO PUBLIC;

CREATE TABLE questions
(
    id      SERIAL PRIMARY KEY,
    question       TEXT,
    rule TEXT,
    category_id VARCHAR(256),
    tag_id VARCHAR(256),
    status SMALLINT DEFAULT 0,
    created_at TIMESTAMP
);

-- +migrate Down

DROP TABLE questions;