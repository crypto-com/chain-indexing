CREATE INDEX view_block_events_block_height_id_brin_index ON view_block_events USING brin (block_height, id);
