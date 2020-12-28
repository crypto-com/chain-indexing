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
    phase_0_task_registration VARCHAR NOT NULL,
    phase_1_task_node_setup VARCHAR NOT NULL,
    phase_1_task_block_valid_commit VARCHAR NOT NULL,
    phase_2_task_keep_node_active VARCHAR NOT NULL,
    phase_2_task_proposal_vote VARCHAR NOT NULL,
    phase_2_task_network_upgrade VARCHAR NOT NULL,
    phase_2_task_network_upgrade_block_commit VARCHAR NOT NULL,
    phase_1_2_task_commitment_count_rank integer default 0,
    phase_3_task_commitment_count_rank integer default 0,
    task_highest_sequence_rank integer default 0,
    PRIMARY KEY (id),
    UNIQUE (operator_address, consensus_node_address)
);

-- Last_account_sequence