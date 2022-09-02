# Governance Module Event List
  - [event::MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED](#event_msg_submit_param_change_proposal_created)
  - [event::MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_FAILED](#event_msg_submit_param_change_proposal_failed)
  - [event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED](#event_msg_submit_community_pool_spend_proposal_created)
  - [event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_FAILED](#event_msg_submit_community_pool_spend_proposal_failed)
  - [event::MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED](#event_msg_submit_software_upgrade_proposal_created)
  - [event::MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_FAILED](#event_msg_submit_software_upgrade_proposal_failed)
  - [event::MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_CREATED](#event_msg_submit_cancel_software_upgrade_proposal_created)
  - [event::MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_FAILED](#event_msg_submit_cancel_software_upgrade_proposal_failed)
  - [event::MSG_DEPOSIT_CREATED](#event_msg_deposit_created)
  - [event::MSG_DEPOSIT_FAILED](#event_msg_deposit_failed)
  - [event::MSG_VOTE_CREATED](#event_msg_vote_created)
  - [event::MSG_VOTE_FAILED](#event_msg_vote_failed)
  - [event::PROPOSAL_ENDED](#event_proposal_ended)
  - [event::PROPOSAL_INACTIVED](#event_proposal_inactived)
  

## event::MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED
*Name* : /cosmos.params.v1beta1.ParameterChangeProposal.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                          | Type            | Description                                                                          |
| ---------------------------- | --------------- | ------------------------------------------------------------------------------------ |
| `proposalId`                 | *string*        | _(Optional)_ Proposal ID                                                             |
| `proposerAddress`            | *string*        | Proposer blockchain address                                                          |
| `initialDeposit`             | *string*        | Initially deposited CRO Amount in base unit                                          |
| `content`                    | *object*        | Content wrapper                                                                      |
| `content.title`              | *string*        | Content title indicating the change                                                  |
| `content.@type`              | *string*        | Cosmos SDK type URL                                                                  |
| `content.description`        | *string*        | Action description                                                                   |
| `content.changes`            | *array(object)* | List of changes executed                                                             |
| `content.changes[].key`      | *string*        | change key                                                                           |
| `content.changes[].value`    | *string*        | change value. (Golang type: `json.RawMessage`)                                       |
| `content.changes[].subspace` | *string*        | change coding subspace                                                               |
| `msgName`                    | *string*        | Blockchain Message type . Value: `MsgSubmitParamUpdateProposal`                      |
| `txHash`                     | *string*        | TxID of the blockchain transaction containing the event                              |
| `msgIndex`                   | *int*           | message index on the block                                                           |
| `name`                       | *string*        | Specific Event Name. Value: `/cosmos.params.v1beta1.ParameterChangeProposal.Created` |
| `version`                    | *int*           | Event Version. Value: `1`                                                            |
| `height`                     | *int64*         | Height of the block containing the transaction                                       |
| `uuid`                       | *string*        | Unique ID that is assigned on event creation                                         |

*Example* : 
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
*Name* : /cosmos.params.v1beta1.ParameterChangeProposal.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                          | Type            | Description                                                                         |
| ---------------------------- | --------------- | ----------------------------------------------------------------------------------- |
| `proposalId`                 | *string*        | _(Optional)_ Proposal ID                                                            |
| `proposerAddress`            | *string*        | Proposer blockchain address                                                         |
| `initialDeposit`             | *string*        | Initially deposited CRO Amount in base unit                                         |
| `content`                    | *object*        | Content wrapper                                                                     |
| `content.@title`             | *string*        | Content title indicating the change                                                 |
| `content.type`               | *string*        | Cosmos SDK type URL                                                                 |
| `content.description`        | *string*        | Action description                                                                  |
| `content.changes`            | *array(object)* | List of changes executed                                                            |
| `content.changes[].key`      | *string*        | change key                                                                          |
| `content.changes[].value`    | *string*        | change value. (Golang type: `json.RawMessage`)                                      |
| `content.changes[].subspace` | *string*        | change coding subspace                                                              |
| `msgName`                    | *string*        | Blockchain Message type . Value: `MsgSubmitParamUpdateProposal`                     |
| `txHash`                     | *string*        | TxID of the blockchain transaction containing the event                             |
| `msgIndex`                   | *int*           | message index on the block                                                          |
| `name`                       | *string*        | Specific Event Name. Value: `/cosmos.params.v1beta1.ParameterChangeProposal.Failed` |
| `version`                    | *int*           | Event Version. Value: `1`                                                           |
| `height`                     | *int64*         | Height of the block containing the transaction                                      |
| `uuid`                       | *string*        | Unique ID that is assigned on event creation                                        |

*Example* : T.B.D  

## event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED
*Name* : /cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* :

| Key                        | Type     | Description                                                                                   |
| -------------------------- | -------- | --------------------------------------------------------------------------------------------- |
| `proposalId`               | *string* | _(Optional)_ Proposal ID                                                                      |
| `proposerAddress`          | *string* | Proposer blockchain address                                                                   |
| `initialDeposit`           | *string* | Initially deposited CRO Amount in base unit                                                   |
| `content`                  | *object* | Content wrapper                                                                               |
| `content.@type`            | *string* | Cosmos SDK type URL                                                                           |
| `content.title`            | *string* | Content title indicating the change                                                           |
| `content.description`      | *string* | Action description                                                                            |
| `content.recipientAddress` | *string* | Recipient blockchain address                                                                  |
| `content.amount`           | *string* | Recipient CRO Amount in base unit                                                             |
| `msgName`                  | *string* | Blockchain Message type . Value: `MsgSubmitCommunityPoolSpendProposal`                        |
| `txHash`                   | *string* | TxID of the blockchain transaction containing the event                                       |
| `msgIndex`                 | *int*    | message index on the block                                                                    |
| `name`                     | *string* | Specific Event Name. Value: `/cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Created` |
| `version`                  | *int*    | Event Version. Value: `1`                                                                     |
| `height`                   | *int64*  | Height of the block containing the transaction                                                |
| `uuid`                     | *string* | Unique ID that is assigned on event creation                                                  |

*Example* : T.B.D

## event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_FAILED
*Name* : /cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                        | Type     | Description                                                                                  |
| -------------------------- | -------- | -------------------------------------------------------------------------------------------- |
| `proposalId`               | *string* | _(Optional)_ Proposal ID                                                                     |
| `proposerAddress`          | *string* | Proposer blockchain address                                                                  |
| `initialDeposit`           | *string* | Initially deposited CRO Amount in base unit                                                  |
| `content`                  | *object* | Content wrapper                                                                              |
| `content.@type`            | *string* | Cosmos SDK type URL                                                                          |
| `content.title`            | *string* | Content title indicating the change                                                          |
| `content.description`      | *string* | Action description                                                                           |
| `content.recipientAddress` | *string* | Recipient blockchain address                                                                 |
| `content.amount`           | *string* | Recipient CRO Amount in base unit                                                            |
| `msgName`                  | *string* | Blockchain Message type . Value: `MsgSubmitCommunityPoolSpendProposal`                       |
| `txHash`                   | *string* | TxID of the blockchain transaction containing the event                                      |
| `msgIndex`                 | *int*    | message index on the block                                                                   |
| `name`                     | *string* | Specific Event Name. Value: `/cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Failed` |
| `version`                  | *int*    | Event Version. Value: `1`                                                                    |
| `height`                   | *int64*  | Height of the block containing the transaction                                               |
| `uuid`                     | *string* | Unique ID that is assigned on event creation                                                 |

*Example* : T.B.D  

## event::MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED
*Name* : /cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                   | Type     | Description                                                                           |
| --------------------- | -------- | ------------------------------------------------------------------------------------- |
| `proposalId`          | *string* | _(Optional)_ Proposal ID                                                              |
| `proposerAddress`     | *string* | Proposer blockchain address                                                           |
| `initialDeposit`      | *bigint* | Initially deposited CRO Amount in base unit                                           |
| `content`             | *object* | Content wrapper                                                                       |
| `content.@type`       | *string* | Cosmos SDK type URL                                                                   |
| `content.title`       | *string* | Content title indicating the change                                                   |
| `content.description` | *string* | Action description                                                                    |
| `content.plan`        | *object* | Content plan details wrapper                                                          |
| `content.plan.name`   | *string* | Plan name                                                                             |
| `content.plan.title`  | *int64*  | Plan Title. Golang type: `utctime.UTCTime`                                            |
| `content.plan.height` | *int64*  | Plan proposal block height                                                            |
| `content.plan.info`   | *string* | Plan information                                                                      |
| `msgName`             | *string* | Blockchain Message type . Value: `MsgSubmitSoftwareUpgradeProposal`                   |
| `txHash`              | *string* | TxID of the blockchain transaction containing the event                               |
| `msgIndex`            | *int*    | message index on the block                                                            |
| `name`                | *string* | Specific Event Name. Value: `/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.Created` |
| `version`             | *int*    | Event Version. Value: `1`                                                             |
| `height`              | *int64*  | Height of the block containing the transaction                                        |
| `uuid`                | *string* | Unique ID that is assigned on event creation                                          |

*Example* : T.B.D  

## event::MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_FAILED
*Name* : /cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                   | Type     | Description                                                                          |
| --------------------- | -------- | ------------------------------------------------------------------------------------ |
| `proposalId`          | *string* | _(Optional)_ Proposal ID                                                             |
| `proposerAddress`     | *string* | Proposer blockchain address                                                          |
| `initialDeposit`      | *bigint* | Initially deposited CRO Amount in base unit                                          |
| `content`             | *object* | Content wrapper                                                                      |
| `content.@type`       | *string* | Cosmos SDK type URL                                                                  |
| `content.title`       | *string* | Content title indicating the change                                                  |
| `content.description` | *string* | Action description                                                                   |
| `content.plan`        | *object* | Content plan details wrapper                                                         |
| `content.plan.name`   | *string* | Plan name                                                                            |
| `content.plan.title`  | *int64*  | Plan Title. Golang type: `utctime.UTCTime`                                           |
| `content.plan.height` | *int64*  | Plan proposal block height                                                           |
| `content.plan.info`   | *string* | Plan information                                                                     |
| `msgName`             | *string* | Blockchain Message type . Value: `MsgSubmitSoftwareUpgradeProposal`                  |
| `txHash`              | *string* | TxID of the blockchain transaction containing the event                              |
| `msgIndex`            | *int*    | message index on the block                                                           |
| `name`                | *string* | Specific Event Name. Value: `/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.Failed` |
| `version`             | *int*    | Event Version. Value: `1`                                                            |
| `height`              | *int64*  | Height of the block containing the transaction                                       |
| `uuid`                | *string* | Unique ID that is assigned on event creation                                         |

*Example* : T.B.D  

## event::MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_CREATED
*Name* : /cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                   | Type     | Description                                                                                 |
| --------------------- | -------- | ------------------------------------------------------------------------------------------- |
| `proposalId`          | *string* | _(Optional)_ Proposal ID                                                                    |
| `proposerAddress`     | *string* | Proposer blockchain address                                                                 |
| `initialDeposit`      | *bigint* | Initially deposited CRO Amount in base unit                                                 |
| `content`             | *object* | Content wrapper                                                                             |
| `content.@type`       | *string* | Cosmos SDK type URL                                                                         |
| `content.title`       | *string* | Content title indicating the change                                                         |
| `content.description` | *string* | Action description                                                                          |
| `msgName`             | *string* | Blockchain Message type . Value: `MsgSubmitCancelSoftwareUpgradeProposal`                   |
| `txHash`              | *string* | TxID of the blockchain transaction containing the event                                     |
| `msgIndex`            | *int*    | message index on the block                                                                  |
| `name`                | *string* | Specific Event Name. Value: `/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Created` |
| `version`             | *int*    | Event Version. Value: `1`                                                                   |
| `height`              | *int64*  | Height of the block containing the transaction                                              |
| `uuid`                | *string* | Unique ID that is assigned on event creation                                                |

*Example* : T.B.D  

## event::MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_FAILED
*Name* : /cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                   | Type     | Description                                                                                |
| --------------------- | -------- | ------------------------------------------------------------------------------------------ |
| `proposalId`          | *string* | _(Optional)_ Proposal ID                                                                   |
| `proposerAddress`     | *string* | Proposer blockchain address                                                                |
| `initialDeposit`      | *bigint* | Initially deposited CRO Amount in base unit                                                |
| `content`             | *object* | Content wrapper                                                                            |
| `content.@type`       | *string* | Cosmos SDK type URL                                                                        |
| `content.title`       | *string* | Content title indicating the change                                                        |
| `content.description` | *string* | Action description                                                                         |
| `msgName`             | *string* | Blockchain Message type . Value: `MsgSubmitCancelSoftwareUpgradeProposal`                  |
| `txHash`              | *string* | TxID of the blockchain transaction containing the event                                    |
| `msgIndex`            | *int*    | message index on the block                                                                 |
| `name`                | *string* | Specific Event Name. Value: `/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Failed` |
| `version`             | *int*    | Event Version. Value: `1`                                                                  |
| `height`              | *int64*  | Height of the block containing the transaction                                             |
| `uuid`                | *string* | Unique ID that is assigned on event creation                                               |

*Example* : T.B.D  

## event::MSG_DEPOSIT_CREATED
*Name* : /cosmos.gov.v1beta1.MsgDeposit.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key          | Type     | Description                                                          |
| ------------ | -------- | -------------------------------------------------------------------- |
| `proposalId` | *string* | _(Optional)_ Proposal ID                                             |
| `depositor`  | *string* | Depositor blockchain address                                         |
| `amount`     | *bigint* | CRO Amount in base unit                                              |
| `msgName`    | *string* | Blockchain Message type . Value: `MsgDeposit`                        |
| `txHash`     | *string* | TxID of the blockchain transaction containing the event              |
| `msgIndex`   | *int*    | message index on the block                                           |
| `name`       | *string* | Specific Event Name. Value: `/cosmos.gov.v1beta1.MsgDeposit.Created` |
| `version`    | *int*    | Event Version. Value: `1`                                            |
| `height`     | *int64*  | Height of the block containing the transaction                       |
| `uuid`       | *string* | Unique ID that is assigned on event creation                         |

*Example* :  
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
*Name* : /cosmos.gov.v1beta1.MsgDeposit.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key          | Type     | Description                                                         |
| ------------ | -------- | ------------------------------------------------------------------- |
| `proposalId` | *string* | _(Optional)_ Proposal ID                                            |
| `depositor`  | *string* | Depositor blockchain address                                        |
| `amount`     | *bigint* | CRO Amount in base unit                                             |
| `msgName`    | *string* | Blockchain Message type . Value: `MsgDeposit`                       |
| `txHash`     | *string* | TxID of the blockchain transaction containing the event             |
| `msgIndex`   | *int*    | message index on the block                                          |
| `name`       | *string* | Specific Event Name. Value: `/cosmos.gov.v1beta1.MsgDeposit.Failed` |
| `version`    | *int*    | Event Version. Value: `1`                                           |
| `height`     | *int64*  | Height of the block containing the transaction                      |
| `uuid`       | *string* | Unique ID that is assigned on event creation                        |

*Example* :  
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
*Name* : /cosmos.gov.v1beta1.MsgVote.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key          | Type     | Description                                                       |
| ------------ | -------- | ----------------------------------------------------------------- |
| `proposalId` | *string* | Proposal ID                                                       |
| `voter`      | *string* | Voter blockchain address                                          |
| `option`     | *string* | Voter option selected                                             |
| `msgName`    | *string* | Blockchain Message type . Value: `MsgVote`                        |
| `txHash`     | *string* | TxID of the blockchain transaction containing the event           |
| `msgIndex`   | *int*    | message index on the block                                        |
| `name`       | *string* | Specific Event Name. Value: `/cosmos.gov.v1beta1.MsgVote.Created` |
| `version`    | *int*    | Event Version. Value: `1`                                         |
| `height`     | *int64*  | Height of the block containing the transaction                    |
| `uuid`       | *string* | Unique ID that is assigned on event creation                      |

*Example* :  
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
*Name* : /cosmos.gov.v1beta1.MsgVote.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key          | Type     | Description                                                      |
| ------------ | -------- | ---------------------------------------------------------------- |
| `proposalId` | *string* | Proposal ID                                                      |
| `voter`      | *string* | Voter blockchain address                                         |
| `option`     | *string* | Voter option selected                                            |
| `msgName`    | *string* | Blockchain Message type . Value: `MsgVote`                       |
| `txHash`     | *string* | TxID of the blockchain transaction containing the event          |
| `msgIndex`   | *int*    | message index on the block                                       |
| `name`       | *string* | Specific Event Name. Value: `/cosmos.gov.v1beta1.MsgVote.Failed` |
| `version`    | *int*    | Event Version. Value: `1`                                        |
| `height`     | *int64*  | Height of the block containing the transaction                   |
| `uuid`       | *string* | Unique ID that is assigned on event creation                     |

*Example* :  
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
*Name* : ProposalEnded

*Type* : [Base](../README.md#Understanding_an_EVENT)

*Structure* : 

| Key          | Type     | Description                                    |
| ------------ | -------- | ---------------------------------------------- |
| `proposalId` | *string* | Proposal ID                                    |
| `result`     | *string* | Proposal end result                            |
| `name`       | *string* | Specific Event Name. Value: `ProposalEnded`    |
| `version`    | *int*    | Event Version. Value: `1`                      |
| `height`     | *int64*  | Height of the block containing the transaction |
| `uuid`       | *string* | Unique ID that is assigned on event creation   |

*Example* :  T.B.D  

## event::PROPOSAL_INACTIVED
*Name* : ProposalInactived

*Type* : [Base](../README.md#Understanding_an_EVENT)

*Structure* : 

| Key          | Type     | Description                                     |
| ------------ | -------- | ----------------------------------------------- |
| `proposalId` | *string* | Proposal ID                                     |
| `result`     | *string* | Proposal end result                             |
| `name`       | *string* | Specific Event Name. Value: `ProposalInactived` |
| `version`    | *int*    | Event Version. Value: `1`                       |
| `height`     | *int64*  | Height of the block containing the transaction  |
| `uuid`       | *string* | Unique ID that is assigned on event creation    |

*Example* : T.B.D
