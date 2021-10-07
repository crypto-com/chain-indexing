package usecase_parser_test

const TX_MSG_EXEC_MSG_SEND_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "4BA2468ADCA23EAAD3CC2818C123173B44FE6CDC0FAD4D3A0799B0A47E8386E3",
      "parts": {
        "total": 1,
        "hash": "B4D103EB478FD8C566E2ACF6E88072FB0C1F77C3658B01926B8443BA0195A25A"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-4",
        "height": "113382",
        "time": "2021-08-26T04:56:23.854458166Z",
        "last_block_id": {
          "hash": "B60748462F8F3E14808979E698B4F2D062E02A90E330A36D3265FBB1B882AFF6",
          "parts": {
            "total": 1,
            "hash": "B3D5B588FF679FA6A37059327A19899A19D7A23B799A6CB7EA4BA304E842072F"
          }
        },
        "last_commit_hash": "F056E5D486FC2D5C9D9CB2774E19B59864A259F673C5E83B6FF8806CA482921F",
        "data_hash": "3ABD2ABA5F7D8A324600DEA2EBB9342971961CB8C14D05D9C7FEDB61A6D8A8E6",
        "validators_hash": "77B520746A7915359B36FBB0FC761C77528FCB68E5998412B6CCDA1CF6B30FC9",
        "next_validators_hash": "77B520746A7915359B36FBB0FC761C77528FCB68E5998412B6CCDA1CF6B30FC9",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "3A1900BB913898F91CFC5940C982591F461D062E8188199D7D8E7FC3E43F582E",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2"
      },
      "data": {
        "txs": [
          "CuYBCuMBCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxLBAQordGNybzE1emg1dG43eGpkZWN1NHpqY2xzbWxubGh0NWVhZDJteDg0Z2F1MhKRAQocL2Nvc21vcy5iYW5rLnYxYmV0YTEuTXNnU2VuZBJxCit0Y3JvMXZ1cmZocWYwajJqZ2ZwamFobGphNmc2dXEydHMycjYwc3dtMmQ5Eit0Y3JvMWE5M3lmbnNjM3g3bTBtNDQ1Y2pzdmVlMm43cXo5YzBwdXJsendxGhUKCGJhc2V0Y3JvEgkxMDAwMDAwMDASawpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA/ZmZrKTHzpFHfml++KJ3lO8qQysug3ls0/McPi6TY3oEgQKAggBGAMSFwoRCghiYXNldGNybxIFNTAwMDAQwJoMGkDUsthteS7lTq+Ldp1JriOoban1HEpio4TGcsDv1dcyei7iVxo2EGiLewWFVZ8RqgNM4SwukAoNmNca55eijx0z"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "113381",
        "round": 0,
        "block_id": {
          "hash": "B60748462F8F3E14808979E698B4F2D062E02A90E330A36D3265FBB1B882AFF6",
          "parts": {
            "total": 1,
            "hash": "B3D5B588FF679FA6A37059327A19899A19D7A23B799A6CB7EA4BA304E842072F"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "F9A4582896E5F2692BEDDD7C251755E33D2DE013",
            "timestamp": "2021-08-26T04:56:23.910207522Z",
            "signature": "LIGswOX+KFHB1wO/kuJDpNuk/DzwsX20U12gvOiAQ4HSl6F5zd1VlvxcNnyVxqQ5whBum38vQ0w4BP2hp8HVAQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2",
            "timestamp": "2021-08-26T04:56:23.854458166Z",
            "signature": "SrwtVL3ZX39u2Yb04plPhdsBb3M67bVggAOwrkji9NvJEAs8bGXoJhNEjjZFzRFxOcuJGtV1FQBDiGuvKqI/Aw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7E2B455523A35AE8F4064C00AF01515FEF673079",
            "timestamp": "2021-08-26T04:56:23.825722636Z",
            "signature": "OQEnrSTYgCWzkBb9SaZiTbEY1isS4AEdBkctN63nazbC/ynjBT8A8cLmRLaN5yc+1QesGVyZSZft5KaDabpwCQ=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_EXEC_MSG_SEND_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "113382",
    "txs_results": [
      {
        "code": 0,
        "data": "CiMKHS9jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dFeGVjEgIKAA==",
        "log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"tcro1a93yfnsc3x7m0m445cjsvee2n7qz9c0purlzwq\"},{\"key\":\"amount\",\"value\":\"100000000basetcro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9\"},{\"key\":\"amount\",\"value\":\"100000000basetcro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgExec\"},{\"key\":\"sender\",\"value\":\"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1a93yfnsc3x7m0m445cjsvee2n7qz9c0purlzwq\"},{\"key\":\"sender\",\"value\":\"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9\"},{\"key\":\"amount\",\"value\":\"100000000basetcro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "76874",
        "events": [
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "dGNybzE1emg1dG43eGpkZWN1NHpqY2xzbWxubGh0NWVhZDJteDg0Z2F1Mg==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMDBiYXNldGNybw==",
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
                "value": "NTAwMDBiYXNldGNybw==",
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
                "value": "dGNybzE1emg1dG43eGpkZWN1NHpqY2xzbWxubGh0NWVhZDJteDg0Z2F1Mg==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMDBiYXNldGNybw==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNybzE1emg1dG43eGpkZWN1NHpqY2xzbWxubGh0NWVhZDJteDg0Z2F1Mg==",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "dGNybzE1emg1dG43eGpkZWN1NHpqY2xzbWxubGh0NWVhZDJteDg0Z2F1Mi8z",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "MUxMWWJYa3U1VTZ2aTNhZFNhNGpxRzJwOVJ4S1lxT0V4bkxBNzlYWE1ub3U0bGNhTmhCb2kzc0ZoVldmRWFvRFRPRXNMcEFLRFpqWEd1ZVhvbzhkTXc9PQ==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2Nvc21vcy5hdXRoei52MWJldGExLk1zZ0V4ZWM=",
                "index": true
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "dGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkOQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwMDAwYmFzZXRjcm8=",
                "index": true
              }
            ]
          },
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "cmVjZWl2ZXI=",
                "value": "dGNybzFhOTN5Zm5zYzN4N20wbTQ0NWNqc3ZlZTJuN3F6OWMwcHVybHp3cQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwMDAwYmFzZXRjcm8=",
                "index": true
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "dGNybzFhOTN5Zm5zYzN4N20wbTQ0NWNqc3ZlZTJuN3F6OWMwcHVybHp3cQ==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "dGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkOQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwMDAwYmFzZXRjcm8=",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkOQ==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "YmFuaw==",
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
            "value": "NDgzNTIzNTYxOGJhc2V0Y3Jv",
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
            "value": "NDgzNTIzNTYxOGJhc2V0Y3Jv",
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
            "value": "NDgzNTIzNTYxOGJhc2V0Y3Jv",
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
            "value": "NDgzNTIzNTYxOGJhc2V0Y3Jv",
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
            "value": "NDgzNTIzNTYxOGJhc2V0Y3Jv",
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
            "value": "MC4wMDA1OTUxMTA0MTAwMjk0NDI=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTIxMDc1NzA5NzQ0NzEyMDI=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MzA1MTc2ODYzMTI2Nzk1NjIuODc3ODk2NzczMzAzOTk3MDQ4",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDgzNTIzNTYxOA==",
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
            "value": "NDgzNTIzNTYxOGJhc2V0Y3Jv",
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
            "value": "NDgzNTIzNTYxOGJhc2V0Y3Jv",
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
            "value": "NDgzNTIzNTYxOGJhc2V0Y3Jv",
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
            "value": "MjQxNzYxNzgwLjkwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
            "value": "MjQxNzYxNzguMDkwMDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
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
            "value": "MjQxNzYxNzgwLjkwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
            "value": "MTUzMTE2MTAxLjIwOTI4NDIzNzkyNTc0OTQyMmJhc2V0Y3Jv",
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
            "value": "MTUzMTE2MTAxMi4wOTI4NDIzNzkyNTc0OTQyMTliYXNldGNybw==",
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
            "value": "MTUzMTE1Nzk0LjM2NTg0NjAxNTY1MDg2MjcxNmJhc2V0Y3Jv",
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
            "value": "MTUzMTE1Nzk0My42NTg0NjAxNTY1MDg2MjcxNjNiYXNldGNybw==",
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
            "value": "MTUzMTE1NDg4LjEzNDg2OTc0NTk2NDA0MDQ3OGJhc2V0Y3Jv",
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
            "value": "MTUzMTE1NDg4MS4zNDg2OTc0NTk2NDA0MDQ3ODBiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxczRnZ3EyenV6dndnNWs4dm54Mnhmd3RkbTRjejZ3dG51cWtsN2E=",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": null,
    "validator_updates": null,
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
