## Event Details
*Name* : MSG_SEND_CREATED

*Type* : [MsgBase](../README.md#MsgBase)

*Structure* : 

| Key            | Type     | Description                                             |
| -------------- | -------- | ------------------------------------------------------- |
| `fromAddress`  | *string* | The source or the origin address                        |
| `toAddress`    | *string* | The destination or the recepient address                |
| `amount`       | *object* | CRO amount wrapper                                      |
| `amount.value` | *bigint* | CRO amount in basic unit                                |
| `msgName`      | *string* | Blockchain Message name. Value: `MsgSend`               |
| `txHash`       | *string* | TxID of the blockchain transaction containing the event |
| `msgIndex`     | *int*    | message index on the block                              |
| `name`         | *string* | Specific Event Name. Value: `MSG_SEND_CREATED`          |
| `version`      | *int*    | Event Version. Value: `1`                               |
| `height`       | *int64*  | Height of the block containing the transaction          |
| `uuid`         | *string* | Unique ID that is assigned on event creation            |