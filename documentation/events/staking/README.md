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

## event::MSG_CREATE_VALIDATOR_FAILED
*Name* : MsgCreateValidatorFailed

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
| `name`                          | *string* | Specific Event Name. Value: `MsgCreateValidatorFailed`  |
| `version`                       | *int*    | Event Version. Value: `1`                               |
| `height`                        | *int64*  | Height of the block containing the transaction          |
| `uuid`                          | *string* | Unique ID that is assigned on event creation            |

*Example* :

```json
{
    "name": "MsgCreateValidatorFailed",
    "uuid": "2c1e1e6b-30e8-403e-ab6c-d922e800a318",
    "amount": "50000000000000",
    "height": 80202,
    "pubkey": "tcrocnclconspub1zcjduepqjl7jdvpa0g3pqczmh8cks9669slpkusa2vlnk7rsqehhhe6du02qvxpzca",
    "txHash": "C0B72C26FBE0A191E9AD534993B2B964BB9BD5FDA27E0EC237457AA70D515D05",
    "msgName": "MsgCreateValidator",
    "version": 1,
    "msgIndex": 0,
    "description": {
        "details": "",
        "moniker": "Fierydev-MCLB",
        "website": "",
        "identity": "",
        "securityContact": "fierydev@gmail.com"
    },
    "commissionRates": {
        "rate": "0.100000000000000000",
        "maxRate": "0.200000000000000000",
        "maxChangeRate": "0.010000000000000000"
    },
    "delegatorAddress": "tcro16p0um8f20sq77xqlpqqmx4t5qwyczgjm85euth",
    "validatorAddress": "tcrocncl16p0um8f20sq77xqlpqqmx4t5qwyczgjmjt69n5",
    "minSelfDelegation": "1"
}
```  

## event::MSG_EDIT_VALIDATOR_CREATED
*Name* : MsgEditValidatorCreated

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                           | Type     | Description                                                                |
| ----------------------------- | -------- | -------------------------------------------------------------------------- |
| `validatorAddress`            | *string* | Validator address                                                          |
| `minSelfDelegation`           | *string* | *(Optional)* New minimum Self Delegation amount.                           |
| `commissionRate`              | *string* | *(Optional)* New commission rate                                           |
| `description`                 | *object* | Validator description wrapper                                              |
| `description.moniker`         | *string* | *(Optional)* New validator node public moniker. Default: `[do-not-modify]` |
| `description.identity`        | *string* | *(Optional)* New validator node identity. Default: `[do-not-modify]`       |
| `description.website`         | *string* | *(Optional)* New website domain name. Default: `[do-not-modify]`           |
| `description.securityContact` | *string* | *(Optional)* New Validator node email contact. Default: `[do-not-modify]`  |
| `description.details`         | *string* | *(Optional)* New Validator node details. Default: `[do-not-modify]`        |
| `msgName`                     | *string* | Blockchain Message type . Value: `MsgEditValidator`                        |
| `txHash`                      | *string* | TxID of the blockchain transaction containing the event                    |
| `msgIndex`                    | *int*    | message index on the block                                                 |
| `name`                        | *string* | Specific Event Name. Value: `MsgEditValidatorCreated`                      |
| `version`                     | *int*    | Event Version. Value: `1`                                                  |
| `height`                      | *int64*  | Height of the block containing the transaction                             |
| `uuid`                        | *string* | Unique ID that is assigned on event creation                               |

*Example* :

```json
{
    "name": "MsgEditValidatorCreated",
    "uuid": "7ea5034f-0ccf-45ea-be46-9f2f82caa9a2",
    "height": 82066,
    "txHash": "8D0AA51BA31B2BCF3A8D7658B421476E614A24B9D8E5A9BD0BC23A1D2F9B5FAD",
    "msgName": "MsgEditValidator",
    "version": 1,
    "msgIndex": 0,
    "description": {
        "details": "[do-not-modify]",
        "moniker": "1NP_SirJN",
        "website": "[do-not-modify]",
        "identity": "[do-not-modify]",
        "securityContact": "[do-not-modify]"
    },
    "commissionRate": null,
    "validatorAddress": "tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu",
    "minSelfDelegation": null
}
```  