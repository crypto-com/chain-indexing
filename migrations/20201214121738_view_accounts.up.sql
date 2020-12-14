CREATE TABLE view_accounts (
    id BIGSERIAL,
    AccountAddress VARCHAR NOT NULL,
    AccountType VARCHAR ,
    PubKey VARCHAR ,
    AccountNumber BIGINT DEFAULT -1,
    SequenceNumber BIGINT DEFAULT -1,
    AccountBalance BIGINT DEFAULT 0,
    AccountDenom VARCHAR  DEFAULT 'basecro',
    PRIMARY KEY(id),
    UNIQUE(AccountAddress)
);