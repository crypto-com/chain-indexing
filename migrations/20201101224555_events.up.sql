CREATE TABLE events (
    id BIGSERIAL,
    name VARCHAR NOT NULL,
    version INT NOT NULL,
    payload JSONB NOT NULL,
    PRIMARY KEY(id)
);