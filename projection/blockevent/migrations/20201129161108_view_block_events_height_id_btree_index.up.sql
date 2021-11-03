CREATE INDEX view_block_events_block_height_id_btree_index ON view_block_events USING btree (block_height, id);
