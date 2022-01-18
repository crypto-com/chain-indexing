CREATE TABLE view_vd_redelegations (
    id BIGSERIAL,
    height BIGINT NOT NULL,
    delegator_address VARCHAR NOT NULL,
    validator_src_address VARCHAR NOT NULL,
    validator_dst_address VARCHAR NOT NULL,
    entries JSONB NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(delegator_address, validator_src_address, validator_dst_address, height)
);

CREATE INDEX view_vd_redelegations_src_validator_height_index ON view_vd_redelegations USING btree (validator_src_address, height);