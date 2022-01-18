CREATE TABLE view_vd_unbonding_validators (
    id BIGSERIAL,
    operator_address VARCHAR NOT NULL UNIQUE,
    unbonding_time BIGINT NOT NULL,
    PRIMARY KEY (id)
);
