ALTER TABLE view_validators
ADD total_signed_block BIGINT,
ADD total_active_block BIGINT,
ADD imprecise_up_time NUMERIC;
