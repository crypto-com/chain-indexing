ALTER TABLE view_validators
    ADD attention BOOL NOT NULL DEFAULT FALSE;

CREATE INDEX view_validators_attention_btree_index ON view_validators USING btree (attention);
