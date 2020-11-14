CREATE TABLE events (
    row_id BIGSERIAL,
    id VARCHAR,
    height INT NOT NULL,
    name VARCHAR NOT NULL,
    version INT NOT NULL,
    payload JSONB NOT NULL,
    PRIMARY KEY(row_id),
    UNIQUE(id)
);