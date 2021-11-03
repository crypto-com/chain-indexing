CREATE TABLE view_account_transactions (
   id BIGSERIAL,
   block_height BIGINT,
   block_hash VARCHAR NOT NULL,
   block_time BIGINT NOT NULL,
   account VARCHAR NOT NULL,
   transaction_hash VARCHAR NOT NULL,
   success BOOLEAN NOT NULL,
   message_types JSONB NOT NULL,
   PRIMARY KEY (id)
)