ALTER TABLE view_accounts
    ADD name VARCHAR NULL;
ALTER TABLE view_accounts
    ADD pubkey VARCHAR NULL;
ALTER TABLE view_accounts
    ADD account_type VARCHAR;
ALTER TABLE view_accounts
    ADD balance JSONB;
