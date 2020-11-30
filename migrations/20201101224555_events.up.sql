CREATE TABLE events (
    id BIGSERIAL,
    uuid VARCHAR NOT NULL,
    height INT NOT NULL,
    name VARCHAR NOT NULL,
    version INT NOT NULL,
    payload JSONB NOT NULL,
    UNIQUE(uuid, height)
) partition by range (height);

