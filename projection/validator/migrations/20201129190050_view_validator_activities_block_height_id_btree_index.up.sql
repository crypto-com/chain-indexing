CREATE INDEX view_validator_activities_block_height_id_btree_index ON view_validator_activities USING btree (block_height, id);
