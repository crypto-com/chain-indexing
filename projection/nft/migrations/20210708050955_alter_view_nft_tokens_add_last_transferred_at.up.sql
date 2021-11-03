ALTER TABLE view_nft_tokens
    ADD last_transferred_at BIGINT NOT NULL DEFAULT 0,
    ADD last_transferred_at_block_height BIGINT NOT NULL DEFAULT 0;

UPDATE view_nft_tokens SET
    last_transferred_at = last_edited_at,
    last_transferred_at_block_height = last_edited_at_block_height
WHERE last_transferred_at = 0;
