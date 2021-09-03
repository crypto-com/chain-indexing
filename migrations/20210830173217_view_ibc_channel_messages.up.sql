CREATE TABLE view_ibc_channel_messages (
    id BIGSERIAL,
    channel_id VARCHAR NOT NULL,
    block_height BIGINT NOT NULL,
    source_channel VARCHAR NOT NULL,
    destination_channel VARCHAR NOT NULL,
    denom VARCHAR NOT NULL,
    amount VARCHAR NOT NULL,
    success VARCHAR NOT NULL,
    error VARCHAR NOT NULL,
    message_type VARCHAR NOT NULL,
    message JSONB NOT NULL,
    updated_bonded_tokens JSONB NOT NULL,
    PRIMARY KEY(id)
);