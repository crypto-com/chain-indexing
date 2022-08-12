package usecase_parser_test

const TX_MSG_EXEC_MSG_DELEGATE_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "A7E71E7931CB42BC7BE4A7397D32F9AB5985751B2B053756498601CEAB679AD5",
      "parts": {
        "total": 1,
        "hash": "751231A4D4A9383D8C69497E8B5D9F3BB13A5A77E55376C9641365F8F8DCAB45"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-4",
        "height": "170493",
        "time": "2021-08-29T17:15:46.150170633Z",
        "last_block_id": {
          "hash": "B419EF055C39B597CB3ECF6A62AE588527EDA2BCCA5D96679E02AE292BACAA87",
          "parts": {
            "total": 1,
            "hash": "CB6D2047E8EC5EE8BA1D1A145AD0A181D3335B827BA0C55E956073D9D89FF14F"
          }
        },
        "last_commit_hash": "2B293501E6C1EBB195A4BE5BB6BC3DD002305EBC10D4BF65790F8BB9A983A550",
        "data_hash": "020E880F725CFBD904B3266E6526258373292D9C26382555E50A0BB529F521E7",
        "validators_hash": "77B520746A7915359B36FBB0FC761C77528FCB68E5998412B6CCDA1CF6B30FC9",
        "next_validators_hash": "77B520746A7915359B36FBB0FC761C77528FCB68E5998412B6CCDA1CF6B30FC9",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "3B0D6B9E55CFE1A7B679F70FCF1E36B5236042200AB344FE2AD363C82C114ED0",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2"
      },
      "data": {
        "txs": [
          "CvEBCu4BCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxLMAQordGNybzE1emg1dG43eGpkZWN1NHpqY2xzbWxubGh0NWVhZDJteDg0Z2F1MhKcAQojL2Nvc21vcy5zdGFraW5nLnYxYmV0YTEuTXNnRGVsZWdhdGUSdQordGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkORIvdGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQaFQoIYmFzZXRjcm8SCTEwMDAwMDAwMBJsClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiED9mZmspMfOkUd+aX74oneU7ypDKy6DeWzT8xw+LpNjegSBAoCCAEYDhIYChIKCGJhc2V0Y3JvEgYxMDAwMDAQwJoMGkAlVwW36DZkZ3MtfRH2xorcF1N85ubGD1Kgws8+TcDkq1BgaIhBJQmqNu9KpFHprSYAxqY14faXpJgQK55M/o/H"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "170492",
        "round": 0,
        "block_id": {
          "hash": "B419EF055C39B597CB3ECF6A62AE588527EDA2BCCA5D96679E02AE292BACAA87",
          "parts": {
            "total": 1,
            "hash": "CB6D2047E8EC5EE8BA1D1A145AD0A181D3335B827BA0C55E956073D9D89FF14F"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "F9A4582896E5F2692BEDDD7C251755E33D2DE013",
            "timestamp": "2021-08-29T17:15:46.199521055Z",
            "signature": "m+H5ZtQFrzUxBSs/qQXYGTLVjIIyxGF++kefr6//m9mfKW16mkjqbhZfMYLf209SsosIdUzbA9ahGe1V3kziAQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2",
            "timestamp": "2021-08-29T17:15:46.150170633Z",
            "signature": "LJ/xK9ZbD6nCCAKQvjWoEErx8IsRZP5JMtNUJ+M/hvGITuz9LwwfxmybMVreKI7miYsZvTNCQb+vKG3e+313BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7E2B455523A35AE8F4064C00AF01515FEF673079",
            "timestamp": "2021-08-29T17:15:46.090847978Z",
            "signature": "apTaYIh2ueNK2jvBaRKIq5m9p2XtLljc04wfKX6BVBrm6ewR62C49J1jXcBA8Ae99/2bpBfLWuKYLSbo3v/tDA=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_EXEC_MSG_DELEGATE_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "170493",
    "txs_results": [
      {
        "code": 0,
        "data": "CiMKHS9jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dFeGVjEgIKAA==",
        "log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"tcro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3r4gj9h\"},{\"key\":\"amount\",\"value\":\"100000000basetcro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9\"},{\"key\":\"amount\",\"value\":\"100000000basetcro\"}]},{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl163tv59yzgeqcap8lrsa2r4zk580h8ddr5a0sdd\"},{\"key\":\"amount\",\"value\":\"100000000basetcro\"},{\"key\":\"new_shares\",\"value\":\"100000000.000000000000000000\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgExec\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "120155",
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
                "value": "MTAwMDAwYmFzZXRjcm8=",
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
                "value": "MTAwMDAwYmFzZXRjcm8=",
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
                "value": "MTAwMDAwYmFzZXRjcm8=",
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
                "value": "dGNybzE1emg1dG43eGpkZWN1NHpqY2xzbWxubGh0NWVhZDJteDg0Z2F1Mi8xNA==",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "SlZjRnQrZzJaR2R6TFgwUjlzYUszQmRUZk9ibXhnOVNvTUxQUGszQTVLdFFZR2lJUVNVSnFqYnZTcVJSNmEwbUFNYW1OZUgybDZTWUVDdWVUUDZQeHc9PQ==",
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
                "value": "dGNybzFmbDQ4dnNubXNkemN2ODVxNWQycTR6NWFqZGhhOHl1M3I0Z2o5aA==",
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
            "type": "delegate",
            "attributes": [
              {
                "key": "dmFsaWRhdG9y",
                "value": "dGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQ=",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwMDAwYmFzZXRjcm8=",
                "index": true
              },
              {
                "key": "bmV3X3NoYXJlcw==",
                "value": "MTAwMDAwMDAwLjAwMDAwMDAwMDAwMDAwMDAwMA==",
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
                "value": "dGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkOQ==",
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
            "value": "NDg1NzQwNDAwMWJhc2V0Y3Jv",
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
            "value": "NDg1NzQwNDAwMWJhc2V0Y3Jv",
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
            "value": "NDg1NzQwNDAwMWJhc2V0Y3Jv",
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
            "value": "NDg1NzQwNDAwMWJhc2V0Y3Jv",
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
            "value": "NDg1NzQwNDAwMWJhc2V0Y3Jv",
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
            "value": "MC4wMDA1OTUwNDU1NDA5MjIxODA=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTIxNjE3NTU0MzI0MjMxMDk=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MzA2NTc2MDI1MDIyMTEwMDUuOTExMjkyMDM0OTcyNTI5MTYz",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDg1NzQwNDAwMQ==",
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
            "value": "NDg1NzQwNDAwMWJhc2V0Y3Jv",
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
            "value": "NDg1NzQwNDAwMWJhc2V0Y3Jv",
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
            "value": "NDg1NzQwNDAwMWJhc2V0Y3Jv",
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
            "value": "MjQyODcwMjAwLjA1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
            "value": "MjQyODcwMjAuMDA1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
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
            "value": "MjQyODcwMjAwLjA1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
            "value": "MTUzODE4MTAxLjQxMDE1MTAxMTEwMzU4NTcyNGJhc2V0Y3Jv",
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
            "value": "MTUzODE4MTAxNC4xMDE1MTAxMTEwMzU4NTcyNDViYXNldGNybw==",
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
            "value": "MTUzODE3NzkzLjE1OTkxMDAxOTEyOTk5NDE2OWJhc2V0Y3Jv",
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
            "value": "MTUzODE3NzkzMS41OTkxMDAxOTEyOTk5NDE2ODliYXNldGNybw==",
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
            "value": "MTUzODE3NDg1LjUyNDkzODk2OTMwNDk2NjcyNmJhc2V0Y3Jv",
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
            "value": "MTUzODE3NDg1NS4yNDkzODk2OTMwNDk2NjcyNjRiYXNldGNybw==",
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
    "validator_updates": [
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "pqwk90+eV9yRpfmUOYfHBfc0z17gI5LHKaqSNBuNvi4="
            }
          }
        },
        "power": "500002102"
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
        "pub_key_types": ["ed25519"]
      }
    }
  }
}
`
