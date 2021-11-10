CREATE TABLE view_nft_denoms (
    id BIGSERIAL,
    denom_id VARCHAR NOT NULL,
    name VARCHAR NULL,
    schema VARCHAR NULL,
    creator VARCHAR NOT NULL,
    created_at BIGINT NOT NULL,
    PRIMARY KEY(denom_id),
    UNIQUE (id)
)