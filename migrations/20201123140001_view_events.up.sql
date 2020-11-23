CREATE TABLE view_events (
   id BIGSERIAL,
   block_height BIGINT,
   block_hash VARCHAR NOT NULL,
   block_time BIGINT NOT NULL,
   events JSONB NOT NULL,
   PRIMARY KEY(id)
);
