package usecase_parser_test

const TX_MSG_CREATE_SOLOMACHINE_CLIENT_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "B147388BB1E0D1FE5ABDD5820D1137C887B8FCC0C6A648799259C31867E8B816",
      "parts": {
        "total": 1,
        "hash": "4DEAB75CA344C466ADEEEEBC9180BB5C261BC4FB8893F8182576990FF927F2CE"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-4",
        "height": "14791",
        "time": "2021-08-20T02:49:21.578863734Z",
        "last_block_id": {
          "hash": "96068C0A0D2EC7EE28F1C3D30707623D0E6AB4EDD16382402796336208D19AEB",
          "parts": {
            "total": 1,
            "hash": "EA8F35958282FA026C0F54297C61616CC7BACCF6D7B0B8794F4AE50D5D66D803"
          }
        },
        "last_commit_hash": "0A15AEB9A5045C595C898D0C3E13909AEC2BC7CEB0DF3BA1D44F2B6F54E0BFA8",
        "data_hash": "337DAE7557BDFF199D98449FEBDEC21832EE656B378C89FA7A59F8B937D70179",
        "validators_hash": "856C3B06E9040E535C3E1D46B2863876FAF1F001E7B7D1C70C0BA571A737E0C9",
        "next_validators_hash": "856C3B06E9040E535C3E1D46B2863876FAF1F001E7B7D1C70C0BA571A737E0C9",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "69B504EF011137526AD994981108C522BBFBD9071D59BCE35A013CB8C67273B6",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2"
      },
      "data": {
        "txs": [
          "CqoDCpQDCiMvaWJjLmNvcmUuY2xpZW50LnYxLk1zZ0NyZWF0ZUNsaWVudBLsAgqeAQosL2liYy5saWdodGNsaWVudHMuc29sb21hY2hpbmUudjIuQ2xpZW50U3RhdGUSbggBGmgKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEC6yndcJ14UBZcc15MFvclljnp8ydxn+2lGcZ9IVV0RKYSGHNvbG8tbWFjaGluZS1kaXZlcnNpZmllchiqsPyIBiABEpsBCi8vaWJjLmxpZ2h0Y2xpZW50cy5zb2xvbWFjaGluZS52Mi5Db25zZW5zdXNTdGF0ZRJoCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAusp3XCdeFAWXHNeTBb3JZY56fMncZ/tpRnGfSFVdESmEhhzb2xvLW1hY2hpbmUtZGl2ZXJzaWZpZXIYqrD8iAYaK3Rjcm8xNHdrdTRocjc0bTBtNHR2ZXhzNGY2anZ1eTZ2bnUyeDJkZzdoc3kSEXNvbG8tbWFjaGluZS1tZW1vEmkKTgpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQLrKd1wnXhQFlxzXkwW9yWWOenzJ3Gf7aUZxn0hVXREphIECgIIARIXChEKCGJhc2V0Y3JvEgUxMDAwMBDgpxIaQKUFo8cobjRAe3tvtvj4OkkNoiKgduqPfb1kuMJQAQ8zU8589A4TEXUBFVM6On0XzS0/S43boiY52lUSp20EvV8="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "14790",
        "round": 0,
        "block_id": {
          "hash": "96068C0A0D2EC7EE28F1C3D30707623D0E6AB4EDD16382402796336208D19AEB",
          "parts": {
            "total": 1,
            "hash": "EA8F35958282FA026C0F54297C61616CC7BACCF6D7B0B8794F4AE50D5D66D803"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2",
            "timestamp": "2021-08-20T02:49:21.611930663Z",
            "signature": "A7QiugCxXqMEZzu7Dk5baylyFQP5LKN7ynBV43JQO62GxZKpGAceSTpGSOM1CN1W2LNrAFcvdKN4FBVCHKDECw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7E2B455523A35AE8F4064C00AF01515FEF673079",
            "timestamp": "2021-08-20T02:49:21.578863734Z",
            "signature": "yeHvvZeSrQgsxmLotELOBtpRUjWhVqUwv8VmDoXbwy66d6Wye34Z5fPPhy+wiEAyh0CC1AKUsmEn4X63b+/yCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F9A4582896E5F2692BEDDD7C251755E33D2DE013",
            "timestamp": "2021-08-20T02:49:21.573351718Z",
            "signature": "KaelkJqSKi+oHdhEz438ongc94LRgllC5wZ79Bdx3gThbK02CY0psFf9F499j1L0IjOG8lSgdWJJa850y0C9AQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "19982846C26E6C39745F85FFCEA9950895563EE1",
            "timestamp": "2021-08-20T02:49:21.736717Z",
            "signature": "39uboeTmZrn2ljkS5dbKGuvYcc/OGkT7PElpb4chyKQGe9s5fl4GaZvH/CdKwYEAzS2vDd+ieyEPxWnh6VGzDQ=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_CREATE_SOLOMACHINE_CLIENT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "14791",
    "txs_results": [
      {
        "code": 0,
        "data": "CiUKIy9pYmMuY29yZS5jbGllbnQudjEuTXNnQ3JlYXRlQ2xpZW50",
        "log": "[{\"events\":[{\"type\":\"create_client\",\"attributes\":[{\"key\":\"client_id\",\"value\":\"06-solomachine-0\"},{\"key\":\"client_type\",\"value\":\"06-solomachine\"},{\"key\":\"consensus_height\",\"value\":\"0-1\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/ibc.core.client.v1.MsgCreateClient\"},{\"key\":\"module\",\"value\":\"ibc_client\"}]}]}]",
        "info": "",
        "gas_wanted": "300000",
        "gas_used": "76405",
        "events": [
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "dGNybzE0d2t1NGhyNzRtMG00dHZleHM0ZjZqdnV5NnZudTJ4MmRnN2hzeQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDBiYXNldGNybw==",
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
                "value": "MTAwMDBiYXNldGNybw==",
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
                "value": "dGNybzE0d2t1NGhyNzRtMG00dHZleHM0ZjZqdnV5NnZudTJ4MmRnN2hzeQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDBiYXNldGNybw==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNybzE0d2t1NGhyNzRtMG00dHZleHM0ZjZqdnV5NnZudTJ4MmRnN2hzeQ==",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "dGNybzE0d2t1NGhyNzRtMG00dHZleHM0ZjZqdnV5NnZudTJ4MmRnN2hzeS8w",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "cFFXanh5aHVORUI3ZTIrMitQZzZTUTJpSXFCMjZvOTl2V1M0d2xBQkR6TlR6bnowRGhNUmRRRVZVem82ZlJmTkxUOUxqZHVpSmpuYVZSS25iUVM5WHc9PQ==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2liYy5jb3JlLmNsaWVudC52MS5Nc2dDcmVhdGVDbGllbnQ=",
                "index": true
              }
            ]
          },
          {
            "type": "create_client",
            "attributes": [
              {
                "key": "Y2xpZW50X2lk",
                "value": "MDYtc29sb21hY2hpbmUtMA==",
                "index": true
              },
              {
                "key": "Y2xpZW50X3R5cGU=",
                "value": "MDYtc29sb21hY2hpbmU=",
                "index": true
              },
              {
                "key": "Y29uc2Vuc3VzX2hlaWdodA==",
                "value": "MC0x",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "aWJjX2NsaWVudA==",
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
            "value": "NDc5Njk3NjQ0M2Jhc2V0Y3Jv",
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
            "value": "NDc5Njk3NjQ0M2Jhc2V0Y3Jv",
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
            "value": "NDc5Njk3NjQ0M2Jhc2V0Y3Jv",
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
            "value": "NDc5Njk3NjQ0M2Jhc2V0Y3Jv",
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
            "value": "NDc5Njk3NjQ0M2Jhc2V0Y3Jv",
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
            "value": "MC4wMDA1OTUyNjEwMjkwMTA2MDM=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTIwMTQwMzIxMDY2NTE3MjY=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MzAyNzYyMTI3NjE4OTI3NjUuMDY2ODg3NTA5NjMyNjI5ODQ4",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDc5Njk3NjQ0Mw==",
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
            "value": "NDc5Njk3NjQ0M2Jhc2V0Y3Jv",
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
            "value": "NDc5Njk3NjQ0M2Jhc2V0Y3Jv",
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
            "value": "NDc5Njk3NjQ0M2Jhc2V0Y3Jv",
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
            "value": "MjM5ODQ4ODIyLjE1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
            "value": "MjM5ODQ4ODIuMjE1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
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
            "value": "MjM5ODQ4ODIyLjE1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
            "value": "MTUxODk0MTI3Ljc1MzE0OTc4OTY3NDc1ODcyMmJhc2V0Y3Jv",
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
            "value": "MTUxODk0MTI3Ny41MzE0OTc4OTY3NDc1ODcyMThiYXNldGNybw==",
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
            "value": "MTUxODk0MTI3Ljc1MzE0OTc4OTY3NDc1ODcyMmJhc2V0Y3Jv",
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
            "value": "MTUxODk0MTI3Ny41MzE0OTc4OTY3NDc1ODcyMThiYXNldGNybw==",
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
            "value": "MTUxODk0MTI3Ljc1MzE0OTc4OTY3NDc1ODcyMmJhc2V0Y3Jv",
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
            "value": "MTUxODk0MTI3Ny41MzE0OTc4OTY3NDc1ODcyMThiYXNldGNybw==",
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
            "value": "MzAzNzguODI1NTUwNjI5NjA4NTg1NTQ4YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxazRuejJ5YWxzdXp1enhleGN4dXVnaHAzM2NmcWxzN2Uwcmd5bTk=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzAzNzg4LjI1NTUwNjI5NjA4NTg1NTQ4M2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxazRuejJ5YWxzdXp1enhleGN4dXVnaHAzM2NmcWxzN2Uwcmd5bTk=",
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

const TX_MSG_CREATE_SOLOMACHINE_CLIENT_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/ibc.core.client.v1.MsgCreateClient",
          "client_state": {
            "@type": "/ibc.lightclients.solomachine.v2.ClientState",
            "sequence": "1",
            "is_frozen": false,
            "consensus_state": {
              "public_key": {
                "@type": "/cosmos.crypto.secp256k1.PubKey",
                "key": "Ausp3XCdeFAWXHNeTBb3JZY56fMncZ/tpRnGfSFVdESm"
              },
              "diversifier": "solo-machine-diversifier",
              "timestamp": "1629427754"
            },
            "allow_update_after_proposal": true
          },
          "consensus_state": {
            "@type": "/ibc.lightclients.solomachine.v2.ConsensusState",
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "Ausp3XCdeFAWXHNeTBb3JZY56fMncZ/tpRnGfSFVdESm"
            },
            "diversifier": "solo-machine-diversifier",
            "timestamp": "1629427754"
          },
          "signer": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy"
        }
      ],
      "memo": "solo-machine-memo",
      "timeout_height": "0",
      "extension_options": [
      ],
      "non_critical_extension_options": [
      ]
    },
    "auth_info": {
      "signer_infos": [
        {
          "public_key": {
            "@type": "/cosmos.crypto.secp256k1.PubKey",
            "key": "Ausp3XCdeFAWXHNeTBb3JZY56fMncZ/tpRnGfSFVdESm"
          },
          "mode_info": {
            "single": {
              "mode": "SIGN_MODE_DIRECT"
            }
          },
          "sequence": "0"
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "basetcro",
            "amount": "10000"
          }
        ],
        "gas_limit": "300000",
        "payer": "",
        "granter": ""
      }
    },
    "signatures": [
      "pQWjxyhuNEB7e2+2+Pg6SQ2iIqB26o99vWS4wlABDzNTznz0DhMRdQEVUzo6fRfNLT9LjduiJjnaVRKnbQS9Xw=="
    ]
  },
  "tx_response": {
    "height": "14791",
    "txhash": "0DB632201805BC399035F0B0CD0E1DABC061E5209D8A709EC0A2B29B1A306BA5",
    "codespace": "",
    "code": 0,
    "data": "0A250A232F6962632E636F72652E636C69656E742E76312E4D7367437265617465436C69656E74",
    "raw_log": "[{\"events\":[{\"type\":\"create_client\",\"attributes\":[{\"key\":\"client_id\",\"value\":\"06-solomachine-0\"},{\"key\":\"client_type\",\"value\":\"06-solomachine\"},{\"key\":\"consensus_height\",\"value\":\"0-1\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/ibc.core.client.v1.MsgCreateClient\"},{\"key\":\"module\",\"value\":\"ibc_client\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "create_client",
            "attributes": [
              {
                "key": "client_id",
                "value": "06-solomachine-0"
              },
              {
                "key": "client_type",
                "value": "06-solomachine"
              },
              {
                "key": "consensus_height",
                "value": "0-1"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/ibc.core.client.v1.MsgCreateClient"
              },
              {
                "key": "module",
                "value": "ibc_client"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "300000",
    "gas_used": "76405",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/ibc.core.client.v1.MsgCreateClient",
            "client_state": {
              "@type": "/ibc.lightclients.solomachine.v2.ClientState",
              "sequence": "1",
              "is_frozen": false,
              "consensus_state": {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "Ausp3XCdeFAWXHNeTBb3JZY56fMncZ/tpRnGfSFVdESm"
                },
                "diversifier": "solo-machine-diversifier",
                "timestamp": "1629427754"
              },
              "allow_update_after_proposal": true
            },
            "consensus_state": {
              "@type": "/ibc.lightclients.solomachine.v2.ConsensusState",
              "public_key": {
                "@type": "/cosmos.crypto.secp256k1.PubKey",
                "key": "Ausp3XCdeFAWXHNeTBb3JZY56fMncZ/tpRnGfSFVdESm"
              },
              "diversifier": "solo-machine-diversifier",
              "timestamp": "1629427754"
            },
            "signer": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy"
          }
        ],
        "memo": "solo-machine-memo",
        "timeout_height": "0",
        "extension_options": [
        ],
        "non_critical_extension_options": [
        ]
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "Ausp3XCdeFAWXHNeTBb3JZY56fMncZ/tpRnGfSFVdESm"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "0"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "basetcro",
              "amount": "10000"
            }
          ],
          "gas_limit": "300000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "pQWjxyhuNEB7e2+2+Pg6SQ2iIqB26o99vWS4wlABDzNTznz0DhMRdQEVUzo6fRfNLT9LjduiJjnaVRKnbQS9Xw=="
      ]
    },
    "timestamp": "2021-08-20T02:49:21Z"
  }
}`
