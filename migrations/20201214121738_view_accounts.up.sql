CREATE TABLE view_accounts (
    id BIGSERIAL,
    address VARCHAR NOT NULL,
    account_type VARCHAR,
    pubkey VARCHAR,
    account_number BIGINT,
    sequence_number BIGINT,
    balance VARCHAR,
    PRIMARY KEY(id),
    UNIQUE(address)
);