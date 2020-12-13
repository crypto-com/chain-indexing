# Block Module Event List
  - [event::TRANSACTION_CREATED](#event_transaction_created)
  - [event::TRANSACTION_FAILED](#event_transaction_failed)
  - [event::ACCOUNT_TRANSFERRED](#event_account_transferred)
  - [event::MINTED](#event_minted)
  - [event::POWER_CHANGED](#event_power_changed)
  - [event::VALIDATOR_SLASHED](#event_validator_slashed)
  - [event::VALIDATOR_JAILED](#event_validator_jailed)

## event::TRANSACTION_CREATED
*Name* : TransactionCreated

*Type* : [Base](../README.md#understanding_an_event)

*Structure* : 

| Key             | Type     | Description                                      |
| --------------- | -------- | ------------------------------------------------ |
| `txHash`        | *string* | Transaction Hash                                 |
| `fee`           | *bigint* | Fee amount in baseunit                           |
| `log`           | *string* | Stringified transaction state logs               |
| `code`          | *int*    | Transaction status code                          |
| `memo`          | *string* | User attached memo message                       |
| `gasUsed`       | *int*    | On chain gas amount in base unit                 |
| `feePayer`      | *string* | Fee payer candidate                              |
| `msgCount`      | *int*    | Number of messages in this transaction           |
| `gasWanted`     | *int*    | Gas estimated for this transaction               |
| `feeGranter`    | *string* | Fee granter if any                               |
| `timeoutHeight` | *int64*  | Block height at which the transaction timeouts   |
| `name`          | *string* | Specific Event Name. Value: `TransactionCreated` |
| `version`       | *int*    | Event Version. Value: `1`                        |
| `height`        | *int64*  | Height of the block containing the transaction   |
| `uuid`          | *string* | Unique ID that is assigned on event creation     |

*Example* :  
```json
{
    "fee": "0",
    "log": "[{\"msgIndex\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3\"},{\"key\":\"sender\",\"value\":\"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3\"},{\"key\":\"amount\",\"value\":\"60561202basetcro\"}]}]}]",
    "code": 0,
    "memo": "",
    "name": "TransactionCreated",
    "uuid": "da23ea7d-6afa-49ad-be8a-fb14b4febabc",
    "height": 69096,
    "txHash": "0E713183EBDE565408ADCD34D0F0A39C8342709E7863B260CB3816908C2EC824",
    "gasUsed": 51039,
    "version": 1,
    "feePayer": "",
    "msgCount": 1,
    "gasWanted": 200000,
    "feeGranter": "",
    "timeoutHeight": 0
}
```  

## event::TRANSACTION_FAILED
*Name* : TransactionFailed

*Type* : [Base](../README.md#understanding_an_event)

*Structure* : 

| Key             | Type     | Description                                     |
| --------------- | -------- | ----------------------------------------------- |
| `txHash`        | *string* | Transaction Hash                                |
| `fee`           | *bigint* | Fee amount in baseunit                          |
| `log`           | *string* | Stringified transaction state logs              |
| `code`          | *int*    | Transaction status code                         |
| `memo`          | *string* | User attached memo message                      |
| `gasUsed`       | *int*    | On chain gas amount in base unit                |
| `feePayer`      | *string* | Fee payer candidate                             |
| `msgCount`      | *int*    | Number of messages in this transaction          |
| `gasWanted`     | *int*    | Gas estimated for this transaction              |
| `feeGranter`    | *string* | Fee granter if any                              |
| `timeoutHeight` | *int64*  | Block height at which the transaction timeouts  |
| `name`          | *string* | Specific Event Name. Value: `TransactionFailed` |
| `version`       | *int*    | Event Version. Value: `1`                       |
| `height`        | *int64*  | Height of the block containing the transaction  |
| `uuid`          | *string* | Unique ID that is assigned on event creation    |

*Example* :  
```json
{
    "fee": "20000",
    "log": "failed to execute message; message index: 0: validator not jailed; cannot be unjailed",
    "code": 5,
    "memo": "",
    "name": "TransactionFailed",
    "uuid": "6b1b4af3-db62-4ead-8372-786f2c20db95",
    "height": 69154,
    "txHash": "09846EBF61641170178DC094FBB751AE025DB94DD0247339068D302B1999B4F8",
    "gasUsed": 52652,
    "version": 1,
    "feePayer": "",
    "msgCount": 1,
    "gasWanted": 200000,
    "feeGranter": "",
    "timeoutHeight": 0
}
```  