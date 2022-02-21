CREATE TABLE view_validator_delegation_redelegation_queue (
    id BIGSERIAL,
    completion_time BIGINT NOT NULL UNIQUE,
    -- dvv_triplets is a list of (delegator_addr, src_validator_addr, dst_validator_addr) triplet
    dvv_triplets JSONB NOT NULL,
    PRIMARY KEY (id)
);
