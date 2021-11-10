CREATE TABLE view_nft_token_transfers (
    id BIGSERIAL,
    denom_id VARCHAR NOT NULL,
    token_id VARCHAR NOT NULL,
    block_height BIGINT NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    sender VARCHAR NOT NULL,
    recipient VARCHAR NOT NULL,
    transferred_at BIGINT NOT NULL,
    PRIMARY KEY(id)
)