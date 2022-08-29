# Slashing Module Event List
 - [event::MSG_UNJAIL_CREATED](#event_MSG_UNJAIL_CREATED)
 - [event::MSG_UNJAIL_FAILED](#event_MSG_UNJAIL_FAILED)

## event::MSG_UNJAIL_CREATED
*Name* : /cosmos.slashing.v1beta1.MsgUnjail.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                              |
| ------------------ | -------- | ------------------------------------------------------------------------ |
| `validatorAddress` | *string* | Validator address                                                        |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgUnjail`                             |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event                  |
| `msgIndex`         | *int*    | message index on the block                                               |
| `name`             | *string* | Specific Event Name. Value: `/cosmos.slashing.v1beta1.MsgUnjail.Created` |
| `version`          | *int*    | Event Version. Value: `1`                                                |
| `height`           | *int64*  | Height of the block containing the transaction                           |
| `uuid`             | *string* | Unique ID that is assigned on event creation                             |

*Example* :

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
*Name* : /cosmos.slashing.v1beta1.MsgUnjail.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key                | Type     | Description                                                             |
| ------------------ | -------- | ----------------------------------------------------------------------- |
| `validatorAddress` | *string* | Validator address                                                       |
| `msgName`          | *string* | Blockchain Message type . Value: `MsgUnjail`                            |
| `txHash`           | *string* | TxID of the blockchain transaction containing the event                 |
| `msgIndex`         | *int*    | message index on the block                                              |
| `name`             | *string* | Specific Event Name. Value: `/cosmos.slashing.v1beta1.MsgUnjail.Failed` |
| `version`          | *int*    | Event Version. Value: `1`                                               |
| `height`           | *int64*  | Height of the block containing the transaction                          |
| `uuid`             | *string* | Unique ID that is assigned on event creation                            |

*Example* :

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
