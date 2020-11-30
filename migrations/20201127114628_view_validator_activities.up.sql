CREATE TABLE view_validator_activities (
    id BIGSERIAL,
    block_height BIGINT NOT NULL,
    block_hash VARCHAR NOT NULL,
    block_time BIGINT NOT NULL,
    transaction_hash VARCHAR NULL,
    operator_address VARCHAR NOT NULL,
    success BOOLEAN NOT NULL,
    data JSONB NOT NULL,
    PRIMARY KEY (id)
);

