CREATE TABLE view_bridge_activities(
    id BIGSERIAL PRIMARY KEY,
    uuid VARCHAR UNIQUE,
    bridge_type VARCHAR NOT NULL,
    link_id VARCHAR NOT NULL,
    source_block_height BIGINT NOT NULL,
    source_block_time BIGINT NOT NULL,
    source_transaction_id VARCHAR NOT NULL,
    source_chain VARCHAR NOT NULL,
    source_address VARCHAR NOT NULL,
    maybe_source_smart_contract_address VARCHAR NULL,
    maybe_destination_block_height BIGINT NULL,
    maybe_destination_block_time BIGINT NULL,
    maybe_destination_transaction_id VARCHAR NULL,
    destination_chain VARCHAR NOT NULL,
    destination_address VARCHAR NOT NULL,
    maybe_destination_smart_contract_address VARCHAR NULL,
    maybe_channel_id VARCHAR NULL,
    amount VARCHAR NOT NULL,
    denom VARCHAR NOT NULL,
    maybe_bridge_fee_amount VARCHAR NULL,
    maybe_bridge_fee_denom VARCHAR NULL,
    status VARCHAR NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE INDEX view_bridge_activities_source_address_btree_index
    ON view_bridge_activities USING btree (source_chain, source_address);

CREATE INDEX view_bridge_activities_destination_address_btree_index
    ON view_bridge_activities USING btree (destination_chain, destination_address);

CREATE INDEX view_bridge_activities_source_transaction_id_btree_index
    ON view_bridge_activities USING btree (source_transaction_id);
