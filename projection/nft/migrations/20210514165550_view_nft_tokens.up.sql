CREATE TABLE view_nft_tokens (
    id BIGSERIAL,
    denom_id VARCHAR NOT NULL,
    token_id VARCHAR NOT NULL,
    drop VARCHAR NULL,
    burned BOOL NOT NULL,
    name VARCHAR NULL,
    uri VARCHAR NULL,
    data VARCHAR NULL,
    minter VARCHAR NOT NULL,
    owner VARCHAR NOT NULL,
    minted_at BIGINT NOT NULL,
    last_edited_at BIGINT NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(denom_id, token_id)
);

CREATE INDEX view_nft_tokens_drop_btree_index ON view_nft_tokens USING btree (drop);

CREATE INDEX view_nft_tokens_token_id_btree_index ON view_nft_tokens USING btree (token_id);
