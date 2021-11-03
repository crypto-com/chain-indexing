CREATE TABLE view_validator_block_commitments (
    id BIGSERIAL,
    consensus_node_address VARCHAR NOT NULL,
    block_height BIGINT NOT NULL,
    signature VARCHAR NOT NULL,
    timestamp BIGINT NOT NULL,
    PRIMARY KEY (id)
)