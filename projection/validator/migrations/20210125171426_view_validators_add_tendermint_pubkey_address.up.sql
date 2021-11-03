ALTER TABLE view_validators
    ADD tendermint_pubkey VARCHAR NOT NULL DEFAULT '',
    ADD tendermint_address VARCHAR NOT NULL DEFAULT '';