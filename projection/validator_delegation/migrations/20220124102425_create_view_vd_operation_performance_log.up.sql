CREATE TABLE view_vd_operation_performance_log (
    id BIGSERIAL,
    height BIGINT NOT NULL,
    action VARCHAR NOT NULL,
    duration BIGINT NOT NULL,
    PRIMARY KEY (id)
);
