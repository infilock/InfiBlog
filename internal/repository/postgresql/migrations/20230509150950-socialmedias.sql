-- +migrate Up
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
