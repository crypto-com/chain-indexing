# Bank Module Event List
- [MSG_SEND_CREATED](./README.md#event::MSG_SEND_CREATED)
- [MSG_SEND_FAILED](./README.md)
- [MSG_MULTI_SEND_CREATED](./README.md)
- [MSG_MULTI_SEND_FAILED](./README.md)


## event::MSG_SEND_CREATED
*Name* : MSG_SEND_CREATED

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key           | Type     | Description                                             |
| ------------- | -------- | ------------------------------------------------------- |
| `fromAddress` | *string* | The source or the origin address                        |
| `toAddress`   | *string* | The destination or the recepient address                |
| `amount`      | *object* | CRO amount in basic unit                                |
| `msgName`     | *string* | Blockchain Message type . Value: `MsgSend`              |
| `txHash`      | *string* | TxID of the blockchain transaction containing the event |
| `msgIndex`    | *int*    | message index on the block                              |
| `name`        | *string* | Specific Event Name. Value: `MsgSendCreated`          |
| `version`     | *int*    | Event Version. Value: `1`                               |
| `height`      | *int64*  | Height of the block containing the transaction          |
| `uuid`        | *string* | Unique ID that is assigned on event creation            |

*Example* :

```json
{
    "name": "MsgSendCreated",
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