ALTER TABLE view_accounts
    ADD usable_balance JSONB NOT NULL DEFAULT '{}';
