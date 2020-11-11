CREATE TABLE events (
    id VARCHAR PRIMARY KEY,
    height INT NOT NULL,
    name VARCHAR NOT NULL,
    version INT NOT NULL,
    payload JSONB NOT NULL
);