# Block Module Event List
  - [event::BLOCK_CREATED](#event_block_created)
  - [event::RAW_BLOCK_CREATED](#event_raw_block_created)
  - [event::BLOCK_PROPOSER_REWARDED](#event_block_proposer_rewarded)
  - [event::BLOCK_REWARDED](#event_block_rewarded)
  - [event::BLOCK_COMMISSIONED](#event_block_commissioned)

## event::BLOCK_CREATED
*Name* : BlockCreated

*Type* : [Base](../README.md#understanding_an_event)

*Structure* : 

| Key                                | Type            | Description                                         |
| ---------------------------------- | --------------- | --------------------------------------------------- |
| `block`                            | *object*        | Block details wrapper                               |
| `block.height`                     | *int64*         | Block height                                        |
| `block.hash`                       | *string*        | Block hash                                          |
| `block.time`                       | *string*        | Block timestamp. Golang type: `utctime.UTCTime`     |
| `block.appHash`                    | *string*        | Block application hash                              |
| `block.proposerAddress`            | *object*        | Block proposer address                              |
| `block.txs`                        | *array(string)* | Array of encoded and stringified Transactions       |
| `block.signature`                  | *array(object)* | Block signature wrapper                             |
| `block.signature.blockIdFlag`      | *int*           | Block signature Flag ID                             |
| `block.signature.validatorAddress` | *string*        | Block signing validator address                     |
| `block.signature.timestamp`        | *string*        | Signature timestamp. Golang type: `utctime.UTCTime` |
| `block.signature.signature`        | *string*        | Signature value encoded as `base64`                 |
| `name`                             | *string*        | Specific Event Name. Value: `BlockCreated`          |
| `version`                          | *int*           | Event Version. Value: `1`                           |
| `height`                           | *int64*         | Height of the block containing the transaction      |
| `uuid`                             | *string*        | Unique ID that is assigned on event creation        |

*Example* :  
```json
{
    "name": "BlockCreated",
    "uuid": "13b11ec0-0bdd-4670-b973-eec3e5c6a698",
    "block": {
        "txs": [
            "CpMBCpABChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnAKK3Rjcm8xNjV0emNyaDJ5bDgzZzhxZXF4dWVnMmc1Z3pndTU3eTNmZTNrYzMSK3Rjcm8xODRsdGEybHN5dTQ3dnd5cDJlOHptdGNhM2s1eXE4NXA2YzR2cDMaFAoIYmFzZXRjcm8SCDYwNTYxMjAyElkKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQIIi3p/bsKb7K3mIm8IJ0Lc7qYpe5MjSZeKExt1peqCBBIECgIIARirBxIEEMCaDBpAT2f5hdgcqRhtGE0McGt/FSt0RiDTpIWAgt1GeyWfZSxobgHun0MW5OxAJ8IQi5ME5CbEHJBJpPH8aeOBEx35Vw=="
        ],
        "hash": "1D91E55C40D8CCA378DCCDC1C1B61F6B0745CD99F6B0B7875A93916305F4DFE3",
        "time": "2020-10-19T18:23:33.34998225Z",
        "height": 69096,
        "appHash": "952AFF3E808318FEDA01282874310F7C475CF25E8141ABDE4EE23B367092BED2",
        "signature": [
            {
                "signature": "Ul9OKX815AgK8jVpXTq6cClAfFXdnQ+FvUE6Uwl9TlEVqnijZKEzj1CrWT9FMAQEItIjNPrGHSTJdZQswyw2Cg==",
                "timestamp": "2020-10-19T18:23:33.30156328Z",
                "blockIdFlag": 2,
                "validatorAddress": "A1E8AAEBBC82929B852748734BA39D67A62F201B"
            }
        ],
        "proposerAddress": "679424928F3AB39FD197BEA70896D3D4E107D9FC"
    },
    "height": 69096,
    "version": 1
}
```