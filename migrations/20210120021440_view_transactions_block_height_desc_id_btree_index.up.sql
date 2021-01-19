CREATE INDEX view_transactions_block_height_desc_id_btree_index ON view_transactions USING btree (block_height DESC, id);
