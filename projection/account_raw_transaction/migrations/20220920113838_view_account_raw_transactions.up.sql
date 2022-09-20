create table view_account_raw_transactions
(
    id               BIGSERIAL,
    block_height     BIGINT  NOT NULL,
    block_hash       VARCHAR NOT NULL,
    account          VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    success          BOOLEAN NOT NULL,
    code             INTEGER NOT NULL,
    log              VARCHAR NOT NULL,
    fee              jsonb   NOT NULL,
    fee_payer        VARCHAR NOT NULL,
    fee_granter      VARCHAR NOT NULL,
    gas_wanted       BIGINT  NOT NULL,
    gas_used         BIGINT  NOT NULL,
    memo             VARCHAR NOT NULL,
    timeout_height   BIGINT  NOT NULL,
    messages         JSONB   NOT NULL,
    signers          JSONB DEFAULT '[]'::JSONB NOT NULL,
    PRIMARY KEY (id)
);

create index view_account_raw_transactions_account_btree_index
    on view_account_raw_transactions (account);

create index view_account_raw_transactions_account_memo_btree_index
    on view_account_raw_transactions (account, memo);