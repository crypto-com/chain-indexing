CREATE TABLE view_accounts (
    id BIGSERIAL,
    account_address VARCHAR NOT NULL,
    account_type VARCHAR ,
    pubkey VARCHAR ,
    account_number BIGINT DEFAULT -1,
    sequence_number BIGINT DEFAULT -1,
    balance BIGINT DEFAULT 0,
    bonded_balance BIGINT DEFAULT 0,
    unbonding_balance BIGINT DEFAULT 0,
    unbonded_balance BIGINT DEFAULT 0,
    account_denom VARCHAR  DEFAULT 'basecro',
    PRIMARY KEY(id),
    UNIQUE(account_address)
);