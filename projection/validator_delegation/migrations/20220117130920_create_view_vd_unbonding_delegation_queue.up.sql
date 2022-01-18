CREATE TABLE view_vd_unbonding_delegation_queue (
    id BIGSERIAL,
    completion_time BIGINT NOT NULL UNIQUE,
    dv_pairs JSONB NOT NULL,
    PRIMARY KEY (id)
);
