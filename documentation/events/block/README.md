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

## event::RAW_BLOCK_CREATED
*Name* : RawBlockCreated

*Type* : [Base](../README.md#understanding_an_event)

*Structure* : (T.B.D)

| Key        | Type     | Description                                       |
| ---------- | -------- | ------------------------------------------------- |
| `rawBlock` | *object* | Structure for Tendermint /block API response JSON |
| `name`     | *string* | Specific Event Name. Value: `RawBlockCreated`     |
| `version`  | *int*    | Event Version. Value: `1`                         |
| `height`   | *int64*  | Height of the block containing the transaction    |
| `uuid`     | *string* | Unique ID that is assigned on event creation      |

*Example* :  
```json
{
    "name": "RawBlockCreated",
    "uuid": "b7683cd9-2190-4346-b4c2-81b7c0c7fe8b",
    "height": 69091,
    "version": 1,
    "rawBlock": {
        "block": {
            "data": {
                "txs": []
            },
            "header": {
                "time": "2020-10-19T18:23:01.734341392Z",
                "height": "69091",
                "version": {
                    "block": "11"
                },
                "app_hash": "6A655A2A5262C6EC407AB22F15C60CA2BC6852873D0A374F8944F9A49D30B30D",
                "chain_id": "testnet-croeseid-1",
                "data_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "last_block_id": {
                    "hash": "E8DD5B4316B8A3C745713B246B8239E13419EE457FE84966662FC184FFBCB1F2",
                    "parts": {
                        "hash": "D87906584F4DC755EB6445A24A75B900DDD1AA05AE55B398F711253DCF2BB6D2",
                        "total": 1
                    }
                },
                "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
                "validators_hash": "430EABF2FD3BD8B1EF2309D0DCFF2C228F1728D66816C6A223B1630869873FF9",
                "last_commit_hash": "FD7160DE3F2548BB759866590F2EAC2F12C44AA1EFBEF4EFC51C19C4AF0D6DAB",
                "proposer_address": "4D9F47C5A19D5550685D2D55DF2FF8C5D7504CEB",
                "last_results_hash": "B8ED77450942E97670532D596E572C23BB3BEA5BEE607C7A66F3314195EB31CE",
                "next_validators_hash": "430EABF2FD3BD8B1EF2309D0DCFF2C228F1728D66816C6A223B1630869873FF9"
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "round": 0,
                "height": "69090",
                "block_id": {
                    "hash": "E8DD5B4316B8A3C745713B246B8239E13419EE457FE84966662FC184FFBCB1F2",
                    "parts": {
                        "hash": "D87906584F4DC755EB6445A24A75B900DDD1AA05AE55B398F711253DCF2BB6D2",
                        "total": 1
                    }
                },
                "signatures": [
                    {
                        "signature": "l5L/SCpT9LL9uiVeNfmdqhj+NuWICoKeaX2dlXSSwpYtaU5wzwaOWjM8iX+CwOLtmCE/4GY88b5IHMPDwk5PAA==",
                        "timestamp": "2020-10-19T18:23:01.744252687Z",
                        "block_id_flag": 2,
                        "validator_address": "A1E8AAEBBC82929B852748734BA39D67A62F201B"
                    }
                ]
            }
        },
        "block_id": {
            "hash": "7B2BDC15C73DC6E7A602153B28B305475FFD180513030D510643E15D594C3312",
            "parts": {
                "hash": "1447EB01153B903A4D5EEAED1732A3CD6206C98DC79EE4C0115D260D16C94B32",
                "total": 1
            }
        }
    }
}
```  

## event::BLOCK_PROPOSER_REWARDED
*Name* : BlockProposerRewarded

*Type* : [Base](../README.md#understanding_an_event)

*Structure* : 

| Key         | Type     | Description                                         |
| ----------- | -------- | --------------------------------------------------- |
| `validator` | *string* | Validator address                                   |
| `amount`    | *string* | Reward amount in decimals                           |
| `name`      | *string* | Specific Event Name. Value: `BlockProposerRewarded` |
| `version`   | *int*    | Event Version. Value: `1`                           |
| `height`    | *int64*  | Height of the block containing the transaction      |
| `uuid`      | *string* | Unique ID that is assigned on event creation        |

*Example* :  
```json
{
    "name": "BlockProposerRewarded",
    "uuid": "88dfc53a-3c2e-4d29-a1b6-c910f2a9bc5f",
    "amount": "833034152.277599639227521390",
    "height": 69090,
    "version": 1,
    "validator": "tcrocncl18ylchgmxyphw3ctsl75n53ujequkmmag2n6x3f"
}
```  

## event::BLOCK_REWARDED
*Name* : BlockRewarded

*Type* : [Base](../README.md#understanding_an_event)

*Structure* : 

| Key         | Type     | Description                                    |
| ----------- | -------- | ---------------------------------------------- |
| `validator` | *string* | Validator address                              |
| `amount`    | *string* | Reward amount in decimals                      |
| `name`      | *string* | Specific Event Name. Value: `BlockRewarded`    |
| `version`   | *int*    | Event Version. Value: `1`                      |
| `height`    | *int64*  | Height of the block containing the transaction |
| `uuid`      | *string* | Unique ID that is assigned on event creation   |

*Example* :  
```json
{
    "name": "BlockRewarded",
    "uuid": "d20be8fc-a2df-430c-9a33-9527b1fa7964",
    "amount": "1521432571.172691693837460632",
    "height": 69090,
    "version": 1,
    "validator": "tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6"
}
```  