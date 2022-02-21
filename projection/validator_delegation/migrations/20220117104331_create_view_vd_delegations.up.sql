CREATE EXTENSION btree_gist;

CREATE TABLE view_validator_delegation_delegations (
    id BIGSERIAL,
    height INT8RANGE NOT NULL,
    validator_address VARCHAR NOT NULL,
    delegator_address VARCHAR NOT NULL,
    shares VARCHAR NOT NULL,
    PRIMARY KEY (id),
    -- Below is a constraint and it is also an index.
    -- It prevents a delegation record appear twice at any given height. 
    EXCLUDE USING gist (validator_address WITH =, delegator_address WITH =, height WITH &&)
);
