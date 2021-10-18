CREATE TABLE view_bridge_pending_activities(
    id BIGSERIAL PRIMARY KEY,
    block_height BIGINT NOT NULL,
    block_time BIGINT NOT NULL,
    maybe_transaction_id VARCHAR NULL,
    bridge_type VARCHAR NOT NULL,
    link_id VARCHAR NOT NULL,
    direction VARCHAR NOT NULL,
    from_chain_id VARCHAR NOT NULL,
    maybe_from_address VARCHAR NOT NULL,
    maybe_from_smart_contract_address VARCHAR NULL,
    to_chain_id VARCHAR NOT NULL,
    to_address VARCHAR NOT NULL,
    maybe_to_smart_contract_address VARCHAR NULL,
    maybe_channel_id VARCHAR NULL,
    amount VARCHAR NOT NULL,
    maybe_denom VARCHAR NULL,
    maybe_bridge_fee_amount VARCHAR NULL,
    maybe_bridge_fee_denom VARCHAR NULL,
    status VARCHAR NOT NULL,
    is_processed BOOL NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE INDEX view_bridge_pending_activities_link_id_btree_index
    ON view_bridge_pending_activities USING btree (link_id, is_processed);

CREATE INDEX view_bridge_pending_activities_is_processed_btree_index
    ON view_bridge_pending_activities USING btree (direction, is_processed);
