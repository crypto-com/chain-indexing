CREATE TABLE view_crossfire_validators (
    id BIGSERIAL,
    operator_address VARCHAR NOT NULL,
    consensus_node_address VARCHAR,
    initial_delegator_address VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    jailed BOOL NOT NULL,
    joined_at_block_height BIGINT NOT NULL,
    joined_at_block_height_time BIGINT NOT NULL,
    moniker VARCHAR NOT NULL,
    identity VARCHAR NULL,
    isPrimaryNode BOOLEAN default true,
    phase_0_task_registration integer default -1,
    phase_1_task_node_setup integer default -1,
    phase_1_task_block_valid_commit integer default -1,
    phase_2_task_keep_node_active integer default -1,
    phase_2_task_proposal_vote integer default -1,
    phase_2_task_network_upgrade integer default -1,
    phase_2_task_network_upgrade_block_commit integer default -1,
    phase_1_2_task_commitment_count_rank integer default -1,
    phase_3_task_commitment_count_rank integer default -1,
    task_highest_sequence_rank integer default -1,
    PRIMARY KEY (id),
    UNIQUE (operator_address, consensus_node_address)
);

-- Last_account_sequence