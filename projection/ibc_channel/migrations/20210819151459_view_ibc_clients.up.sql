CREATE TABLE view_ibc_clients (
    id BIGSERIAL,
    client_id VARCHAR NOT NULL UNIQUE,
    counterparty_chain_id VARCHAR NOT NULL,
    PRIMARY KEY(id)
);