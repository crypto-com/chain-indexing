CREATE TABLE view_ibc_denom_hash_mapping (
    id BIGSERIAL,
    denom VARCHAR NOT NULL UNIQUE,
    hash VARCHAR NOT NULL UNIQUE,
    PRIMARY KEY(id)
);