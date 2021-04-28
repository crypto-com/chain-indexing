CREATE TABLE view_proposal_depositors (
    id BIGSERIAL,
    proposal_id VARCHAR NOT NULL,
    depositor_address VARCHAR NOT NULL,
    maybe_depositor_operator_address VARCHAR NULL,
    transaction_hash VARCHAR NOT NULL,
    deposit_at_block_height BIGINT NOT NULL,
    deposit_at_block_time BIGINT NOT NULL,
    amount JSONB NOT NULL,
    PRIMARY KEY (id)
)
