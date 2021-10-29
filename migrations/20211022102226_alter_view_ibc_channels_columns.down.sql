UPDATE view_ibc_channels
SET status = 'true'
WHERE status = 'OPENED';

UPDATE view_ibc_channels
SET status = 'false'
WHERE status = 'CLOSED';

UPDATE view_ibc_channels
SET status = 'false'
WHERE status = 'NOT_ESTABLISHED';

ALTER TABLE view_ibc_channels
ALTER COLUMN status TYPE BOOLEAN USING status::boolean;