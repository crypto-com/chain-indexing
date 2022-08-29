# Slashing Module Event List

-   [event::MSG_UNJAIL_CREATED](#event_MSG_UNJAIL_CREATED)
-   [event::MSG_UNJAIL_FAILED](#event_MSG_UNJAIL_FAILED)

## event::MSG_UNJAIL_CREATED

_Name_ : /cosmos.slashing.v1beta1.MsgUnjail.Created

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                | Type     | Description                                                              |
| ------------------ | -------- | ------------------------------------------------------------------------ |
| `validatorAddress` | _string_ | Validator address                                                        |
| `msgName`          | _string_ | Blockchain Message type . Value: `MsgUnjail`                             |
| `txHash`           | _string_ | TxID of the blockchain transaction containing the event                  |
| `msgIndex`         | _int_    | message index on the block                                               |
| `name`             | _string_ | Specific Event Name. Value: `/cosmos.slashing.v1beta1.MsgUnjail.Created` |
| `version`          | _int_    | Event Version. Value: `1`                                                |
| `height`           | _int64_  | Height of the block containing the transaction                           |
| `uuid`             | _string_ | Unique ID that is assigned on event creation                             |

_Example_ :

```json
{
    "name": "/cosmos.slashing.v1beta1.MsgUnjail.Created",
    "uuid": "102121bc-178a-465e-8b91-fb35b2e99a0b",
    "height": 29161,
    "txHash": "110870C6DE486220B6A89E54B3B4AADFED5FC56D547E043C73A47F657BC328BC",
    "msgName": "MsgUnjail",
    "version": 1,
    "msgIndex": 0,
    "validatorAddress": "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5"
}
```

## event::MSG_UNJAIL_FAILED

_Name_ : /cosmos.slashing.v1beta1.MsgUnjail.Failed

_Type_ : [MsgBase](../README.md#MsgBase)

_Structure_ :

| Key                | Type     | Description                                                             |
| ------------------ | -------- | ----------------------------------------------------------------------- |
| `validatorAddress` | _string_ | Validator address                                                       |
| `msgName`          | _string_ | Blockchain Message type . Value: `MsgUnjail`                            |
| `txHash`           | _string_ | TxID of the blockchain transaction containing the event                 |
| `msgIndex`         | _int_    | message index on the block                                              |
| `name`             | _string_ | Specific Event Name. Value: `/cosmos.slashing.v1beta1.MsgUnjail.Failed` |
| `version`          | _int_    | Event Version. Value: `1`                                               |
| `height`           | _int64_  | Height of the block containing the transaction                          |
| `uuid`             | _string_ | Unique ID that is assigned on event creation                            |

_Example_ :

```json
{
    "name": "/cosmos.slashing.v1beta1.MsgUnjail.Failed",
    "uuid": "a4a03f80-4545-4c01-a9a0-2a017bb4e0d2",
    "height": 11977,
    "txHash": "69D10851CCBA7DFEB4D760119B99034B784F1B8081CC8AA7CDA71E54CAA2B74B",
    "msgName": "MsgUnjail",
    "version": 1,
    "msgIndex": 0,
    "validatorAddress": "tcrocncl1gs80n8fpc5mc3ywkgfy93l23tg0gdqj5m4uxzk"
}
```
