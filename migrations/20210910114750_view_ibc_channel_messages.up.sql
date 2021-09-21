CREATE TABLE view_ibc_channel_messages (
    id BIGSERIAL,
    channel_id VARCHAR NOT NULL,
    block_height BIGINT NOT NULL,
    block_time BIGINT NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    relayer VARCHAR,
    success BOOLEAN,
    error VARCHAR,
    sender VARCHAR,
    receiver VARCHAR,
    denom VARCHAR,
    amount VARCHAR,
    message_type VARCHAR NOT NULL,
    message JSONB NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX view_ibc_channel_messages_channel_id_btree_index ON view_ibc_channel_messages USING btree (channel_id);