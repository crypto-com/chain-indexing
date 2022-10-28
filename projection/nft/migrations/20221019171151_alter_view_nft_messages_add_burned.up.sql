ALTER TABLE view_nft_messages
    ADD burned BOOL NOT NULL DEFAULT FALSE;

CREATE INDEX view_nft_messages_burned_btree_index ON view_nft_messages USING btree (burned);