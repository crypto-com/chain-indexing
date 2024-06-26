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
*Name* : /cosmos.staking.v1beta1.MsgCreateValidator.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                             | Type     | Description                                                                      |
| ------------------------------- | -------- | -------------------------------------------------------------------------------- |
| `amount`                        | *bigint* | CRO amount in basic unit                                                         |
| `pubkey`                        | *string* | Consensus public key                                                             |
| `validatorAddress`              | *string* | Validator address                                                                |
| `delegatorAddress`              | *string* | Delegator address                                                                |
| `minSelfDelegation`             | *string* | Minimum Self Delegation amount                                                   |
| `description`                   | *object* | Validator description wrapper                                                    |
| `description.moniker`           | *string* | Validator node public moniker                                                    |
| `description.identity`          | *string* | _(Optional)_ Validator node identity                                             |
| `description.website`           | *string* | _(Optional)_ Website domain name                                                 |
| `description.securityContact`   | *string* | _(Optional)_ Validator node email contact                                        |
| `description.details`           | *string* | _(Optional)_ Validator node details                                              |
| `commissionRates`               | *object* | Validator node commission wrapper                                                |
| `commissionRates.rate`          | *string* | Validator commission rate                                                        |
| `commissionRates.maxRate`       | *string* | Validator commission maximum rate                                                |
| `commissionRates.maxChangeRate` | *string* | Validator commission maximum change rate                                         |
| `msgName`                       | *string* | Blockchain Message type . Value: `MsgCreateValidator`                            |
| `txHash`                        | *string* | TxID of the blockchain transaction containing the event                          |
| `msgIndex`                      | *int*    | message index on the block                                                       |
| `name`                          | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgCreateValidator.Created` |
| `version`                       | *int*    | Event Version. Value: `1`                                                        |
| `height`                        | *int64*  | Height of the block containing the transaction                                   |
| `uuid`                          | *string* | Unique ID that is assigned on event creation                                     |

*Example* :

```json
{
    "name": "/cosmos.staking.v1beta1.MsgCreateValidator.Created",
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
        "moniker": "jer-node",
        "website": "",
        "identity": "",
        "securityContact": "Jme@cry.bzh"
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
*Name* : /cosmos.staking.v1beta1.MsgCreateValidator.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                             | Type     | Description                                                                     |
| ------------------------------- | -------- | ------------------------------------------------------------------------------- |
| `amount`                        | *bigint* | CRO amount in basic unit                                                        |
| `pubkey`                        | *string* | Consensus public key                                                            |
| `validatorAddress`              | *string* | Validator address                                                               |
| `delegatorAddress`              | *string* | Delegator address                                                               |
| `minSelfDelegation`             | *string* | Minimum Self Delegation amount                                                  |
| `description`                   | *object* | Validator description wrapper                                                   |
| `description.moniker`           | *string* | Validator node public moniker                                                   |
| `description.identity`          | *string* | _(Optional)_ Validator node identity                                            |
| `description.website`           | *string* | _(Optional)_ Website domain name                                                |
| `description.securityContact`   | *string* | _(Optional)_ Validator node email contact                                       |
| `description.details`           | *string* | _(Optional)_ Validator node details                                             |
| `commissionRates`               | *object* | Validator node commission wrapper                                               |
| `commissionRates.rate`          | *string* | Validator commission rate                                                       |
| `commissionRates.maxRate`       | *string* | Validator commission maximum rate                                               |
| `commissionRates.maxChangeRate` | *string* | Validator commission maximum change rate                                        |
| `msgName`                       | *string* | Blockchain Message type . Value: `MsgCreateValidator`                           |
| `txHash`                        | *string* | TxID of the blockchain transaction containing the event                         |
| `msgIndex`                      | *int*    | message index on the block                                                      |
| `name`                          | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgCreateValidator.Failed` |
| `version`                       | *int*    | Event Version. Value: `1`                                                       |
| `height`                        | *int64*  | Height of the block containing the transaction                                  |
| `uuid`                          | *string* | Unique ID that is assigned on event creation                                    |

*Example* :

```json
{
    "name": "/cosmos.staking.v1beta1.MsgCreateValidator.Failed",
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
        "moniker": "FasdfydeCLB",
        "website": "",
        "identity": "",
        "securityContact": "fdsryddev@gxmail.com"
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
*Name* : /cosmos.staking.v1beta1.MsgEditValidator.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                           | Type     | Description                                                                    |
| ----------------------------- | -------- | ------------------------------------------------------------------------------ |
| `validatorAddress`            | *string* | Validator address                                                              |
| `minSelfDelegation`           | *string* | _(Optional)_ New minimum Self Delegation amount.                               |
| `commissionRate`              | *string* | _(Optional)_ New commission rate                                               |
| `description`                 | *object* | Validator description wrapper                                                  |
| `description.moniker`         | *string* | _(Optional)_ New validator node public moniker. Default: `[do-not-modify]`     |
| `description.identity`        | *string* | _(Optional)_ New validator node identity. Default: `[do-not-modify]`           |
| `description.website`         | *string* | _(Optional)_ New website domain name. Default: `[do-not-modify]`               |
| `description.securityContact` | *string* | _(Optional)_ New Validator node email contact. Default: `[do-not-modify]`      |
| `description.details`         | *string* | _(Optional)_ New Validator node details. Default: `[do-not-modify]`            |
| `msgName`                     | *string* | Blockchain Message type . Value: `MsgEditValidator`                            |
| `txHash`                      | *string* | TxID of the blockchain transaction containing the event                        |
| `msgIndex`                    | *int*    | message index on the block                                                     |
| `name`                        | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgEditValidator.Created` |
| `version`                     | *int*    | Event Version. Value: `1`                                                      |
| `height`                      | *int64*  | Height of the block containing the transaction                                 |
| `uuid`                        | *string* | Unique ID that is assigned on event creation                                   |

*Example* :

```json
{
    "name": "/cosmos.staking.v1beta1.MsgEditValidator.Created",
    "uuid": "7ea5034f-0ccf-45ea-be46-9f2f82caa9a2",
    "height": 82066,
    "txHash": "8D0AA51BA31B2BCF3A8D7658B421476E614A24B9D8E5A9BD0BC23A1D2F9B5FAD",
    "msgName": "MsgEditValidator",
    "version": 1,
    "msgIndex": 0,
    "description": {
        "details": "[do-not-modify]",
        "moniker": "1asP_asJN",
        "website": "[do-not-modify]",
        "identity": "[do-not-modify]",
        "securityContact": "[do-not-modify]"
    },
    "commissionRate": null,
    "validatorAddress": "tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu",
    "minSelfDelegation": null
}
```  

## event::MSG_EDIT_VALIDATOR_FAILED
*Name* : /cosmos.staking.v1beta1.MsgEditValidator.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                           | Type     | Description                                                                   |
| ----------------------------- | -------- | ----------------------------------------------------------------------------- |
| `validatorAddress`            | *string* | Validator address                                                             |
| `minSelfDelegation`           | *string* | _(Optional)_ New minimum Self Delegation amount.                              |
| `commissionRate`              | *string* | _(Optional)_ New commission rate                                              |
| `description`                 | *object* | Validator description wrapper                                                 |
| `description.moniker`         | *string* | _(Optional)_ New validator node public moniker. Default: `[do-not-modify]`    |
| `description.identity`        | *string* | _(Optional)_ New validator node identity. Default: `[do-not-modify]`          |
| `description.website`         | *string* | _(Optional)_ New website domain name. Default: `[do-not-modify]`              |
| `description.securityContact` | *string* | _(Optional)_ New Validator node email contact. Default: `[do-not-modify]`     |
| `description.details`         | *string* | _(Optional)_ New Validator node details. Default: `[do-not-modify]`           |
| `msgName`                     | *string* | Blockchain Message type . Value: `MsgEditValidator`                           |
| `txHash`                      | *string* | TxID of the blockchain transaction containing the event                       |
| `msgIndex`                    | *int*    | message index on the block                                                    |
| `name`                        | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgEditValidator.Failed` |
| `version`                     | *int*    | Event Version. Value: `1`                                                     |
| `height`                      | *int64*  | Height of the block containing the transaction                                |
| `uuid`                        | *string* | Unique ID that is assigned on event creation                                  |

*Example* :

```json
{
    "name": "/cosmos.staking.v1beta1.MsgEditValidator.Failed",
    "uuid": "070b919c-e105-4c60-b524-24ec0f3f19d9",
    "height": 88293,
    "txHash": "296E492F7F35A1FBD4BF1C8C34E30BF07417F92E17A61279031DE650A1B98E91",
    "msgName": "MsgEditValidator",
    "version": 1,
    "msgIndex": 0,
    "description": {
        "details": "[do-not-modify]",
        "moniker": "[do-not-modify]",
        "website": "[do-not-modify]",
        "identity": "random0290Frandom",
        "securityContact": "[do-not-modify]"
    },
    "commissionRate": "0.150000000000000000",
    "validatorAddress": "tcrocncl122w9fhc0pu3ey9r6hekznd2fkl5jswl5aqsvgy",
    "minSelfDelegation": null
}
```  

## event::MSG_DELEGATE_CREATED
*Name* : /cosmos.staking.v1beta1.MsgDelegate.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                               |
| ------------------ | -------- | ------------------------------------------------------------------------- |
| `delegatorAddress` | *string* | Delegator address                                                         |
| `validatorAddress` | *string* | Validator address                                                         |
| `amount`           | *bigint* | CRO Amount in base unit                                                   |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgDelegate`                            |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event                   |
| `msgIndex`         | *int*    | message index on the block                                                |
| `name`             | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgDelegate.Created` |
| `version`          | *int*    | Event Version. Value: `1`                                                 |
| `height`           | *int64*  | Height of the block containing the transaction                            |
| `uuid`             | *string* | Unique ID that is assigned on event creation                              |

*Example* :

```json
{
    "name": "/cosmos.staking.v1beta1.MsgDelegate.Created",
    "uuid": "c51d9e37-3c8c-46eb-b2de-9e79137dfe8a",
    "amount": "24077737755",
    "height": 80190,
    "txHash": "DD148B9B34A215B3FD15097579E4F45237C0DBD0A7983A043F1A3BA436F1590A",
    "msgName": "MsgDelegate",
    "version": 1,
    "msgIndex": 0,
    "delegatorAddress": "tcro152ena75gh5nqnu2nlarwmpzxa2czxs8y9e3slh",
    "validatorAddress": "tcrocncl152ena75gh5nqnu2nlarwmpzxa2czxs8ysxjf85"
}
```  

## event::MSG_DELEGATE_FAILED
*Name* : /cosmos.staking.v1beta1.MsgDelegate.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                              |
| ------------------ | -------- | ------------------------------------------------------------------------ |
| `delegatorAddress` | *string* | Delegator address                                                        |
| `validatorAddress` | *string* | Validator address                                                        |
| `amount`           | *bigint* | CRO Amount in base unit                                                  |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgDelegate`                           |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event                  |
| `msgIndex`         | *int*    | message index on the block                                               |
| `name`             | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgDelegate.Failed` |
| `version`          | *int*    | Event Version. Value: `1`                                                |
| `height`           | *int64*  | Height of the block containing the transaction                           |
| `uuid`             | *string* | Unique ID that is assigned on event creation                             |

*Example* :

```json
{
    "name": "/cosmos.staking.v1beta1.MsgDelegate.Failed",
    "uuid": "e6324f2f-ce4b-41d5-b91d-cf46c7ea31ba",
    "amount": "29700000000",
    "height": 84771,
    "txHash": "701D8B249B10AF9F0EE7E991262F296C9C9B747424CC58C068A693D019EB8FD1",
    "msgName": "MsgDelegate",
    "version": 1,
    "msgIndex": 0,
    "delegatorAddress": "tcro1urmrrmmt6gdf077dmgt95cmj6tc0z9045d5xmw",
    "validatorAddress": "tcrocncl1urmrrmmt6gdf077dmgt95cmj6tc0z904pjhlrd"
}
```  

## event::MSG_UNDELEGATE_CREATED
*Name* : /cosmos.staking.v1beta1.MsgUndelegate.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                                  |
| ------------------ | -------- | ---------------------------------------------------------------------------- |
| `delegatorAddress` | *string* | Delegator address                                                            |
| `validatorAddress` | *string* | Validator address                                                            |
| `amount`           | *bigint* | CRO Amount in base unit                                                      |
| `unbondCompleteAt` | *string* | _(Optional)_ Unbonding completion timestamp. Golang type: `*utctime.UTCTime` |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgUndelegate`                             |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event                      |
| `msgIndex`         | *int*    | message index on the block                                                   |
| `name`             | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgUndelegate.Created`  |
| `version`          | *int*    | Event Version. Value: `1`                                                    |
| `height`           | *int64*  | Height of the block containing the transaction                               |
| `uuid`             | *string* | Unique ID that is assigned on event creation                                 |

*Example* :

```json
{
    "name": "/cosmos.staking.v1beta1.MsgUndelegate.Created",
    "uuid": "13e87da9-2ca9-4510-a1b3-f99decefb718",
    "amount": "50000000000000",
    "height": 192744,
    "txHash": "D6B15F498A1E31452738F43AF124FAE7007EB3344E66A5423941667B46EA41EE",
    "msgName": "MsgUndelegate",
    "version": 1,
    "msgIndex": 0,
    "delegatorAddress": "tcro15e69kdrtczajjdlzyt2qgs5q2anc5qpmr4mrlp",
    "unbondCompleteAt": "2020-10-29T03:40:20Z",
    "validatorAddress": "tcrocncl15e69kdrtczajjdlzyt2qgs5q2anc5qpmk2c68z"
}
```  

## event::MSG_UNDELEGATE_FAILED
*Name* : /cosmos.staking.v1beta1.MsgUndelegate.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                                  |
| ------------------ | -------- | ---------------------------------------------------------------------------- |
| `delegatorAddress` | *string* | Delegator address                                                            |
| `validatorAddress` | *string* | Validator address                                                            |
| `amount`           | *bigint* | CRO Amount in base unit                                                      |
| `unbondCompleteAt` | *string* | _(Optional)_ Unbonding completion timestamp. Golang type: `*utctime.UTCTime` |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgUndelegate`                             |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event                      |
| `msgIndex`         | *int*    | message index on the block                                                   |
| `name`             | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgUndelegate.Failed`   |
| `version`          | *int*    | Event Version. Value: `1`                                                    |
| `height`           | *int64*  | Height of the block containing the transaction                               |
| `uuid`             | *string* | Unique ID that is assigned on event creation                                 |

*Example* :

```json
{
    "name": "/cosmos.staking.v1beta1.MsgUndelegate.Failed",
    "uuid": "d5ddb893-2ce7-4fbd-b85f-e8749dbe8d2e",
    "amount": "50000000000000",
    "height": 185255,
    "txHash": "257D324CE4ED85357126FA82B9CD31CDE301BA5109945D27968A8A5D87C7B6BD",
    "msgName": "MsgUndelegate",
    "version": 1,
    "msgIndex": 0,
    "delegatorAddress": "tcro15e69kdrtczajjdlzyt2qgs5q2anc5qpmr4mrlp",
    "unbondCompleteAt": null,
    "validatorAddress": "tcrocncl1llst0cguh5azl9t8wr6mz5yzjuwukz7f0pp83e"
}
```  

## event::MSG_BEGIN_REDELEGATE_CREATED
*Name* : /cosmos.staking.v1beta1.MsgBeginRedelegate.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                   | Type     | Description                                                                      |
| --------------------- | -------- | -------------------------------------------------------------------------------- |
| `delegatorAddress`    | *string* | Delegator address                                                                |
| `validatorSrcAddress` | *string* | Source Validator address                                                         |
| `validatorDstAddress` | *string* | Destination Validator address                                                    |
| `amount`              | *bigint* | CRO Amount in base unit                                                          |
| `msgName`             | *string* | Blockchain Message type . Value: `MsgBeginRedelegate`                            |
| `txHash`              | *string* | TxID of the blockchain transaction containing the event                          |
| `msgIndex`            | *int*    | message index on the block                                                       |
| `name`                | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgBeginRedelegate.Created` |
| `version`             | *int*    | Event Version. Value: `1`                                                        |
| `height`              | *int64*  | Height of the block containing the transaction                                   |
| `uuid`                | *string* | Unique ID that is assigned on event creation                                     |

*Example* : T.B.D  

## event::MSG_BEGIN_REDELEGATE_FAILED
*Name* : /cosmos.staking.v1beta1.MsgBeginRedelegate.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                   | Type     | Description                                                                     |
| --------------------- | -------- | ------------------------------------------------------------------------------- |
| `delegatorAddress`    | *string* | Delegator address                                                               |
| `validatorSrcAddress` | *string* | Source Validator address                                                        |
| `validatorDstAddress` | *string* | Destination Validator address                                                   |
| `amount`              | *bigint* | CRO Amount in base unit                                                         |
| `msgName`             | *string* | Blockchain Message type . Value: `MsgBeginRedelegate`                           |
| `txHash`              | *string* | TxID of the blockchain transaction containing the event                         |
| `msgIndex`            | *int*    | message index on the block                                                      |
| `name`                | *string* | Specific Event Name. Value: `/cosmos.staking.v1beta1.MsgBeginRedelegate.Failed` |
| `version`             | *int*    | Event Version. Value: `1`                                                       |
| `height`              | *int64*  | Height of the block containing the transaction                                  |
| `uuid`                | *string* | Unique ID that is assigned on event creation                                    |

*Example* : T.B.D  

## event::BONDING_COMPLETED
*Name* : BondingCompleted

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key         | Type     | Description                                    |
| ----------- | -------- | ---------------------------------------------- |
| `delegator` | *string* | Delegator address                              |
| `validator` | *string* | Validator address                              |
| `amount`    | *bigint* | CRO Amount in base unit                        |
| `name`      | *string* | Specific Event Name. Value: `BondingCompleted` |
| `version`   | *int*    | Event Version. Value: `1`                      |
| `height`    | *int64*  | Height of the block containing the transaction |
| `uuid`      | *string* | Unique ID that is assigned on event creation   |

*Example* : T.B.D  
