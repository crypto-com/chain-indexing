ALTER TABLE view_nft_tokens DROP CONSTRAINT view_nft_tokens_denom_id_token_id_key;

CREATE INDEX view_nft_tokens_denom_id_token_id_btree_index ON view_nft_tokens USING btree (denom_id, token_id);