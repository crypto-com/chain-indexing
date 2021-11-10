ALTER TABLE view_ibc_channels 
RENAME COLUMN total_transfer_in_count TO total_relay_in_count;

ALTER TABLE view_ibc_channels 
RENAME COLUMN total_transfer_out_count TO total_relay_out_count;

ALTER TABLE view_ibc_channels 
RENAME COLUMN total_transfer_out_success_count TO total_relay_out_success_count;

ALTER TABLE view_ibc_channels 
RENAME COLUMN total_transfer_out_success_rate TO total_relay_out_success_rate;