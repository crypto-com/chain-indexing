CREATE TABLE view_channels (
    id BIGSERIAL,
    channel_id VARCHAR NOT NULL,
    port_id VARCHAR NOT NULL,
    connection_id VARCHAR NOT NULL,
    counterparty_channel_id VARCHAR NOT NULL,
    counterparty_port_id VARCHAR NOT NULL,
    success_packet_count BIGINT NOT NULL,
    failure_packet_count BIGINT NOT NULL,
    PRIMARY KEY(id),
    UNIQUE (channel_id, port_id)
);