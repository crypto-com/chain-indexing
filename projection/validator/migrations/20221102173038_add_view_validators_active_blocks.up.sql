CREATE TABLE view_validator_active_blocks
(
    operator_address VARCHAR,
    block_height     BIGINT  NOT NULL,
    signed           BOOLEAN NOT NULL,
    PRIMARY KEY (operator_address, block_height)
);

CREATE INDEX view_validator_active_blocks_block_height_index
    ON view_validator_active_blocks USING btree (block_height);

CREATE INDEX view_validator_active_blocks_operator_address_block_height_signed_index
    ON view_validator_active_blocks USING btree (operator_address, block_height, signed);

ALTER TABLE view_validators
    DROP COLUMN recent_signed_blocks;

ALTER TABLE view_validators
    DROP COLUMN recent_active_blocks;

