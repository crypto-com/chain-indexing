CREATE TABLE view_vd_validators (
    id BIGSERIAL,
    height BIGINT NOT NULL,
    operator_address VARCHAR NOT NULL,
    consensus_node_address VARCHAR NOT NULL,
    tendermint_address VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    jailed BOOLEAN NOT NULL,
    power VARCHAR NOT NULL,
    unbonding_height BIGINT NOT NULL,
    unbonding_time BIGINT NOT NULL,
    tokens VARCHAR NOT NULL,
    shares VARCHAR NOT NULL,
    min_self_delegation VARCHAR NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(operator_address, height)
    UNIQUE(consensus_node_address, height)
    UNIQUE(tendermint_address, height)
);

CREATE INDEX view_vd_validators_height_index ON view_vd_validators USING btree (height);
