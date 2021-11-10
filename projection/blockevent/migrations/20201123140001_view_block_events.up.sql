CREATE TABLE view_block_events (
   id BIGSERIAL,
   block_height BIGINT,
   block_hash VARCHAR NOT NULL,
   block_time BIGINT NOT NULL,
   data JSONB NOT NULL,
   PRIMARY KEY(id)
);
