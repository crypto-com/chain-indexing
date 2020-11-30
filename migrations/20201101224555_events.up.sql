CREATE TABLE events (
    id BIGSERIAL,
    uuid VARCHAR,
    height INT NOT NULL,
    name VARCHAR NOT NULL,
    version INT NOT NULL,
    payload JSONB NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(uuid)
);

