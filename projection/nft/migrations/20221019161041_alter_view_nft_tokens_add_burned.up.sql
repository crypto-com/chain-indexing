ALTER TABLE view_nft_tokens
    ADD burned BOOL NOT NULL DEFAULT FALSE;

CREATE INDEX view_nft_tokens_burned_btree_index ON view_nft_tokens USING btree (burned);