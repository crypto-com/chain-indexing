ALTER TABLE view_nft_tokens
    ADD minted_at_block_height BIGINT NOT NULL,
    ADD last_edited_at_block_height BIGINT NOT NULL;