# Bank Module Event List

-   [event::MSG_SEND_CREATED](#eventmsg_send_created)
-   [event::MSG_SEND_FAILED](#eventmsg_send_failed)
-   [event::MSG_MULTI_SEND_CREATED](#eventmsg_multi_send_created)
-   [event::MSG_MULTI_SEND_FAILED](#eventmsg_multi_send_failed)

## event::MSG_SEND_CREATED

*Name* : /cosmos.bank.v1beta1.MsgSend.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* :

| Key           | Type     | Description                                                        |
| ------------- | -------- | ------------------------------------------------------------------ |
| `fromAddress` | *string* | The source or the origin address                                   |
| `toAddress`   | *string* | The destination or the recipient address                           |
| `amount`      | *object* | CRO amount in basic unit                                           |
| `msgName`     | *string* | Blockchain Message type . Value: `MsgSend`                         |
| `txHash`      | *string* | TxID of the blockchain transaction containing the event            |
| `msgIndex`    | *int*    | message index on the block                                         |
| `name`        | *string* | Specific Event Name. Value: `/cosmos.bank.v1beta1.MsgSend.Created` |
| `version`     | *int*    | Event Version. Value: `1`                                          |
| `height`      | *int64*  | Height of the block containing the transaction                     |
| `uuid`        | *string* | Unique ID that is assigned on event creation                       |

*Example* :

```json
{
    "name": "/cosmos.bank.v1beta1.MsgSend.Created",
    "uuid": "d125ad50-14ce-4e58-b6d5-292db54534f4",
    "amount": "1000000000",
    "height": 29421,
    "txHash": "1B73517984DAD6CB9D19390541A24849E7E9F8A10B2B072F30AD2B62B698A6E7",
    "msgName": "MsgSend",
    "version": 1,
    "msgIndex": 0,
    "toAddress": "tcro1j7pej8kplem4wt50p4hfvndhuw5jprxxn5625q",
    "fromAddress": "tcro1nj7zlmkuek5rl67ew2k8cle7cyalp3p6a9tj5t"
}
```

## event::MSG_SEND_FAILED

*Name* : /cosmos.bank.v1beta1.MsgSend.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* :

| Key           | Type     | Description                                                       |
| ------------- | -------- | ----------------------------------------------------------------- |
| `fromAddress` | *string* | The source or the origin address                                  |
| `toAddress`   | *string* | The destination or the recipient address                          |
| `amount`      | *object* | CRO amount in basic unit                                          |
| `msgName`     | *string* | Blockchain Message type . Value: `MsgSend`                        |
| `txHash`      | *string* | TxID of the blockchain transaction containing the event           |
| `msgIndex`    | *int*    | message index on the block                                        |
| `name`        | *string* | Specific Event Name. Value: `/cosmos.bank.v1beta1.MsgSend.Failed` |
| `version`     | *int*    | Event Version. Value: `1`                                         |
| `height`      | *int64*  | Height of the block containing the transaction                    |
| `uuid`        | *string* | Unique ID that is assigned on event creation                      |

*Example* :

```json
{
    "name": "/cosmos.bank.v1beta1.MsgSend.Failed",
    "uuid": "f2f00a5b-0cef-4652-b12f-e31bac7cb927",
    "amount": "1000000000",
    "height": 115930,
    "txHash": "B2D15532E2DD5160EA9CA137E96AF0142E86E6A89C0C24C01C0F5AF49689C901",
    "msgName": "MsgSend",
    "version": 1,
    "msgIndex": 0,
    "toAddress": "tcro1pet9pezper24qmf5k23wkews8ha68xs2vz00q9",
    "fromAddress": "tcro17wnekjfsllm8au3e8yuptxd24zll3m55655wl9"
}
```

## event::MSG_MULTI_SEND_CREATED

*Name* : /cosmos.bank.v1beta1.MsgMultiSend.Created

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* :

| Key                 | Type            | Description                                                             |
| ------------------- | --------------- | ----------------------------------------------------------------------- |
| `inputs`            | _array(object)_ | Inputs array for the multisig tx                                        |
| `inputs[].address`  | *string*        | Participating Input blockchain address                                  |
| `inputs[].amount`   | *bigint*        | Participating Input amount                                              |
| `outputs`           | _array(object)_ | Outputs array for the multisig tx                                       |
| `outputs[].address` | *string*        | Participating Input blockchain address                                  |
| `outputs[].amount`  | *bigint*        | Participating Input amount                                              |
| `msgName`           | *string*        | Blockchain Message type. Value: `MsgMultiSend`                          |
| `txHash`            | *string*        | TxID of the blockchain transaction containing the event                 |
| `msgIndex`          | *int*           | message index on the block                                              |
| `name`              | *string*        | Specific Event Name. Value: `/cosmos.bank.v1beta1.MsgMultiSend.Created` |
| `version`           | *int*           | Event Version. Value: `1`                                               |
| `height`            | *int64*         | Height of the block containing the transaction                          |
| `uuid`              | *string*        | Unique ID that is assigned on event creation                            |

*Example* : T.B.D

## event::MSG_MULTI_SEND_FAILED

*Name* : /cosmos.bank.v1beta1.MsgMultiSend.Failed

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* :

| Key                 | Type            | Description                                                            |
| ------------------- | --------------- | ---------------------------------------------------------------------- |
| `inputs`            | _array(object)_ | Inputs array for the multisig tx                                       |
| `inputs[].address`  | *string*        | Participating Input blockchain address                                 |
| `inputs[].amount`   | *bigint*        | Participating Input amount                                             |
| `outputs`           | _array(object)_ | Outputs array for the multisig tx                                      |
| `outputs[].address` | *string*        | Participating Input blockchain address                                 |
| `outputs[].amount`  | *bigint*        | Participating Input amount                                             |
| `msgName`           | *string*        | Blockchain Message type. Value: `MsgMultiSend`                         |
| `txHash`            | *string*        | TxID of the blockchain transaction containing the event                |
| `msgIndex`          | *int*           | message index on the block                                             |
| `name`              | *string*        | Specific Event Name. Value: `/cosmos.bank.v1beta1.MsgMultiSend.Failed` |
| `version`           | *int*           | Event Version. Value: `1`                                              |
| `height`            | *int64*         | Height of the block containing the transaction                         |
| `uuid`              | *string*        | Unique ID that is assigned on event creation                           |

*Example* : T.B.D
