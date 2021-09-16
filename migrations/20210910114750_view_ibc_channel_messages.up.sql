CREATE TABLE view_ibc_channel_messages (
    id BIGSERIAL,
    channel_id VARCHAR NOT NULL,
    block_height BIGINT NOT NULL,
    block_time BIGINT NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    relayer VARCHAR NOT NULL,
    success BOOLEAN NOT NULL,
    error VARCHAR NOT NULL,
    sender VARCHAR NOT NULL,
    receiver VARCHAR NOT NULL,
    denom VARCHAR NOT NULL,
    amount VARCHAR NOT NULL,
    message_type VARCHAR NOT NULL,
    message JSONB NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX view_ibc_channel_messages_channel_id_btree_index ON view_ibc_channel_messages USING btree (channel_id);