# Bank Module Event List
- [MSG_SEND_CREATED](#EVENT::MSG_SEND_CREATED)
- [MSG_SEND_FAILED](./README.md)
- [MSG_MULTI_SEND_CREATED](./README.md)
- [MSG_MULTI_SEND_FAILED](./README.md)


## event::MSG_SEND_CREATED
*Name* : MSG_SEND_CREATED

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key            | Type     | Description                                             |
| -------------- | -------- | ------------------------------------------------------- |
| `fromAddress`  | *string* | The source or the origin address                        |
| `toAddress`    | *string* | The destination or the recepient address                |
| `amount`       | *object* | CRO amount wrapper                                      |
| `amount.value` | *bigint* | CRO amount in basic unit                                |
| `msgName`      | *string* | Blockchain Message type . Value: `MsgSend`              |
| `txHash`       | *string* | TxID of the blockchain transaction containing the event |
| `msgIndex`     | *int*    | message index on the block                              |
| `name`         | *string* | Specific Event Name. Value: `MSG_SEND_CREATED`          |
| `version`      | *int*    | Event Version. Value: `1`                               |
| `height`       | *int64*  | Height of the block containing the transaction          |
| `uuid`         | *string* | Unique ID that is assigned on event creation            |