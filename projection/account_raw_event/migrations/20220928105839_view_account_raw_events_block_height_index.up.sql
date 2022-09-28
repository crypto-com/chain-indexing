create index view_account_raw_events_block_height_btree_index
    on view_account_raw_events (block_height);

create index view_account_raw_events_account_block_height_btree_index
    on view_account_raw_events (account, block_height);
