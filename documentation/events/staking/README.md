# Staking Module Event List
  - [event::MSG_CREATE_VALIDATOR_CREATED](#event_msg_create_validator_created)
  - [event::MSG_CREATE_VALIDATOR_FAILED](#event_msg_create_validator_failed)
  - [event::MSG_EDIT_VALIDATOR_CREATED](#event_msg_edit_validator_created)
  - [event::MSG_EDIT_VALIDATOR_FAILED](#event_msg_edit_validator_failed)
  - [event::MSG_DELEGATE_CREATED](#event_msg_delegate_created)
  - [event::MSG_DELEGATE_FAILED](#event_msg_delegate_failed)
  - [event::MSG_UNDELEGATE_CREATED](#event_msg_undelegate_created)
  - [event::MSG_UNDELEGATE_FAILED](#event_msg_undelegate_failed)
  - [event::MSG_BEGIN_REDELEGATE_CREATED](#event_msg_begin_redelegate_created)
  - [event::MSG_BEGIN_REDELEGATE_FAILED](#event_msg_begin_redelegate_failed)


## event::MSG_CREATE_VALIDATOR_CREATED
*Name* : MsgCreateValidatorCreated

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                             | Type     | Description                                             |
| ------------------------------- | -------- | ------------------------------------------------------- |
| `amount`                        | *bigint* | CRO amount in basic unit                                |
| `pubkey`                        | *string* | Consensus public key                                    |
| `validatorAddress`              | *string* | Validator address                                       |
| `delegatorAddress`              | *string* | Delegator address                                       |
| `minSelfDelegation`             | *string* | Minimum Self Delegation amount                          |
| `description`                   | *object* | Validator description wrapper                           |
| `description.moniker`           | *string* | Validator node public moniker                           |
| `description.identity`          | *string* | *(Optional)* Validator node identity                    |
| `description.website`           | *string* | *(Optional)* Website domain name                        |
| `description.securityContact`   | *string* | *(Optional)* Validator node email contact               |
| `description.details`           | *string* | *(Optional)* Validator node details                     |
| `commissionRates`               | *object* | Validator node commission wrapper                       |
| `commissionRates.rate`          | *string* | Validator commission rate                               |
| `commissionRates.maxRate`       | *string* | Validator commission maximum rate                       |
| `commissionRates.maxChangeRate` | *string* | Validator commission maximum change rate                |
| `msgName`                       | *string* | Blockchain Message type . Value: `MsgCreateValidator`   |
| `txHash`                        | *string* | TxID of the blockchain transaction containing the event |
| `msgIndex`                      | *int*    | message index on the block                              |
| `name`                          | *string* | Specific Event Name. Value: `MsgCreateValidatorCreated` |
| `version`                       | *int*    | Event Version. Value: `1`                               |
| `height`                        | *int64*  | Height of the block containing the transaction          |
| `uuid`                          | *string* | Unique ID that is assigned on event creation            |

*Example* :

```json
{
    "name": "MsgCreateValidatorCreated",
    "uuid": "ee0920d1-3d34-4371-b7e3-dab9aa81273d",
    "amount": "100000000",
    "height": 174508,
    "pubkey": "tcrocnclconspub1zcjduepqwz563kykhvntl7epx2p0ewejlmprxpyq68wf3v5hy7ja2vu3zd2qg65v20",
    "txHash": "3C5D9DA6396EAF68DAC96D74A1CF59DE5D6E8A8C880E6B561E09A6C642B29ECD",
    "msgName": "MsgCreateValidator",
    "version": 1,
    "msgIndex": 0,
    "description": {
        "details": "",
        "moniker": "jerome-node",
        "website": "",
        "identity": "",
        "securityContact": "Jerome@crypto.bzh"
    },
    "commissionRates": {
        "rate": "0.100000000000000000",
        "maxRate": "0.200000000000000000",
        "maxChangeRate": "0.010000000000000000"
    },
    "delegatorAddress": "tcro1f0q0k4yysavkxs75a83w70384dqu7vnxnfw7jy",
    "validatorAddress": "tcrocncl1f0q0k4yysavkxs75a83w70384dqu7vnxxkd828",
    "minSelfDelegation": "1"
}
```  