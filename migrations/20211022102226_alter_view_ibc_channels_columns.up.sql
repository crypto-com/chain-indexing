ALTER TABLE view_ibc_channels
ALTER COLUMN status TYPE VARCHAR;

UPDATE view_ibc_channels
SET status = 'NotEstablished'
WHERE status = 'false';

UPDATE view_ibc_channels
SET status = 'Opened'
WHERE status = 'true';