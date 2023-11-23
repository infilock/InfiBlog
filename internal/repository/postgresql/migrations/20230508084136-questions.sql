-- +migrate Up
CREATE TABLE questions
(
    id      SERIAL PRIMARY KEY,
    question       TEXT,
    rule TEXT,
    category_id VARCHAR(256),
    tag_id VARCHAR(256),
    status VARCHAR(16) DEFAULT 'pending', -- pending OR completed
    created_at TIMESTAMP
);

-- +migrate Down
DROP TABLE questions;
