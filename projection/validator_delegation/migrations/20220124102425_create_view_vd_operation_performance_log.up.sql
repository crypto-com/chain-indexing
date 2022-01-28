CREATE TABLE view_vd_operation_performance_log (
    id BIGSERIAL,
    height BIGINT NOT NULL,
    action VARCHAR NOT NULL,
    duration BIGINT NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX view_vd_operation_performance_log_height_index ON view_vd_operation_performance_log USING btree (height);
