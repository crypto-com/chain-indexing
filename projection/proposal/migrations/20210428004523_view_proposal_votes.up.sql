CREATE TABLE view_proposal_votes (
    id BIGSERIAL,
    proposal_id VARCHAR NOT NULL,
    voter_address VARCHAR NOT NULL,
    maybe_voter_operator_address VARCHAR NULL,
    transaction_hash VARCHAR NOT NULL,
    vote_at_block_height BIGINT NOT NULL,
    vote_at_block_time BIGINT NOT NULL,
    answer VARCHAR NOT NULL,
    histories JSONB NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX view_proposal_votes_vote_at_block_height_btree_index ON view_proposal_votes USING btree(vote_at_block_height);