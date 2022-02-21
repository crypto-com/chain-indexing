CREATE TABLE view_validator_delegation_unbonding_delegation_queue (
    id BIGSERIAL,
    completion_time BIGINT NOT NULL UNIQUE,
    -- dv_pairs is a list of (delegator_addr, validator_addr) pair
    dv_pairs JSONB NOT NULL,
    PRIMARY KEY (id)
);
