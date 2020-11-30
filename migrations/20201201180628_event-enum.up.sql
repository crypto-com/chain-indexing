-- migrate Drop don't drop custom types
DO $$ BEGIN
    CREATE TYPE event_name AS ENUM (
        'AccountTransferred',
        'BlockCommissioned',
        'BlockCreated',
        'BlockProposerRewarded',
        'BlockRewarded',
        'GenesisCreated',
        'Minted',
        'MsgCreateValidatorCreated',
        'MsgCreateValidatorFailed',
        'MsgDelegateCreated',
        'MsgDelegateFailed',
        'MsgDepositCreated',
        'MsgEditValidatorCreated',
        'MsgEditValidatorFailed',
        'MsgSendCreated',
        'MsgSubmitParamUpdateProposalCreated',
        'MsgUndelegateCreated',
        'MsgUnjailCreated',
        'MsgUnjailFailed',
        'MsgVoteCreated',
        'MsgWithdrawDelegatorRewardCreated',
        'MsgWithdrawDelegatorRewardFailed',
        'MsgWithdrawValidatorCommissionCreated',
        'PowerChanged',
        'RawBlockCreated',
        'TransactionCreated',
        'TransactionFailed',
        'ValidatorJailed',
        'ValidatorSlashed',

        'FakeEvent',
        'MockEvent'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

alter table events
    alter column name type event_name using name::event_name;
