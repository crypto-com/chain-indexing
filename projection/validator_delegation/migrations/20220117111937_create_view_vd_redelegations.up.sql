CREATE TABLE view_vd_redelegations (
    id BIGSERIAL,
    height INT8RANGE NOT NULL,
    delegator_address VARCHAR NOT NULL,
    validator_src_address VARCHAR NOT NULL,
    validator_dst_address VARCHAR NOT NULL,
    entries JSONB NOT NULL,
    PRIMARY KEY (id),
    -- Below is a constraint and it is also an index.
    -- It prevents a delegation record appear twice at any given height. 
    EXCLUDE USING gist (
        delegator_address WITH =, 
        validator_src_address WITH =, 
        validator_dst_address WITH =, 
        height WITH &&
    )
);
