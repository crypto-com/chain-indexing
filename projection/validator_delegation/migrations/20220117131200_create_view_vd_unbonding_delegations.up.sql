CREATE TABLE view_vd_unbonding_delegations (
    id BIGSERIAL,
    height INT8RANGE NOT NULL,
    delegator_address VARCHAR NOT NULL,
    validator_address VARCHAR NOT NULL,
    entries JSONB NOT NULL,
    PRIMARY KEY (id),
    -- Below is a constraint and it is also an index.
    -- It prevents a delegation record appear twice at any given height. 
    EXCLUDE USING gist (delegator_address WITH =, validator_address WITH =, height WITH &&)
);

CREATE INDEX view_vd_unbonding_delegations_validator_addr_height_index ON view_vd_unbonding_delegations USING gist (height, validator_address);
