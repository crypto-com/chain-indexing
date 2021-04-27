CREATE TABLE view_proposal_validators (
    id BIGSERIAL,
    consensus_node_address VARCHAR NOT NULL,
    operator_address VARCHAR NOT NULL,
    initial_delegator_address VARCHAR NOT NULL,
    tendermint_pubkey VARCHAR NOT NULL,
    tendermint_address VARCHAR NOT NULL,
    moniker VARCHAR NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(consensus_node_address, operator_address)
);

CREATE INDEX view_proposal_validators_initial_delegator_address_btree_index ON view_proposal_validators USING btree (initial_delegator_address);

CREATE INDEX view_proposal_validators_operator_address_btree_index ON view_proposal_validators USING btree (operator_address);
