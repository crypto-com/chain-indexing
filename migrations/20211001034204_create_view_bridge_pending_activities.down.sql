CREATE TABLE view_bridge_pending_activities(
    id BIGSERIAL PRIMARY KEY,
    block_height BIGINT NOT NULL,
    block_time BIGINT NOT NULL,
    maybe_transaction_id VARCHAR NULL,
    link_id VARCHAR NOT NULL,
    direction VARCHAR NOT NULL,
    from_chain_id VARCHAR NOT NULL,
    maybe_from_address VARCHAR NOT NULL,
    to_chain_id VARCHAR NOT NULL,
    to_address VARCHAR NOT NULL,
    maybe_channel_id VARCHAR NULL,
    token_id VARCHAR NOT NULL,
    amount VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    is_processed BOOL NOT NULL
);

CREATE INDEX view_bridge_pending_activities_link_id_btree_index
    ON view_bridge_pending_activities USING btree (link_id);

CREATE INDEX view_bridge_pending_activities_is_processed_btree_index
    ON view_bridge_pending_activities USING btree (is_processed);
