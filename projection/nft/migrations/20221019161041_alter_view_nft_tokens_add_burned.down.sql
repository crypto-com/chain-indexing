ALTER TABLE view_nft_tokens
    DROP burned;

DROP INDEX IF EXISTS view_nft_tokens_burned_btree_index;


