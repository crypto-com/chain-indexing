create table view_raw_transactions
(
    id             BIGSERIAL,
    block_height   BIGINT                    NOT NULL,
    block_hash     VARCHAR                   NOT NULL,
    block_time     BIGINT                    NOT NULL,
    hash           VARCHAR                   NOT NULL,
    index          INTEGER                   NOT NULL,
    success        BOOLEAN                   NOT NULL,
    code           INTEGER                   NOT NULL,
    log            VARCHAR                   NOT NULL,
    fee            jsonb                     NOT NULL,
    fee_payer      VARCHAR                   NOT NULL,
    fee_granter    VARCHAR                   NOT NULL,
    gas_wanted     BIGINT                    NOT NULL,
    gas_used       BIGINT                    NOT NULL,
    memo           VARCHAR                   NOT NULL,
    timeout_height BIGINT                    NOT NULL,
    messages       JSONB                     NOT NULL,
    signers        JSONB DEFAULT '[]'::JSONB NOT NULL,
    PRIMARY KEY(id)
);

create index view_raw_transactions_hash_btree_index
    on view_raw_transactions (hash);

create index view_raw_transactions_block_height_btree_index
    on view_raw_transactions (block_height);


