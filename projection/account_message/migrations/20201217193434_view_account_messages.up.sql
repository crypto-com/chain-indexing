CREATE TABLE view_account_messages (
    id BIGSERIAL,
    block_height BIGINT,
    block_hash VARCHAR NOT NULL,
    block_time BIGINT NOT NULL,
    account VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    success BOOLEAN NOT NULL,
    message_index INT NOT NULL,
    message_type VARCHAR NOT nULL,
    data JSONB NOT NULL,
    PRIMARY KEY (id)
)