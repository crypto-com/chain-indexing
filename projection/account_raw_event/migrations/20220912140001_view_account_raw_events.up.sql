CREATE TABLE view_account_raw_events (
   id BIGSERIAL,
   account VARCHAR NOT NULL,
   block_height BIGINT,
   block_hash VARCHAR NOT NULL,
   block_time BIGINT NOT NULL,
   data JSONB NOT NULL,
   PRIMARY KEY(id)
);
create index view_account_raw_events_account_btree_index
    on view_account_raw_events (account);

create index view_account_raw_events_block_hash_btree_index
    on view_account_raw_events (block_hash);