CREATE TABLE view_channels (
    id BIGSERIAL,
    channel_id VARCHAR NOT NULL,
    port_id VARCHAR NOT NULL,
    connection_id VARCHAR,
    counterparty_channel_id VARCHAR,
    counterparty_port_id VARCHAR,
    packet_ordering VARCHAR,
    last_in_packet_sequence BIGINT,
    last_out_packet_sequence BIGINT,
    total_transfer_in_count BIGINT,
    total_transfer_out_count BIGINT,
    total_transfer_out_success_count BIGINT,
    total_transfer_out_success_rate FLOAT,
    PRIMARY KEY(id),
    UNIQUE (channel_id, port_id)
);