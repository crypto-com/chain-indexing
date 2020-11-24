CREATE TABLE view_transactions (
    id BIGSERIAL,
    block_height BIGINT,
    block_hash VARCHAR NOT NULL,
    block_time BIGINT NOT NULL,
    hash VARCHAR NOT NULL,
    success BOOLEAN NOT NULL,
    code INT NOT NULL,
    log VARCHAR NOT NULL,
    fee VARCHAR NOT NULL,
    fee_payer VARCHAR NOT NULL,
    fee_granter VARCHAR NOT NULL,
    gas_wanted BIGINT NOT NULL,
    gas_used BIGINT NOT NULL,
    memo VARCHAR NOT NULL,
    timeout_height BIGINT NOT NULL,
    messages JSONB NOT NULL,
    PRIMARY KEY(id)
);

CREATE INDEX view_transactions_block_height_index ON view_transactions(block_height);

CREATE INDEX view_transactions_order_by_block_height_desc_index ON view_transactions(block_height DESC);
