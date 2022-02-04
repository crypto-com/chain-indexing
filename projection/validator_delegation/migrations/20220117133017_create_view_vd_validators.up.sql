CREATE TABLE view_vd_validators (
    id BIGSERIAL,
    height INT8RANGE NOT NULL,
    operator_address VARCHAR NOT NULL,
    consensus_node_address VARCHAR NOT NULL,
    tendermint_address VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    jailed BOOLEAN NOT NULL,
    power VARCHAR NOT NULL,
    unbonding_height BIGINT NOT NULL,
    unbonding_time BIGINT NOT NULL,
    tokens VARCHAR NOT NULL,
    shares VARCHAR NOT NULL,
    min_self_delegation VARCHAR NOT NULL,
    PRIMARY KEY (id),
    -- Below is a constraint and it is also an index.
    -- It prevents a delegation record appear twice at any given height. 
    EXCLUDE USING gist (
        operator_address WITH =, 
        consensus_node_address WITH =, 
        tendermint_address WITH =, 
        height WITH &&
    )
);
