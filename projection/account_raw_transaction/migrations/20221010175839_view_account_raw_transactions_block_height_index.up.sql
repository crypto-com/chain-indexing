create index view_account_raw_transactions_block_height_btree_index
    on view_account_raw_transactions (block_height);

create index view_account_raw_transactions_account_block_height_btree_index
    on view_account_raw_transactions (account, block_height);
