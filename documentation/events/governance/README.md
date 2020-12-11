# Distribution Module Event List
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
*Name* : MsgSubmitParamUpdateProposalCreated

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                          | Type            | Description                                                       |
| ---------------------------- | --------------- | ----------------------------------------------------------------- |
| `proposalId`                 | *string*        | *(Optional)* Proposal ID                                          |
| `proposerAddress`            | *string*        | Proposer blockchain address                                       |
| `initialDeposit`             | *string*        | Initially deposited CRO Amount in base unit                       |
| `content`                    | *object*        | Content wrapper                                                   |
| `content.title`              | *string*        | Content title indicating the change                               |
| `content.type`               | *string*        | Cosmos SDK type URL                                               |
| `content.description`        | *string*        | Action description                                                |
| `content.changes`            | *array(object)* | List of changes executed                                          |
| `content.changes[].key`      | *string*        | change key                                                        |
| `content.changes[].value`    | *string*        | change value. (Golang type: `json.RawMessage`)                    |
| `content.changes[].subspace` | *string*        | change coding subspace                                            |
| `msgName`                    | *string*        | Blockchain Message type . Value: `MsgSubmitParamUpdateProposal`   |
| `txHash`                     | *string*        | TxID of the blockchain transaction containing the event           |
| `msgIndex`                   | *int*           | message index on the block                                        |
| `name`                       | *string*        | Specific Event Name. Value: `MsgSubmitParamUpdateProposalCreated` |
| `version`                    | *int*           | Event Version. Value: `1`                                         |
| `height`                     | *int64*         | Height of the block containing the transaction                    |
| `uuid`                       | *string*        | Unique ID that is assigned on event creation                      |

*Example* :  
```json
{
    "name": "MsgSubmitParamUpdateProposalCreated",
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
*Name* : MsgSubmitParamUpdateProposalFailed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                          | Type            | Description                                                       |
| ---------------------------- | --------------- | ----------------------------------------------------------------- |
| `proposalId`                 | *string*        | *(Optional)* Proposal ID                                          |
| `proposerAddress`            | *string*        | Proposer blockchain address                                       |
| `initialDeposit`             | *string*        | Initially deposited CRO Amount in base unit                       |
| `content`                    | *object*        | Content wrapper                                                   |
| `content.title`              | *string*        | Content title indicating the change                               |
| `content.type`               | *string*        | Cosmos SDK type URL                                               |
| `content.description`        | *string*        | Action description                                                |
| `content.changes`            | *array(object)* | List of changes executed                                          |
| `content.changes[].key`      | *string*        | change key                                                        |
| `content.changes[].value`    | *string*        | change value. (Golang type: `json.RawMessage`)                    |
| `content.changes[].subspace` | *string*        | change coding subspace                                            |
| `msgName`                    | *string*        | Blockchain Message type . Value: `MsgSubmitParamUpdateProposal`   |
| `txHash`                     | *string*        | TxID of the blockchain transaction containing the event           |
| `msgIndex`                   | *int*           | message index on the block                                        |
| `name`                       | *string*        | Specific Event Name. Value: `MsgSubmitParamUpdateProposalFailed` |
| `version`                    | *int*           | Event Version. Value: `1`                                         |
| `height`                     | *int64*         | Height of the block containing the transaction                    |
| `uuid`                       | *string*        | Unique ID that is assigned on event creation                      |

*Example* : T.B.D  

## event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED
*Name* : MsgSubmitCommunityPoolSpendProposalCreated

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                                                                |
| ------------------ | -------- | ---------------------------------------------------------------------------------------------------------- |
| `delegatorAddress` | *string* | Delegator address. More [here](https://chain.crypto.com/docs/chain-details/module_overview.html#delegator) |
| `recipientAddress` | *string* | Recipient blockchain address                                                                               |
| `validatorAddress` | *string* | Validator address                                                                                          |
| `amount`           | *bigint* | CRO amount in base unit                                                                                    |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgSubmitCommunityPoolSpendProposal`                                              |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event                                                    |
| `msgIndex`         | *int*    | message index on the block                                                                                 |
| `name`             | *string* | Specific Event Name. Value: `MsgSubmitCommunityPoolSpendProposalCreated`                                            |
| `version`          | *int*    | Event Version. Value: `1`                                                                                  |
| `height`           | *int64*  | Height of the block containing the transaction                                                             |
| `uuid`             | *string* | Unique ID that is assigned on event creation                                                               |

*Example* : 
```json
{
    "name": "MsgWithdrawDelegatorRewardCreated",
    "uuid": "f092e7e8-aba4-4705-88ec-6692c407fa3c",
    "amount": "1460368",
    "height": 67598,
    "txHash": "F4700511C2DAE0D4CB50EE16C5ACE6946E83425330064A6C22D1CBA48FA727A9",
    "msgName": "MsgWithdrawDelegatorReward",
    "version": 1,
    "msgIndex": 0,
    "delegatorAddress": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
    "recipientAddress": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
    "validatorAddress": "tcrocncl1923pz03mhjaztgcv3gey0hj0amwx02dyskau52"
}
```  

## event::MSG_WITHDRAW_DELEGATOR_REWARD_FAILED
*Name* : MsgWithdrawDelegatorRewardFailed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                                                                |
| ------------------ | -------- | ---------------------------------------------------------------------------------------------------------- |
| `delegatorAddress` | *string* | Delegator address. More [here](https://chain.crypto.com/docs/chain-details/module_overview.html#delegator) |
| `recipientAddress` | *string* | Recipient blockchain address                                                                               |
| `validatorAddress` | *string* | Validator address                                                                                          |
| `amount`           | *bigint* | CRO amount in base unit. Value : `0`                                                                       |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgWithdrawDelegatorReward`                                              |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event                                                    |
| `msgIndex`         | *int*    | message index on the block                                                                                 |
| `name`             | *string* | Specific Event Name. Value: `MsgWithdrawDelegatorRewardFailed`                                             |
| `version`          | *int*    | Event Version. Value: `1`                                                                                  |
| `height`           | *int64*  | Height of the block containing the transaction                                                             |
| `uuid`             | *string* | Unique ID that is assigned on event creation                                                               |

*Example* : 
```json
{
    "name": "MsgWithdrawDelegatorRewardFailed",
    "uuid": "ea860f82-b819-4ec9-a270-856c8bb9b545",
    "amount": "0",
    "height": 73170,
    "txHash": "913CCA67143D57CB7EECB9A10145B202714F8049425F1FFE8044E9ED27276795",
    "msgName": "MsgWithdrawDelegatorReward",
    "version": 1,
    "msgIndex": 0,
    "delegatorAddress": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
    "recipientAddress": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
    "validatorAddress": "tcrocncl1923pz03mhjaztgcv3gey0hj0amwx02dyskau52"
} 
```

## event::MSG_WITHDRAW_VALIDATOR_COMMISSION_CREATED
*Name* : MsgWithdrawValidatorCommissionCreated

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                         |
| ------------------ | -------- | ------------------------------------------------------------------- |
| `recipientAddress` | *string* | Recipient blockchain address                                        |
| `validatorAddress` | *string* | Validator address                                                   |
| `amount`           | *bigint* | CRO amount in base unit                                             |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgWithdrawValidatorCommission`   |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event             |
| `msgIndex`         | *int*    | message index on the block                                          |
| `name`             | *string* | Specific Event Name. Value: `MsgWithdrawValidatorCommissionCreated` |
| `version`          | *int*    | Event Version. Value: `1`                                           |
| `height`           | *int64*  | Height of the block containing the transaction                      |
| `uuid`             | *string* | Unique ID that is assigned on event creation                        |

*Example* : 
```json
{
    "name": "MsgWithdrawValidatorCommissionCreated",
    "uuid": "4bbfe4dc-ec5e-48a7-bf74-1a7b0356e132",
    "amount": "25487036151",
    "height": 79845,
    "txHash": "A83C5F1A06376D584B395E37D9DEA4DE08C037E4C3F654A617BEE1F9E10E4D9C",
    "msgName": "MsgWithdrawValidatorCommission",
    "version": 1,
    "msgIndex": 1,
    "recipientAddress": "tcro1xwd3k8xterdeft3nxqg92szhpz6vx43q97wcke",
    "validatorAddress": "tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6"
}
```  

## event::MSG_WITHDRAW_VALIDATOR_COMMISSION_FAILED
*Name* : MsgWithdrawValidatorCommissionFailed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                        |
| ------------------ | -------- | ------------------------------------------------------------------ |
| `recipientAddress` | *string* | Recipient blockchain address                                       |
| `validatorAddress` | *string* | Validator address                                                  |
| `amount`           | *bigint* | CRO amount in base unit. Value : `0`                               |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgWithdrawValidatorCommission`  |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event            |
| `msgIndex`         | *int*    | message index on the block                                         |
| `name`             | *string* | Specific Event Name. Value: `MsgWithdrawValidatorCommissionFailed` |
| `version`          | *int*    | Event Version. Value: `1`                                          |
| `height`           | *int64*  | Height of the block containing the transaction                     |
| `uuid`             | *string* | Unique ID that is assigned on event creation                       |

*Example* : 
```json
{
    "name": "MsgWithdrawValidatorCommissionFailed",
    "uuid": "4bbfe4dc-ec5e-48a7-bf74-1a7b0356e132",
    "amount": "0",
    "height": 79845,
    "txHash": "A83C5F1A06376D584B395E37D9DEA4DE08C037E4C3F654A617BEE1F9E10E4D9C",
    "msgName": "MsgWithdrawValidatorCommission",
    "version": 1,
    "msgIndex": 1,
    "recipientAddress": "tcro1xwd3k8xterdeft3nxqg92szhpz6vx43q97wcke",
    "validatorAddress": "tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6"
}
```  

## event::MSG_FUND_COMMUNITY_POOL_CREATED
*Name* : MsgFundCommunityPoolCreated

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key         | Type     | Description                                               |
| ----------- | -------- | --------------------------------------------------------- |
| `depositor` | *string* | Recipient blockchain address                              |
| `amount`    | *bigint* | CRO amount in base unit                                   |
| `msgName`   | *string* | Blockchain Message type. Value: `MsgFundCommunityPool`    |
| `txHash`    | *string* | TxID of the blockchain transaction containing the event   |
| `msgIndex`  | *int*    | message index on the block                                |
| `name`      | *string* | Specific Event Name. Value: `MsgFundCommunityPoolCreated` |
| `version`   | *int*    | Event Version. Value: `1`                                 |
| `height`    | *int64*  | Height of the block containing the transaction            |
| `uuid`      | *string* | Unique ID that is assigned on event creation              |

*Example* : 
```json
{
    "name": "MsgFundCommunityPoolCreated",
    "uuid": "4bbfe4dc-ec5e-48a7-bf74-1a7b0356e132",
    "amount": "10000",
    "height": 79845,
    "txHash": "A83C5F1A06376D584B395E37D9DEA4DE08C037E4C3F654A617BEE1F9E10E4D9C",
    "msgName": "MsgWithdrawValidatorCommission",
    "version": 1,
    "msgIndex": 1,
    "depositor": "tcro1xwd3k8xterdeft3nxqg92szhpz6vx43q97wcke"
}
```

## event::MSG_FUND_COMMUNITY_POOL_FAILED
*Name* : MsgFundCommunityPoolFailed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key         | Type     | Description                                              |
| ----------- | -------- | -------------------------------------------------------- |
| `depositor` | *string* | Recipient blockchain address                             |
| `amount`    | *bigint* | CRO amount in base unit                                  |
| `msgName`   | *string* | Blockchain Message type. Value: `MsgFundCommunityPool`   |
| `txHash`    | *string* | TxID of the blockchain transaction containing the event  |
| `msgIndex`  | *int*    | message index on the block                               |
| `name`      | *string* | Specific Event Name. Value: `MsgFundCommunityPoolFailed` |
| `version`   | *int*    | Event Version. Value: `1`                                |
| `height`    | *int64*  | Height of the block containing the transaction           |
| `uuid`      | *string* | Unique ID that is assigned on event creation             |

*Example* : 
```json
{
    "name": "MsgFundCommunityPoolFailed",
    "uuid": "4bbfe4dc-ec5e-48a7-bf74-1a7b0356e132",
    "amount": "10000",
    "height": 79845,
    "txHash": "A83C5F1A06376D584B395E37D9DEA4DE08C037E4C3F654A617BEE1F9E10E4D9C",
    "msgName": "MsgWithdrawValidatorCommission",
    "version": 1,
    "msgIndex": 1,
    "depositor": "tcro1xwd3k8xterdeft3nxqg92szhpz6vx43q97wcke"
}
```