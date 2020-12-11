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
| `content.@type`              | *string*        | Cosmos SDK type URL                                               |
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

| Key                          | Type            | Description                                                      |
| ---------------------------- | --------------- | ---------------------------------------------------------------- |
| `proposalId`                 | *string*        | *(Optional)* Proposal ID                                         |
| `proposerAddress`            | *string*        | Proposer blockchain address                                      |
| `initialDeposit`             | *string*        | Initially deposited CRO Amount in base unit                      |
| `content`                    | *object*        | Content wrapper                                                  |
| `content.@title`             | *string*        | Content title indicating the change                              |
| `content.type`               | *string*        | Cosmos SDK type URL                                              |
| `content.description`        | *string*        | Action description                                               |
| `content.changes`            | *array(object)* | List of changes executed                                         |
| `content.changes[].key`      | *string*        | change key                                                       |
| `content.changes[].value`    | *string*        | change value. (Golang type: `json.RawMessage`)                   |
| `content.changes[].subspace` | *string*        | change coding subspace                                           |
| `msgName`                    | *string*        | Blockchain Message type . Value: `MsgSubmitParamUpdateProposal`  |
| `txHash`                     | *string*        | TxID of the blockchain transaction containing the event          |
| `msgIndex`                   | *int*           | message index on the block                                       |
| `name`                       | *string*        | Specific Event Name. Value: `MsgSubmitParamUpdateProposalFailed` |
| `version`                    | *int*           | Event Version. Value: `1`                                        |
| `height`                     | *int64*         | Height of the block containing the transaction                   |
| `uuid`                       | *string*        | Unique ID that is assigned on event creation                     |

*Example* : T.B.D  

## event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED
*Name* : MsgSubmitCommunityPoolSpendProposalCreated

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                        | Type     | Description                                                              |
| -------------------------- | -------- | ------------------------------------------------------------------------ |
| `proposalId`               | *string* | *(Optional)* Proposal ID                                                 |
| `proposerAddress`          | *string* | Proposer blockchain address                                              |
| `initialDeposit`           | *string* | Initially deposited CRO Amount in base unit                              |
| `content`                  | *object* | Content wrapper                                                          |
| `content.@type`            | *string* | Cosmos SDK type URL                                                      |
| `content.title`            | *string* | Content title indicating the change                                      |
| `content.description`      | *string* | Action description                                                       |
| `content.recipientAddress` | *string* | Recipient blockchain address                                             |
| `content.amount`           | *string* | Recipient CRO Amount in base unit                                        |
| `msgName`                  | *string* | Blockchain Message type . Value: `MsgSubmitCommunityPoolSpendProposal`   |
| `txHash`                   | *string* | TxID of the blockchain transaction containing the event                  |
| `msgIndex`                 | *int*    | message index on the block                                               |
| `name`                     | *string* | Specific Event Name. Value: `MsgSubmitCommunityPoolSpendProposalCreated` |
| `version`                  | *int*    | Event Version. Value: `1`                                                |
| `height`                   | *int64*  | Height of the block containing the transaction                           |
| `uuid`                     | *string* | Unique ID that is assigned on event creation                             |

*Example* : T.B.D  

## event::MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_FAILED
*Name* : MsgSubmitCommunityPoolSpendProposalFailed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                        | Type     | Description                                                             |
| -------------------------- | -------- | ----------------------------------------------------------------------- |
| `proposalId`               | *string* | *(Optional)* Proposal ID                                                |
| `proposerAddress`          | *string* | Proposer blockchain address                                             |
| `initialDeposit`           | *string* | Initially deposited CRO Amount in base unit                             |
| `content`                  | *object* | Content wrapper                                                         |
| `content.@type`            | *string* | Cosmos SDK type URL                                                     |
| `content.title`            | *string* | Content title indicating the change                                     |
| `content.description`      | *string* | Action description                                                      |
| `content.recipientAddress` | *string* | Recipient blockchain address                                            |
| `content.amount`           | *string* | Recipient CRO Amount in base unit                                       |
| `msgName`                  | *string* | Blockchain Message type . Value: `MsgSubmitCommunityPoolSpendProposal`  |
| `txHash`                   | *string* | TxID of the blockchain transaction containing the event                 |
| `msgIndex`                 | *int*    | message index on the block                                              |
| `name`                     | *string* | Specific Event Name. Value: `MsgSubmitCommunityPoolSpendProposalFailed` |
| `version`                  | *int*    | Event Version. Value: `1`                                               |
| `height`                   | *int64*  | Height of the block containing the transaction                          |
| `uuid`                     | *string* | Unique ID that is assigned on event creation                            |

*Example* : T.B.D  