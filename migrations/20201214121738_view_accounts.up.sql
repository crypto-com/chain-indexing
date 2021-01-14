CREATE TABLE view_accounts (
    address VARCHAR,
    account_type VARCHAR,
    name VARCHAR NULL,
    pubkey VARCHAR NULL,
    account_number BIGINT,
    sequence_number BIGINT,
    balance JSONB,
    PRIMARY KEY(address)
);