CREATE TABLE view_vd_delegations (
    id BIGSERIAL,
    height BIGINT NOT NULL,
    validator_address VARCHAR NOT NULL,
    delegator_address VARCHAR NOT NULL,
    shares VARCHAR NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(validator_address, delegator_address, height)
);

CREATE INDEX view_vd_delegations_height_index ON view_vd_delegations USING btree (height);

CREATE INDEX view_vd_delegations_validator_height_index ON view_vd_delegations USING btree (validator_address, height);

CREATE INDEX view_vd_delegations_delegator_height_index ON view_vd_delegations USING btree (delegator_address, height);