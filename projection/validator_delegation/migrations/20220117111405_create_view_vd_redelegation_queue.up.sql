CREATE TABLE view_vd_redelegation_queue (
    id BIGSERIAL,
    completion_time BIGINT NOT NULL UNIQUE,
    dvv_triplets JSONB NOT NULL,
    PRIMARY KEY (id)
);
