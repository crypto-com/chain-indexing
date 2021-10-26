UPDATE view_ibc_channels
SET status = 'true'
WHERE status = 'Opened';

UPDATE view_ibc_channels
SET status = 'false'
WHERE status = 'Closed';

UPDATE view_ibc_channels
SET status = 'false'
WHERE status = 'NotEstablished';

ALTER TABLE view_ibc_channels
ALTER COLUMN status TYPE BOOLEAN;