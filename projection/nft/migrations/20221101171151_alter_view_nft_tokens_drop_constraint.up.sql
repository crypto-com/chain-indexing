ALTER TABLE view_nft_tokens DROP CONSTRAINT view_nft_tokens_ids_btree_index;

CREATE INDEX view_nft_tokens_denom_id_token_id_btree_index ON view_nft_tokens USING btree (denom_id, token_id);