CREATE TABLE projections (
    id VARCHAR NOT NULL,
    last_handled_event_height BIGINT NOT NULL,
    PRIMARY KEY(id)
);