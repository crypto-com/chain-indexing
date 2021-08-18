CREATE TABLE view_channels (
    id BIGSERIAL,
    channel_id VARCHAR NOT NULL,
    port_id VARCHAR NOT NULL,
    connection_id VARCHAR NOT NULL,
    counterparty_channel_id VARCHAR NOT NULL,
    counterparty_port_id VARCHAR NOT NULL,
    packet_ordering VARCHAR NOT NULL,
    last_in_packet_sequence BIGINT NOT NULL,
    last_out_packet_sequence BIGINT NOT NULL,
    total_transfer_in_count BIGINT NOT NULL,
    total_transfer_out_count BIGINT NOT NULL,
    total_transfer_out_success_count BIGINT NOT NULL,
    total_transfer_out_success_rate FLOAT NOT NULL,
    last_activity_time BIGINT NOT NULL,
    bonded_tokens JSONB NOT NULL,
    PRIMARY KEY(id),
    UNIQUE (channel_id, port_id)
);