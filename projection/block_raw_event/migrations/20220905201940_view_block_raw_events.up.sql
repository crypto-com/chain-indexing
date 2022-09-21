CREATE TABLE view_block_raw_events (
   id BIGSERIAL,
   block_height BIGINT,
   block_hash VARCHAR NOT NULL,
   block_time BIGINT NOT NULL,
   from_result VARCHAR NOT NULL,
   data JSONB NOT NULL,
   PRIMARY KEY(id)
);
create index view_block_raw_events_block_height_btree_index
    on view_block_raw_events (block_height);
    
create index view_block_raw_events_block_hash_btree_index
    on view_block_raw_events (block_hash);