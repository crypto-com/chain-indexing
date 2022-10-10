ALTER TABLE view_validators
    ADD recent_active_blocks integer[] DEFAULT array[]::integer[] ;
