CREATE TABLE view_proposal_params (
    module VARCHAR,
    key VARCHAR,
    value VARCHAR NOT NULL,
    PRIMARY KEY (module, key)
);