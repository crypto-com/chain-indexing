CREATE INDEX view_transactions_block_height_brin_index ON view_transactions USING brin (block_height);
