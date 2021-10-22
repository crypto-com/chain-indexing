ALTER TABLE view_ibc_channels
DROP COLUMN closed;

ALTER TABLE view_ibc_channels
RENAME COLUMN established TO closed;