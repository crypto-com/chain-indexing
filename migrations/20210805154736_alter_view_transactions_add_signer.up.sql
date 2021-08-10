ALTER TABLE view_transactions
    ADD signers JSONB NOT NULL DEFAULT '[]';