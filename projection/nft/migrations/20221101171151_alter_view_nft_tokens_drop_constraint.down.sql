ALTER TABLE view_nft_tokens ADD CONSTRAINT view_nft_tokens_denom_id_token_id_key UNIQUE (denom_id, token_id);

DROP INDEX if exists view_nft_tokens_denom_id_token_id_btree_index;