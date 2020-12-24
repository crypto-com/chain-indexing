CREATE TABLE crossfire_chain_stats
(
    metric VARCHAR  NOT NULL,
    value VARCHAR NOT NULL DEFAULT '0',
    PRIMARY KEY (metric)
)

