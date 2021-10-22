ALTER TABLE view_ibc_channels
RENAME COLUMN status TO established;

ALTER TABLE view_ibc_channels
ADD COLUMN closed BOOLEAN NOT NULL DEFAULT FALSE;