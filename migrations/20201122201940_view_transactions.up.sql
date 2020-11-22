CREATE TABLE view_transactions (
     block_height BIGINT,
     block_hash VARCHAR NOT NULL,
     block_time BIGINT NOT NULL,
     hash VARCHAR NOT NULL,
     success BOOLEAN NOT NULL,
     code INT NOT NULL,
     log VARCHAR NOT NULL,
     fee VARCHAR NOT NULL,
     fee_payer VARCHAR NOT NULL,
     fee_granter VARCHAR NOT NULL,
     gas_wanted INT NOT NULL,
     gas_used INT NOT NULL,
     memo VARCHAR NOT NULL,
     timeout_height BIGINT NOT NULL,
     messages JSONB NOT NULL,
     PRIMARY KEY(block_height)
);
