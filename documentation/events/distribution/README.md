# Distribution Module Event List

-   [event::MSG_SET_WITHDRAW_ADDRESS_CREATED](#event_msg_set_withdraw_address_created)
-   [event::MSG_SET_WITHDRAW_ADDRESS_FAILED](#event_msg_set_withdraw_address_failed)
-   [event::MSG_WITHDRAW_DELEGATOR_REWARD_CREATED](#event_msg_withdraw_delegator_reward_created)
-   [event::MSG_WITHDRAW_DELEGATOR_REWARD_FAILED](#event_msg_withdraw_delegator_reward_failed)
-   [event::MSG_WITHDRAW_VALIDATOR_COMMISSION_CREATED](#event_msg_withdraw_validator_commission_created)
-   [event::MSG_WITHDRAW_VALIDATOR_COMMISSION_FAILED](#event_msg_withdraw_validator_commission_failed)
-   [event::MSG_FUND_COMMUNITY_POOL_CREATED](#event_msg_fund_community_pool_created)
-   [event::MSG_FUND_COMMUNITY_POOL_FAILED](#event_msg_fund_community_pool_failed)

## event::MSG_SET_WITHDRAW_ADDRESS_CREATED

_Name_ : /cosmos.distribution.v1beta1.MsgSetWithdrawAddress.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                | Type     | Description                                                                                                |
| ------------------ | -------- | ---------------------------------------------------------------------------------------------------------- |
| `delegatorAddress` | _string_ | Delegator address. More [here](https://chain.crypto.com/docs/chain-details/module_overview.html#delegator) |
| `withdrawAddress`  | _string_ | Intended withdrawal address                                                                                |
| `msgName`          | _string_ | Blockchain Message type . Value: `MsgSetWithdrawAddress`                                                   |
| `txHash`           | _string_ | TxID of the blockchain transaction containing the event                                                    |
| `msgIndex`         | _int_    | message index on the block                                                                                 |
| `name`             | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.MsgSetWithdrawAddress.Created`                   |
| `version`          | _int_    | Event Version. Value: `1`                                                                                  |
| `height`           | _int64_  | Height of the block containing the transaction                                                             |
| `uuid`             | _string_ | Unique ID that is assigned on event creation                                                               |

_Example_ : T.B.D

## event::MSG_SET_WITHDRAW_ADDRESS_FAILED

_Name_ : /cosmos.distribution.v1beta1.MsgSetWithdrawAddress.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                | Type     | Description                                                                                                |
| ------------------ | -------- | ---------------------------------------------------------------------------------------------------------- |
| `delegatorAddress` | _string_ | Delegator address. More [here](https://chain.crypto.com/docs/chain-details/module_overview.html#delegator) |
| `withdrawAddress`  | _string_ | Intended withdrawal address                                                                                |
| `msgName`          | _string_ | Blockchain Message type . Value: `MsgSetWithdrawAddress`                                                   |
| `txHash`           | _string_ | TxID of the blockchain transaction containing the event                                                    |
| `msgIndex`         | _int_    | message index on the block                                                                                 |
| `name`             | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.MsgSetWithdrawAddress.Failed`                    |
| `version`          | _int_    | Event Version. Value: `1`                                                                                  |
| `height`           | _int64_  | Height of the block containing the transaction                                                             |
| `uuid`             | _string_ | Unique ID that is assigned on event creation                                                               |

_Example_ : T.B.D

## event::MSG_WITHDRAW_DELEGATOR_REWARD_CREATED

_Name_ : /cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                | Type     | Description                                                                                                |
| ------------------ | -------- | ---------------------------------------------------------------------------------------------------------- |
| `delegatorAddress` | _string_ | Delegator address. More [here](https://chain.crypto.com/docs/chain-details/module_overview.html#delegator) |
| `recipientAddress` | _string_ | Recipient blockchain address                                                                               |
| `validatorAddress` | _string_ | Validator address                                                                                          |
| `amount`           | _bigint_ | CRO amount in base unit                                                                                    |
| `msgName`          | _string_ | Blockchain Message type . Value: `MsgWithdrawDelegatorReward`                                              |
| `txHash`           | _string_ | TxID of the blockchain transaction containing the event                                                    |
| `msgIndex`         | _int_    | message index on the block                                                                                 |
| `name`             | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward.Created`              |
| `version`          | _int_    | Event Version. Value: `1`                                                                                  |
| `height`           | _int64_  | Height of the block containing the transaction                                                             |
| `uuid`             | _string_ | Unique ID that is assigned on event creation                                                               |

_Example_ :

```json
{
    "name": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward.Created",
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

_Name_ : /cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                | Type     | Description                                                                                                |
| ------------------ | -------- | ---------------------------------------------------------------------------------------------------------- |
| `delegatorAddress` | _string_ | Delegator address. More [here](https://chain.crypto.com/docs/chain-details/module_overview.html#delegator) |
| `recipientAddress` | _string_ | Recipient blockchain address                                                                               |
| `validatorAddress` | _string_ | Validator address                                                                                          |
| `amount`           | _bigint_ | CRO amount in base unit. Value : `0`                                                                       |
| `msgName`          | _string_ | Blockchain Message type . Value: `MsgWithdrawDelegatorReward`                                              |
| `txHash`           | _string_ | TxID of the blockchain transaction containing the event                                                    |
| `msgIndex`         | _int_    | message index on the block                                                                                 |
| `name`             | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward.Failed`               |
| `version`          | _int_    | Event Version. Value: `1`                                                                                  |
| `height`           | _int64_  | Height of the block containing the transaction                                                             |
| `uuid`             | _string_ | Unique ID that is assigned on event creation                                                               |

_Example_ :

```json
{
    "name": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward.Failed",
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

_Name_ : /cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                | Type     | Description                                                                                       |
| ------------------ | -------- | ------------------------------------------------------------------------------------------------- |
| `recipientAddress` | _string_ | Recipient blockchain address                                                                      |
| `validatorAddress` | _string_ | Validator address                                                                                 |
| `amount`           | _bigint_ | CRO amount in base unit                                                                           |
| `msgName`          | _string_ | Blockchain Message type . Value: `MsgWithdrawValidatorCommission`                                 |
| `txHash`           | _string_ | TxID of the blockchain transaction containing the event                                           |
| `msgIndex`         | _int_    | message index on the block                                                                        |
| `name`             | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission.Created` |
| `version`          | _int_    | Event Version. Value: `1`                                                                         |
| `height`           | _int64_  | Height of the block containing the transaction                                                    |
| `uuid`             | _string_ | Unique ID that is assigned on event creation                                                      |

_Example_ :

```json
{
    "name": "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission.Created",
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

_Name_ : /cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                | Type     | Description                                                                                      |
| ------------------ | -------- | ------------------------------------------------------------------------------------------------ |
| `recipientAddress` | _string_ | Recipient blockchain address                                                                     |
| `validatorAddress` | _string_ | Validator address                                                                                |
| `amount`           | _bigint_ | CRO amount in base unit. Value : `0`                                                             |
| `msgName`          | _string_ | Blockchain Message type . Value: `MsgWithdrawValidatorCommission`                                |
| `txHash`           | _string_ | TxID of the blockchain transaction containing the event                                          |
| `msgIndex`         | _int_    | message index on the block                                                                       |
| `name`             | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission.Failed` |
| `version`          | _int_    | Event Version. Value: `1`                                                                        |
| `height`           | _int64_  | Height of the block containing the transaction                                                   |
| `uuid`             | _string_ | Unique ID that is assigned on event creation                                                     |

_Example_ :

```json
{
    "name": "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission.Failed",
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

_Name_ : /cosmos.distribution.v1beta1.MsgFundCommunityPool.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key         | Type     | Description                                                                             |
| ----------- | -------- | --------------------------------------------------------------------------------------- |
| `depositor` | _string_ | Recipient blockchain address                                                            |
| `amount`    | _bigint_ | CRO amount in base unit                                                                 |
| `msgName`   | _string_ | Blockchain Message type. Value: `MsgFundCommunityPool`                                  |
| `txHash`    | _string_ | TxID of the blockchain transaction containing the event                                 |
| `msgIndex`  | _int_    | message index on the block                                                              |
| `name`      | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.MsgFundCommunityPool.Created` |
| `version`   | _int_    | Event Version. Value: `1`                                                               |
| `height`    | _int64_  | Height of the block containing the transaction                                          |
| `uuid`      | _string_ | Unique ID that is assigned on event creation                                            |

_Example_ :

```json
{
    "name": "/cosmos.distribution.v1beta1.MsgFundCommunityPool.Created",
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

_Name_ : /cosmos.distribution.v1beta1.MsgFundCommunityPool.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key         | Type     | Description                                                                            |
| ----------- | -------- | -------------------------------------------------------------------------------------- |
| `depositor` | _string_ | Recipient blockchain address                                                           |
| `amount`    | _bigint_ | CRO amount in base unit                                                                |
| `msgName`   | _string_ | Blockchain Message type. Value: `MsgFundCommunityPool`                                 |
| `txHash`    | _string_ | TxID of the blockchain transaction containing the event                                |
| `msgIndex`  | _int_    | message index on the block                                                             |
| `name`      | _string_ | Specific Event Name. Value: `/cosmos.distribution.v1beta1.MsgFundCommunityPool.Failed` |
| `version`   | _int_    | Event Version. Value: `1`                                                              |
| `height`    | _int64_  | Height of the block containing the transaction                                         |
| `uuid`      | _string_ | Unique ID that is assigned on event creation                                           |

_Example_ :

```json
{
    "name": "/cosmos.distribution.v1beta1.MsgFundCommunityPool.Failed",
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
