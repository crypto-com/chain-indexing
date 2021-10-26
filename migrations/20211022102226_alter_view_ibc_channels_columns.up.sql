ALTER TABLE view_ibc_channels
ALTER COLUMN status TYPE VARCHAR;

UPDATE view_ibc_channels
SET status = 'NOT_ESTABLISHED'
WHERE status = 'false';

UPDATE view_ibc_channels
SET status = 'OPENED'
WHERE status = 'true';