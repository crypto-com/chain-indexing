DROP TABLE IF EXISTS view_validator_active_blocks;

ALTER TABLE view_validators
    ADD recent_signed_blocks integer[] DEFAULT '{}';

ALTER TABLE view_validators
    ADD recent_active_blocks integer[] DEFAULT '{}';