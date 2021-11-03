CREATE TABLE view_ibc_connections (
    id BIGSERIAL,
    connection_id VARCHAR NOT NULL UNIQUE,
    client_id VARCHAR NOT NULL,
    counterparty_connection_id VARCHAR NOT NULL,
    counterparty_client_id VARCHAR NOT NULL,
    counterparty_chain_id VARCHAR NOT NULL,
    PRIMARY KEY(id)
);