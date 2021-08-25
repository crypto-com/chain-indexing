package usecase_parser_test

const TX_MSG_CREATE_VALIDATOR_V0_43_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "0AB6259FB2FE569DE413003CED96A765B4442C6029E64D50C07BC696B3123E53",
      "parts": {
        "total": 1,
        "hash": "9DFF6A0BA7420D284282F94C2870275C57D17147B0B1A2C0BB2AA75634A72D09"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-4",
        "height": "3106",
        "time": "2021-08-19T09:26:15.100668005Z",
        "last_block_id": {
          "hash": "F24928F2F2203B5B3D41BF3B11B2D15A8BB78691526A2DF92A5DC8020D48ACDC",
          "parts": {
            "total": 1,
            "hash": "09D546A6A2A6816F0712D70FE565872F41C4FB3B1F734255F1C39285D089EA56"
          }
        },
        "last_commit_hash": "A4F46F631A109D5EC02CDBE50610CFCFAA49625A369A4972D76642D324D8F0B7",
        "data_hash": "9098E35870D8B5C11AA49F327B85892BBF93E0D7C6D948DB748639ECB68ED2E7",
        "validators_hash": "EB1E4C26CAA143F91889147E7D752C09A0019DA2C8F7E74D0A3EB4760DF4CDB3",
        "next_validators_hash": "EB1E4C26CAA143F91889147E7D752C09A0019DA2C8F7E74D0A3EB4760DF4CDB3",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "A0A3156D8DE3D25A9267F89EBC38535F166CFE66ABFC897465893BA8D70B1311",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2"
      },
      "data": {
        "txs": [
          "CsICCr8CCiovY29zbW9zLnN0YWtpbmcudjFiZXRhMS5Nc2dDcmVhdGVWYWxpZGF0b3ISkAIKEQoPYzRfbm9kZV90ZXN0aW5nEjsKEjEwMDAwMDAwMDAwMDAwMDAwMBISMjAwMDAwMDAwMDAwMDAwMDAwGhExMDAwMDAwMDAwMDAwMDAwMBoBMSIrdGNybzFrNG56MnlhbHN1enV6eGV4Y3h1dWdocDMzY2ZxbHM3ZTZ1dGFyeCovdGNyb2NuY2wxazRuejJ5YWxzdXp1enhleGN4dXVnaHAzM2NmcWxzN2Uwcmd5bTkyQwodL2Nvc21vcy5jcnlwdG8uZWQyNTUxOS5QdWJLZXkSIgoghTpJIeVSSo1PNU87B8y1vMvpyv3ShMmu2I0vYON1w2o6GAoIYmFzZXRjcm8SDDEwMDAwMDAwMDAwMBJqClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiECeexZEAnUYmqXO+m681kxBV0pxgKaw8CIlkKG9/dNdoQSBAoCCAEYARIWChAKCGJhc2V0Y3JvEgQ1MDAwEMCaDBpArJg3VIIFTyiQF35d0jAMVd3046kPdcj9heywiBk8+8g3GHW9/qs+YpEVSo4ZyTdnK4aWRXp1bUKoZlxhI5iRvg=="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "3105",
        "round": 0,
        "block_id": {
          "hash": "F24928F2F2203B5B3D41BF3B11B2D15A8BB78691526A2DF92A5DC8020D48ACDC",
          "parts": {
            "total": 1,
            "hash": "09D546A6A2A6816F0712D70FE565872F41C4FB3B1F734255F1C39285D089EA56"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2",
            "timestamp": "2021-08-19T09:26:15.100668005Z",
            "signature": "rpJTeCLh+Yoo1vnW2y7lDHSrA88cIvpkcCSWe9wb+A5TmLIMbpF1PIyG8s9JUyJmfi7SxhYwWiwa2ixzLRi0Bg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7E2B455523A35AE8F4064C00AF01515FEF673079",
            "timestamp": "2021-08-19T09:26:15.162833356Z",
            "signature": "hEzLzV06YHFoG7bu15nigOTxWtuLnEzvCFJ457kMfH9NZGbsNi3iLslIF03Z162pB70UxWK9lPowNDhi542iAA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F9A4582896E5F2692BEDDD7C251755E33D2DE013",
            "timestamp": "2021-08-19T09:26:15.047894049Z",
            "signature": "ek9JkkgPP+8ERmth2TBD+NJDDQVnwG8QjRsYRRHUhqA9kIcJMFrBRmw1IJOQOisgZ1EI9cd76Hn/+Vps9dAmCg=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_CREATE_VALIDATOR_V0_43_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "3106",
    "txs_results": [
      {
        "code": 0,
        "data": "CiwKKi9jb3Ntb3Muc3Rha2luZy52MWJldGExLk1zZ0NyZWF0ZVZhbGlkYXRvcg==",
        "log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"tcro1tygms3xhhs3yv487phx3dw4a95jn7t7lh45rnr\"},{\"key\":\"amount\",\"value\":\"100000000000basetcro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"tcro1k4nz2yalsuzuzxexcxuughp33cfqls7e6utarx\"},{\"key\":\"amount\",\"value\":\"100000000000basetcro\"}]},{\"type\":\"create_validator\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl1k4nz2yalsuzuzxexcxuughp33cfqls7e0rgym9\"},{\"key\":\"amount\",\"value\":\"100000000000basetcro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.staking.v1beta1.MsgCreateValidator\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"tcro1k4nz2yalsuzuzxexcxuughp33cfqls7e6utarx\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "144421",
        "events": [
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "dGNybzFrNG56MnlhbHN1enV6eGV4Y3h1dWdocDMzY2ZxbHM3ZTZ1dGFyeA==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMGJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "cmVjZWl2ZXI=",
                "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMGJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "dGNybzFrNG56MnlhbHN1enV6eGV4Y3h1dWdocDMzY2ZxbHM3ZTZ1dGFyeA==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMGJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNybzFrNG56MnlhbHN1enV6eGV4Y3h1dWdocDMzY2ZxbHM3ZTZ1dGFyeA==",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "dGNybzFrNG56MnlhbHN1enV6eGV4Y3h1dWdocDMzY2ZxbHM3ZTZ1dGFyeC8x",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "ckpnM1ZJSUZUeWlRRjM1ZDBqQU1WZDMwNDZrUGRjajloZXl3aUJrOCs4ZzNHSFc5L3FzK1lwRVZTbzRaeVRkbks0YVdSWHAxYlVLb1pseGhJNWlSdmc9PQ==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2Nvc21vcy5zdGFraW5nLnYxYmV0YTEuTXNnQ3JlYXRlVmFsaWRhdG9y",
                "index": true
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "dGNybzFrNG56MnlhbHN1enV6eGV4Y3h1dWdocDMzY2ZxbHM3ZTZ1dGFyeA==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwMDAwMDAwYmFzZXRjcm8=",
                "index": true
              }
            ]
          },
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "cmVjZWl2ZXI=",
                "value": "dGNybzF0eWdtczN4aGhzM3l2NDg3cGh4M2R3NGE5NWpuN3Q3bGg0NXJucg==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwMDAwMDAwYmFzZXRjcm8=",
                "index": true
              }
            ]
          },
          {
            "type": "create_validator",
            "attributes": [
              {
                "key": "dmFsaWRhdG9y",
                "value": "dGNyb2NuY2wxazRuejJ5YWxzdXp1enhleGN4dXVnaHAzM2NmcWxzN2Uwcmd5bTk=",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwMDAwMDAwYmFzZXRjcm8=",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "c3Rha2luZw==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "dGNybzFrNG56MnlhbHN1enV6eGV4Y3h1dWdocDMzY2ZxbHM3ZTZ1dGFyeA==",
                "index": true
              }
            ]
          }
        ],
        "codespace": ""
      }
    ],
    "begin_block_events": [
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3N2Jhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coinbase",
        "attributes": [
          {
            "key": "bWludGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3N2Jhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3N2Jhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3N2Jhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3N2Jhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC4wMDA1OTUyMzQ1ODA4MzQ2OTI=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTIwMDI5NDU4ODg3NTk1MjE=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MzAyNDc2MDIyMjc0OTc5OTIuMTY1MTAwODI4NDE2OTcwMDg0",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3Nw==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3N2Jhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNybzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODMzOXA0bA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3N2Jhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODMzOXA0bA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5MjQ0MzM3N2Jhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "proposer_reward",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjM5NjIyMTY4Ljg1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQ=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjM5NjIyMTYuODg1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjM5NjIyMTY4Ljg1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQ=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUxNzYwNzA2LjkzODMzMzMzMzE4MTU3MjYyNmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNGZ6a3N6NWg3MmV0NHNzanRxcHdzbWh6NnlzazZyNG5hNXRyNjM=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUxNzYwNzA2OS4zODMzMzMzMzE4MTU3MjYyNjNiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNGZ6a3N6NWg3MmV0NHNzanRxcHdzbWh6NnlzazZyNG5hNXRyNjM=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUxNzYwNzA2LjkzODMzMzMzMzE4MTU3MjYyNmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxczRnZ3EyenV6dndnNWs4dm54Mnhmd3RkbTRjejZ3dG51cWtsN2E=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUxNzYwNzA2OS4zODMzMzMzMzE4MTU3MjYyNjNiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxczRnZ3EyenV6dndnNWs4dm54Mnhmd3RkbTRjejZ3dG51cWtsN2E=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUxNzYwNzA2LjkzODMzMzMzMzE4MTU3MjYyNmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUxNzYwNzA2OS4zODMzMzMzMzE4MTU3MjYyNjNiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQ=",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNybzF0eWdtczN4aGhzM3l2NDg3cGh4M2R3NGE5NWpuN3Q3bGg0NXJucg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNybzFmbDQ4dnNubXNkemN2ODVxNWQycTR6NWFqZGhhOHl1M3I0Z2o5aA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzFmbDQ4dnNubXNkemN2ODVxNWQycTR6NWFqZGhhOHl1M3I0Z2o5aA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzF0eWdtczN4aGhzM3l2NDg3cGh4M2R3NGE5NWpuN3Q3bGg0NXJucg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzF0eWdtczN4aGhzM3l2NDg3cGh4M2R3NGE5NWpuN3Q3bGg0NXJucg==",
            "index": true
          }
        ]
      }
    ],
    "validator_updates": [
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "hTpJIeVSSo1PNU87B8y1vMvpyv3ShMmu2I0vYON1w2o="
            }
          }
        },
        "power": "100000"
      }
    ],
    "consensus_param_updates": {
      "block": {
        "max_bytes": "500000",
        "max_gas": "81500000"
      },
      "evidence": {
        "max_age_num_blocks": "403200",
        "max_age_duration": "2419200000000000",
        "max_bytes": "150000"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      }
    }
  }
}
`
