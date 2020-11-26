CREATE TABLE view_blocks (
    height BIGINT,
    hash VARCHAR NOT NULL,
    time BIGINT NOT NULL,
    app_hash VARCHAR NOT NULL,
    committed_council_nodes JSONB NOT NULL,
    transaction_count INT NOT NULL,
    UNIQUE(hash),
    PRIMARY KEY(height)
);

CREATE INDEX view_blocks_hash_index ON view_blocks(hash);
