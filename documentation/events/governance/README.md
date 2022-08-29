# Governance Module Event List

-   [event::MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED](#event_msg_submit_param_change_proposal_created)
-   [event::MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_FAILED](#event_msg_submit_param_change_proposal_failed)
-   [event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED](#event_msg_submit_community_pool_spend_proposal_created)
-   [event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_FAILED](#event_msg_submit_community_pool_spend_proposal_failed)
-   [event::MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED](#event_msg_submit_software_upgrade_proposal_created)
-   [event::MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_FAILED](#event_msg_submit_software_upgrade_proposal_failed)
-   [event::MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_CREATED](#event_msg_submit_cancel_software_upgrade_proposal_created)
-   [event::MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_FAILED](#event_msg_submit_cancel_software_upgrade_proposal_failed)
-   [event::MSG_DEPOSIT_CREATED](#event_msg_deposit_created)
-   [event::MSG_DEPOSIT_FAILED](#event_msg_deposit_failed)
-   [event::MSG_VOTE_CREATED](#event_msg_vote_created)
-   [event::MSG_VOTE_FAILED](#event_msg_vote_failed)
-   [event::PROPOSAL_ENDED](#event_proposal_ended)
-   [event::PROPOSAL_INACTIVED](#event_proposal_inactived)

## event::MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED

_Name_ : /cosmos.params.v1beta1.ParameterChangeProposal.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                          | Type            | Description                                                                          |
| ---------------------------- | --------------- | ------------------------------------------------------------------------------------ |
| `proposalId`                 | _string_        | _(Optional)_ Proposal ID                                                             |
| `proposerAddress`            | _string_        | Proposer blockchain address                                                          |
| `initialDeposit`             | _string_        | Initially deposited CRO Amount in base unit                                          |
| `content`                    | _object_        | Content wrapper                                                                      |
| `content.title`              | _string_        | Content title indicating the change                                                  |
| `content.@type`              | _string_        | Cosmos SDK type URL                                                                  |
| `content.description`        | _string_        | Action description                                                                   |
| `content.changes`            | _array(object)_ | List of changes executed                                                             |
| `content.changes[].key`      | _string_        | change key                                                                           |
| `content.changes[].value`    | _string_        | change value. (Golang type: `json.RawMessage`)                                       |
| `content.changes[].subspace` | _string_        | change coding subspace                                                               |
| `msgName`                    | _string_        | Blockchain Message type . Value: `MsgSubmitParamUpdateProposal`                      |
| `txHash`                     | _string_        | TxID of the blockchain transaction containing the event                              |
| `msgIndex`                   | _int_           | message index on the block                                                           |
| `name`                       | _string_        | Specific Event Name. Value: `/cosmos.params.v1beta1.ParameterChangeProposal.Created` |
| `version`                    | _int_           | Event Version. Value: `1`                                                            |
| `height`                     | _int64_         | Height of the block containing the transaction                                       |
| `uuid`                       | _string_        | Unique ID that is assigned on event creation                                         |

_Example_ :

```json
{
    "name": "/cosmos.params.v1beta1.ParameterChangeProposal.Created",
    "uuid": "0f45f978-0478-4232-83bf-b3525c7b45e6",
    "height": 463,
    "txHash": "0DD5DE992A07F232F735FCA5E634347DE4EBAB7E2FFA75BFDBEC117DA42E2802",
    "content": {
        "@type": "/cosmos.params.v1beta1.ParameterChangeProposal",
        "title": "Staking Param Change",
        "changes": [
            {
                "key": "MaxValidators",
                "value": "151",
                "subspace": "staking"
            }
        ],
        "description": "Update max validators"
    },
    "msgName": "MsgSubmitParamUpdateProposal",
    "version": 1,
    "msgIndex": 0,
    "proposalId": "1",
    "initialDeposit": "0",
    "proposerAddress": "tcro1j7pej8kplem4wt50p4hfvndhuw5jprxxn5625q"
}
```

## event::MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_FAILED

_Name_ : /cosmos.params.v1beta1.ParameterChangeProposal.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                          | Type            | Description                                                                         |
| ---------------------------- | --------------- | ----------------------------------------------------------------------------------- |
| `proposalId`                 | _string_        | _(Optional)_ Proposal ID                                                            |
| `proposerAddress`            | _string_        | Proposer blockchain address                                                         |
| `initialDeposit`             | _string_        | Initially deposited CRO Amount in base unit                                         |
| `content`                    | _object_        | Content wrapper                                                                     |
| `content.@title`             | _string_        | Content title indicating the change                                                 |
| `content.type`               | _string_        | Cosmos SDK type URL                                                                 |
| `content.description`        | _string_        | Action description                                                                  |
| `content.changes`            | _array(object)_ | List of changes executed                                                            |
| `content.changes[].key`      | _string_        | change key                                                                          |
| `content.changes[].value`    | _string_        | change value. (Golang type: `json.RawMessage`)                                      |
| `content.changes[].subspace` | _string_        | change coding subspace                                                              |
| `msgName`                    | _string_        | Blockchain Message type . Value: `MsgSubmitParamUpdateProposal`                     |
| `txHash`                     | _string_        | TxID of the blockchain transaction containing the event                             |
| `msgIndex`                   | _int_           | message index on the block                                                          |
| `name`                       | _string_        | Specific Event Name. Value: `/cosmos.params.v1beta1.ParameterChangeProposal.Failed` |
| `version`                    | _int_           | Event Version. Value: `1`                                                           |
| `height`                     | _int64_         | Height of the block containing the transaction                                      |
| `uuid`                       | _string_        | Unique ID that is assigned on event creation                                        |

_Example_ : T.B.D

## event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED

_Name_ : /cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                        | Type     | Description                                                                                   |
| -------------------------- | -------- | --------------------------------------------------------------------------------------------- |
| `proposalId`               | _string_ | _(Optional)_ Proposal ID                                                                      |
| `proposerAddress`          | _string_ | Proposer blockchain address                                                                   |
| `initialDeposit`           | _string_ | Initially deposited CRO Amount in base unit                                                   |
| `content`                  | _object_ | Content wrapper                                                                               |
| `content.@type`            | _string_ | Cosmos SDK type URL                                                                           |
| `content.title`            | _string_ | Content title indicating the change                                                           |
| `content.description`      | _string_ | Action description                                                                            |
| `content.recipientAddress` | _string_ | Recipient blockchain address                                                                  |
| `content.amount`           | _string_ | Recipient CRO Amount in base unit                                                             |
| `msgName`                  | _string_ | Blockchain Message type . Value: `MsgSubmitCommunityPoolSpendProposal`                        |
| `txHash`                   | _string_ | TxID of the blockchain transaction containing the event                                       |
| `msgIndex`                 | _int_    | message index on the block                                                                    |
| `name`                     | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Created` |
| `version`                  | _int_    | Event Version. Value: `1`                                                                     |
| `height`                   | _int64_  | Height of the block containing the transaction                                                |
| `uuid`                     | _string_ | Unique ID that is assigned on event creation                                                  |

_Example_ : T.B.D

## event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_FAILED

_Name_ : /cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                        | Type     | Description                                                                                  |
| -------------------------- | -------- | -------------------------------------------------------------------------------------------- |
| `proposalId`               | _string_ | _(Optional)_ Proposal ID                                                                     |
| `proposerAddress`          | _string_ | Proposer blockchain address                                                                  |
| `initialDeposit`           | _string_ | Initially deposited CRO Amount in base unit                                                  |
| `content`                  | _object_ | Content wrapper                                                                              |
| `content.@type`            | _string_ | Cosmos SDK type URL                                                                          |
| `content.title`            | _string_ | Content title indicating the change                                                          |
| `content.description`      | _string_ | Action description                                                                           |
| `content.recipientAddress` | _string_ | Recipient blockchain address                                                                 |
| `content.amount`           | _string_ | Recipient CRO Amount in base unit                                                            |
| `msgName`                  | _string_ | Blockchain Message type . Value: `MsgSubmitCommunityPoolSpendProposal`                       |
| `txHash`                   | _string_ | TxID of the blockchain transaction containing the event                                      |
| `msgIndex`                 | _int_    | message index on the block                                                                   |
| `name`                     | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Failed` |
| `version`                  | _int_    | Event Version. Value: `1`                                                                    |
| `height`                   | _int64_  | Height of the block containing the transaction                                               |
| `uuid`                     | _string_ | Unique ID that is assigned on event creation                                                 |

_Example_ : T.B.D

## event::MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED

_Name_ : /cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                   | Type     | Description                                                                           |
| --------------------- | -------- | ------------------------------------------------------------------------------------- |
| `proposalId`          | _string_ | _(Optional)_ Proposal ID                                                              |
| `proposerAddress`     | _string_ | Proposer blockchain address                                                           |
| `initialDeposit`      | _bigint_ | Initially deposited CRO Amount in base unit                                           |
| `content`             | _object_ | Content wrapper                                                                       |
| `content.@type`       | _string_ | Cosmos SDK type URL                                                                   |
| `content.title`       | _string_ | Content title indicating the change                                                   |
| `content.description` | _string_ | Action description                                                                    |
| `content.plan`        | _object_ | Content plan details wrapper                                                          |
| `content.plan.name`   | _string_ | Plan name                                                                             |
| `content.plan.title`  | _int64_  | Plan Title. Golang type: `utctime.UTCTime`                                            |
| `content.plan.height` | _int64_  | Plan proposal block height                                                            |
| `content.plan.info`   | _string_ | Plan information                                                                      |
| `msgName`             | _string_ | Blockchain Message type . Value: `MsgSubmitSoftwareUpgradeProposal`                   |
| `txHash`              | _string_ | TxID of the blockchain transaction containing the event                               |
| `msgIndex`            | _int_    | message index on the block                                                            |
| `name`                | _string_ | Specific Event Name. Value: `/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.Created` |
| `version`             | _int_    | Event Version. Value: `1`                                                             |
| `height`              | _int64_  | Height of the block containing the transaction                                        |
| `uuid`                | _string_ | Unique ID that is assigned on event creation                                          |

_Example_ : T.B.D

## event::MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_FAILED

_Name_ : /cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                   | Type     | Description                                                                          |
| --------------------- | -------- | ------------------------------------------------------------------------------------ |
| `proposalId`          | _string_ | _(Optional)_ Proposal ID                                                             |
| `proposerAddress`     | _string_ | Proposer blockchain address                                                          |
| `initialDeposit`      | _bigint_ | Initially deposited CRO Amount in base unit                                          |
| `content`             | _object_ | Content wrapper                                                                      |
| `content.@type`       | _string_ | Cosmos SDK type URL                                                                  |
| `content.title`       | _string_ | Content title indicating the change                                                  |
| `content.description` | _string_ | Action description                                                                   |
| `content.plan`        | _object_ | Content plan details wrapper                                                         |
| `content.plan.name`   | _string_ | Plan name                                                                            |
| `content.plan.title`  | _int64_  | Plan Title. Golang type: `utctime.UTCTime`                                           |
| `content.plan.height` | _int64_  | Plan proposal block height                                                           |
| `content.plan.info`   | _string_ | Plan information                                                                     |
| `msgName`             | _string_ | Blockchain Message type . Value: `MsgSubmitSoftwareUpgradeProposal`                  |
| `txHash`              | _string_ | TxID of the blockchain transaction containing the event                              |
| `msgIndex`            | _int_    | message index on the block                                                           |
| `name`                | _string_ | Specific Event Name. Value: `/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.Failed` |
| `version`             | _int_    | Event Version. Value: `1`                                                            |
| `height`              | _int64_  | Height of the block containing the transaction                                       |
| `uuid`                | _string_ | Unique ID that is assigned on event creation                                         |

_Example_ : T.B.D

## event::MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_CREATED

_Name_ : /cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                   | Type     | Description                                                                                 |
| --------------------- | -------- | ------------------------------------------------------------------------------------------- |
| `proposalId`          | _string_ | _(Optional)_ Proposal ID                                                                    |
| `proposerAddress`     | _string_ | Proposer blockchain address                                                                 |
| `initialDeposit`      | _bigint_ | Initially deposited CRO Amount in base unit                                                 |
| `content`             | _object_ | Content wrapper                                                                             |
| `content.@type`       | _string_ | Cosmos SDK type URL                                                                         |
| `content.title`       | _string_ | Content title indicating the change                                                         |
| `content.description` | _string_ | Action description                                                                          |
| `msgName`             | _string_ | Blockchain Message type . Value: `MsgSubmitCancelSoftwareUpgradeProposal`                   |
| `txHash`              | _string_ | TxID of the blockchain transaction containing the event                                     |
| `msgIndex`            | _int_    | message index on the block                                                                  |
| `name`                | _string_ | Specific Event Name. Value: `/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Created` |
| `version`             | _int_    | Event Version. Value: `1`                                                                   |
| `height`              | _int64_  | Height of the block containing the transaction                                              |
| `uuid`                | _string_ | Unique ID that is assigned on event creation                                                |

_Example_ : T.B.D

## event::MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_FAILED

_Name_ : /cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                   | Type     | Description                                                                                |
| --------------------- | -------- | ------------------------------------------------------------------------------------------ |
| `proposalId`          | _string_ | _(Optional)_ Proposal ID                                                                   |
| `proposerAddress`     | _string_ | Proposer blockchain address                                                                |
| `initialDeposit`      | _bigint_ | Initially deposited CRO Amount in base unit                                                |
| `content`             | _object_ | Content wrapper                                                                            |
| `content.@type`       | _string_ | Cosmos SDK type URL                                                                        |
| `content.title`       | _string_ | Content title indicating the change                                                        |
| `content.description` | _string_ | Action description                                                                         |
| `msgName`             | _string_ | Blockchain Message type . Value: `MsgSubmitCancelSoftwareUpgradeProposal`                  |
| `txHash`              | _string_ | TxID of the blockchain transaction containing the event                                    |
| `msgIndex`            | _int_    | message index on the block                                                                 |
| `name`                | _string_ | Specific Event Name. Value: `/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Failed` |
| `version`             | _int_    | Event Version. Value: `1`                                                                  |
| `height`              | _int64_  | Height of the block containing the transaction                                             |
| `uuid`                | _string_ | Unique ID that is assigned on event creation                                               |

_Example_ : T.B.D

## event::MSG_DEPOSIT_CREATED

_Name_ : /cosmos.gov.v1beta1.MsgDeposit.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key          | Type     | Description                                                          |
| ------------ | -------- | -------------------------------------------------------------------- |
| `proposalId` | _string_ | _(Optional)_ Proposal ID                                             |
| `depositor`  | _string_ | Depositor blockchain address                                         |
| `amount`     | _bigint_ | CRO Amount in base unit                                              |
| `msgName`    | _string_ | Blockchain Message type . Value: `MsgDeposit`                        |
| `txHash`     | _string_ | TxID of the blockchain transaction containing the event              |
| `msgIndex`   | _int_    | message index on the block                                           |
| `name`       | _string_ | Specific Event Name. Value: `/cosmos.gov.v1beta1.MsgDeposit.Created` |
| `version`    | _int_    | Event Version. Value: `1`                                            |
| `height`     | _int64_  | Height of the block containing the transaction                       |
| `uuid`       | _string_ | Unique ID that is assigned on event creation                         |

_Example_ :

```json
{
    "name": "/cosmos.gov.v1beta1.MsgDeposit.Created",
    "uuid": "404aa33b-7e0d-4700-a5e8-b9be6236e580",
    "amount": "100000000000",
    "height": 566,
    "txHash": "90CB157FD0CD6C9DF596F81CDF91ECEED056FD96F6BFAB563AFE3989A489BD90",
    "msgName": "MsgDeposit",
    "version": 1,
    "msgIndex": 0,
    "depositor": "tcro1j7pej8kplem4wt50p4hfvndhuw5jprxxn5625q",
    "proposalId": "1"
}
```

## event::MSG_DEPOSIT_FAILED

_Name_ : /cosmos.gov.v1beta1.MsgDeposit.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key          | Type     | Description                                                         |
| ------------ | -------- | ------------------------------------------------------------------- |
| `proposalId` | _string_ | _(Optional)_ Proposal ID                                            |
| `depositor`  | _string_ | Depositor blockchain address                                        |
| `amount`     | _bigint_ | CRO Amount in base unit                                             |
| `msgName`    | _string_ | Blockchain Message type . Value: `MsgDeposit`                       |
| `txHash`     | _string_ | TxID of the blockchain transaction containing the event             |
| `msgIndex`   | _int_    | message index on the block                                          |
| `name`       | _string_ | Specific Event Name. Value: `/cosmos.gov.v1beta1.MsgDeposit.Failed` |
| `version`    | _int_    | Event Version. Value: `1`                                           |
| `height`     | _int64_  | Height of the block containing the transaction                      |
| `uuid`       | _string_ | Unique ID that is assigned on event creation                        |

_Example_ :

```json
{
    "name": "/cosmos.gov.v1beta1.MsgDeposit.Failed",
    "uuid": "404aa33b-7e0d-4700-a5e8-b9be6236e580",
    "amount": "100000000000",
    "height": 566,
    "txHash": "90CB157FD0CD6C9DF596F81CDF91ECEED056FD96F6BFAB563AFE3989A489BD90",
    "msgName": "MsgDeposit",
    "version": 1,
    "msgIndex": 0,
    "depositor": "tcro1j7pej8kplem4wt50p4hfvndhuw5jprxxn5625q",
    "proposalId": "1"
}
```

## event::MSG_VOTE_CREATED

_Name_ : /cosmos.gov.v1beta1.MsgVote.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key          | Type     | Description                                                       |
| ------------ | -------- | ----------------------------------------------------------------- |
| `proposalId` | _string_ | Proposal ID                                                       |
| `voter`      | _string_ | Voter blockchain address                                          |
| `option`     | _string_ | Voter option selected                                             |
| `msgName`    | _string_ | Blockchain Message type . Value: `MsgVote`                        |
| `txHash`     | _string_ | TxID of the blockchain transaction containing the event           |
| `msgIndex`   | _int_    | message index on the block                                        |
| `name`       | _string_ | Specific Event Name. Value: `/cosmos.gov.v1beta1.MsgVote.Created` |
| `version`    | _int_    | Event Version. Value: `1`                                         |
| `height`     | _int64_  | Height of the block containing the transaction                    |
| `uuid`       | _string_ | Unique ID that is assigned on event creation                      |

_Example_ :

```json
{
    "name": "/cosmos.gov.v1beta1.MsgVote.Created",
    "uuid": "a160ce23-b8e2-4ccd-90a4-a6cf51d67687",
    "voter": "tcro16yzcz3ty94awr7nr2txek9dp2klp2av9egkgxn",
    "height": 673,
    "option": "VOTE_OPTION_YES",
    "txHash": "2FBACF3315B481316ADB1DF739E4151E4533ACAC60A08A402430988B6BA9D557",
    "msgName": "MsgVote",
    "version": 1,
    "msgIndex": 0,
    "proposalId": "1"
}
```

## event::MSG_VOTE_FAILED

_Name_ : /cosmos.gov.v1beta1.MsgVote.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key          | Type     | Description                                                      |
| ------------ | -------- | ---------------------------------------------------------------- |
| `proposalId` | _string_ | Proposal ID                                                      |
| `voter`      | _string_ | Voter blockchain address                                         |
| `option`     | _string_ | Voter option selected                                            |
| `msgName`    | _string_ | Blockchain Message type . Value: `MsgVote`                       |
| `txHash`     | _string_ | TxID of the blockchain transaction containing the event          |
| `msgIndex`   | _int_    | message index on the block                                       |
| `name`       | _string_ | Specific Event Name. Value: `/cosmos.gov.v1beta1.MsgVote.Failed` |
| `version`    | _int_    | Event Version. Value: `1`                                        |
| `height`     | _int64_  | Height of the block containing the transaction                   |
| `uuid`       | _string_ | Unique ID that is assigned on event creation                     |

_Example_ :

```json
{
    "name": "/cosmos.gov.v1beta1.MsgVote.Failed",
    "uuid": "a160ce23-b8e2-4ccd-90a4-a6cf51d67687",
    "voter": "tcro16yzcz3ty94awr7nr2txek9dp2klp2av9egkgxn",
    "height": 673,
    "option": "VOTE_OPTION_YES",
    "txHash": "2FBACF3315B481316ADB1DF739E4151E4533ACAC60A08A402430988B6BA9D557",
    "msgName": "MsgVote",
    "version": 1,
    "msgIndex": 0,
    "proposalId": "1"
}
```

## event::PROPOSAL_ENDED

_Name_ : ProposalEnded

_Type_ : [Base](../README.md#Understanding_an_EVENT)

_Structure_ :

| Key          | Type     | Description                                    |
| ------------ | -------- | ---------------------------------------------- |
| `proposalId` | _string_ | Proposal ID                                    |
| `result`     | _string_ | Proposal end result                            |
| `name`       | _string_ | Specific Event Name. Value: `ProposalEnded`    |
| `version`    | _int_    | Event Version. Value: `1`                      |
| `height`     | _int64_  | Height of the block containing the transaction |
| `uuid`       | _string_ | Unique ID that is assigned on event creation   |

_Example_ : T.B.D

## event::PROPOSAL_INACTIVED

_Name_ : ProposalInactived

_Type_ : [Base](../README.md#Understanding_an_EVENT)

_Structure_ :

| Key          | Type     | Description                                     |
| ------------ | -------- | ----------------------------------------------- |
| `proposalId` | _string_ | Proposal ID                                     |
| `result`     | _string_ | Proposal end result                             |
| `name`       | _string_ | Specific Event Name. Value: `ProposalInactived` |
| `version`    | _int_    | Event Version. Value: `1`                       |
| `height`     | _int64_  | Height of the block containing the transaction  |
| `uuid`       | _string_ | Unique ID that is assigned on event creation    |

_Example_ : T.B.D
