CREATE TABLE view_crossfire_validators (
    id BIGSERIAL,
    operator_address VARCHAR NOT NULL,
    consensus_node_address VARCHAR,
    initial_delegator_address VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    jailed BOOL NOT NULL,
    joined_at_block_height BIGINT NOT NULL,
    moniker VARCHAR NOT NULL,
    identity VARCHAR NULL,
    website VARCHAR NULL,
    security_contact VARCHAR NULL,
    details VARCHAR NULL,
    task_phase_1_node_setup VARCHAR NOT NULL,
    task_phase_2_keep_node_active VARCHAR NOT NULL,
    task_phase_2_proposal_vote VARCHAR NOT NULL,
    task_phase_2_network_upgrade VARCHAR NOT NULL,
    rank_task_phase_1_2_commitment_count integer default 0,
    rank_task_phase_3_commitment_count integer default 0,
    rank_task_highest_tx_sent integer default 0,
    PRIMARY KEY (id),
    UNIQUE (operator_address, consensus_node_address)
);

-- Last_account_sequence