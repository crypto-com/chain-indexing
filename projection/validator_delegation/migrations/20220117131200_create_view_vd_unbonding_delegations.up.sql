CREATE TABLE view_vd_unbonding_delegations (
    id BIGSERIAL,
    height BIGINT NOT NULL,
    delegator_address VARCHAR NOT NULL,
    validator_address VARCHAR NOT NULL,
    entries JSONB NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(delegator_address, validator_address, height)
);

CREATE INDEX view_vd_unbonding_delegations_validator_height_index ON view_vd_unbonding_delegations USING btree (validator_address, height);

CREATE INDEX view_vd_unbonding_delegations_delegator_height_index ON view_vd_unbonding_delegations USING btree (delegator_address, height);