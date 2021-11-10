CREATE TABLE view_nft_messages (
    id BIGSERIAL,
    block_height BIGINT,
    block_hash VARCHAR NOT NULL,
    block_time BIGINT NOT NULL,
    denom_id VARCHAR NOT NULL,
    maybe_token_id VARCHAR NULL,
    maybe_drop VARCHAR NULL,
    transaction_hash VARCHAR NOT NULL,
    success BOOLEAN NOT NULL,
    message_index INT NOT NULL,
    message_type VARCHAR NOT nULL,
    data JSONB NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX view_nft_messages_drop_message_type_btree_index ON view_nft_messages USING btree (maybe_drop, message_type);
CREATE INDEX view_nft_messages_denom_message_type_btree_index ON view_nft_messages USING btree (denom_id, maybe_token_id);
CREATE INDEX view_nft_messages_denom_token_message_type_btree_index ON view_nft_messages USING btree (denom_id, maybe_token_id, message_type);
