CREATE TABLE view_accounts (
    id BIGSERIAL,
    address VARCHAR NOT NULL,
    type VARCHAR,
    pubkey VARCHAR,
    account_number BIGINT,
    sequence_number BIGINT,
    balance VARCHAR,
    PRIMARY KEY(id),
    UNIQUE(address)
);