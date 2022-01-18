CREATE TABLE view_vd_evidences (
    id BIGSERIAL,
    height BIGINT NOT NULL,
    tendermint_address VARCHAR NOT NULL,
    infraction_height BIGINT NOT NULL,
    raw_evidence JSONB NOT NULL,
    PRIMARY KEY (id)
);
