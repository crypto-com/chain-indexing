CREATE INDEX view_transactions_block_height_id_btree_index ON view_transactions USING btree (block_height, id);
