CREATE TABLE view_examples
(
    id      BIGSERIAL PRIMARY KEY,
    address VARCHAR NOT NULL,
    balance JSONB
);

CREATE INDEX view_examples_address_index
    ON view_examples USING btree (address);
