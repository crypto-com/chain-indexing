package usecase_parser_test

const TX_MSG_DEPOSIT_V1_AND_START_VOTING_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "2EBC31B269629C9376C277BDB1D54533EA2A55CC8E78E90B8FD932A98BC1B66D",
      "parts": {
        "total": 1,
        "hash": "E90D70FCF27F7A013CB7639C1AB738F1E3D6A8101782B9DECCC4DE048BEA413A"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "cronos_777-1",
        "height": "5691",
        "time": "2023-01-10T01:40:17.152597Z",
        "last_block_id": {
          "hash": "3CDE331DCC13029563EE4FD95CBED1648E55BCF6F6963D17FEDC078613C4AFDD",
          "parts": {
            "total": 1,
            "hash": "4E13C3197F766026FE737F24B1D4A4AA790E5E0A4D447377074854BA8BB5EECE"
          }
        },
        "last_commit_hash": "7B8A593AFEBD316B65C51DD4393148AA7E37C97F83420D63934A8A3767AD400D",
        "data_hash": "4304196D72E80B725C13B8D5EDBA8F18746C08FECE0D59311D75E12721B835CD",
        "validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "next_validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
        "app_hash": "9AA3EDC8E236F30F3D5EE93EB687D566AE4CBF45C828AB51B17DA43F0DC2EB31",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57"
      },
      "data": {
        "txs": [
          "CmAKXgoZL2Nvc21vcy5nb3YudjEuTXNnRGVwb3NpdBJBCAUSKmNyYzE4ejZxMzhtaHZ0c3Z5cjVtYWs4Zmo4czhnNGd3N2tqanRzZ3JuNxoRCgdiYXNlY3JvEgYxMDAwMDASgAEKWQpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQJCeFp1B0RS1ipqwiJw/7j7Ack3XQunKIeugA3GGTFdGxIECgIIARgEEiMKHQoHYmFzZWNybxISMTI5NzU5MDYzNzI0MDAwMDAwEICJehpBomVYxdMPpSEcW7a//iMeZkATQpDpHHD99s1EYHj8P1VaLzcZvA0BQKmW5J/NQHEqrs5ctWU5bw2233qeygNQygA="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "5690",
        "round": 0,
        "block_id": {
          "hash": "3CDE331DCC13029563EE4FD95CBED1648E55BCF6F6963D17FEDC078613C4AFDD",
          "parts": {
            "total": 1,
            "hash": "4E13C3197F766026FE737F24B1D4A4AA790E5E0A4D447377074854BA8BB5EECE"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57",
            "timestamp": "2023-01-10T01:40:17.152597Z",
            "signature": "UpP1bssyB49OR0qEE7F5yrVHBSzH7F7WbT9bN1XqoLs1atehs34axd+aKNfnfBsyq6q/P9Q6keYt6PKuTXJMDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "FFAF83126C47B961299A37B647A112070AC4228A",
            "timestamp": "2023-01-10T01:40:17.344541Z",
            "signature": "cF/Fl+SL1XMBRiPAjpLRifAMlF1cUMpc+mW3b1gvKASHeTuCW/BdlqfvnrZJb02AEdf5dGxBY8ftt8YCCiFQCg=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_DEPOSIT_V1_AND_START_VOTING_BLOCK_RESULT_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "5691",
    "txs_results": [
      {
        "code": 0,
        "data": "EiMKIS9jb3Ntb3MuZ292LnYxLk1zZ0RlcG9zaXRSZXNwb25zZQ==",
        "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgDeposit\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"100000basecro\"},{\"key\":\"proposal_id\",\"value\":\"5\"},{\"key\":\"voting_period_start\",\"value\":\"5\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]}]}]",
        "info": "",
        "gas_wanted": "2000000",
        "gas_used": "170795",
        "events": [
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
                "index": false
              },
              {
                "key": "YW1vdW50",
                "value": "MTI5NzU5MDYzNzI0MDAwMDAwYmFzZWNybw==",
                "index": false
              }
            ]
          },
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "cmVjZWl2ZXI=",
                "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
                "index": false
              },
              {
                "key": "YW1vdW50",
                "value": "MTI5NzU5MDYzNzI0MDAwMDAwYmFzZWNybw==",
                "index": false
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
                "index": false
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
                "index": false
              },
              {
                "key": "YW1vdW50",
                "value": "MTI5NzU5MDYzNzI0MDAwMDAwYmFzZWNybw==",
                "index": false
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "ZmVl",
                "value": "MTI5NzU5MDYzNzI0MDAwMDAwYmFzZWNybw==",
                "index": false
              },
              {
                "key": "ZmVlX3BheWVy",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43LzQ=",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "b21WWXhkTVBwU0VjVzdhLy9pTWVaa0FUUXBEcEhIRDk5czFFWUhqOFAxVmFMemNadkEwQlFLbVc1Si9OUUhFcXJzNWN0V1U1YncyMjMzcWV5Z05ReWdBPQ==",
                "index": false
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2Nvc21vcy5nb3YudjEuTXNnRGVwb3NpdA==",
                "index": false
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
                "index": false
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwYmFzZWNybw==",
                "index": false
              }
            ]
          },
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "cmVjZWl2ZXI=",
                "value": "Y3JjMTBkMDd5MjY1Z21tdXZ0NHowdzlhdzg4MGpuc3I3MDBqZHVmbnlk",
                "index": false
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwYmFzZWNybw==",
                "index": false
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "Y3JjMTBkMDd5MjY1Z21tdXZ0NHowdzlhdzg4MGpuc3I3MDBqZHVmbnlk",
                "index": false
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
                "index": false
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwYmFzZWNybw==",
                "index": false
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
                "index": false
              }
            ]
          },
          {
            "type": "proposal_deposit",
            "attributes": [
              {
                "key": "YW1vdW50",
                "value": "MTAwMDAwYmFzZWNybw==",
                "index": false
              },
              {
                "key": "cHJvcG9zYWxfaWQ=",
                "value": "NQ==",
                "index": false
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "Z292ZXJuYW5jZQ==",
                "index": false
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
                "index": false
              }
            ]
          },
          {
            "type": "proposal_deposit",
            "attributes": [
              {
                "key": "dm90aW5nX3BlcmlvZF9zdGFydA==",
                "value": "NQ==",
                "index": false
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
            "value": "Y3JjMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04bDBhdzQ3",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzNzdGFrZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "coinbase",
        "attributes": [
          {
            "key": "bWludGVy",
            "value": "Y3JjMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04bDBhdzQ3",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzNzdGFrZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "Y3JjMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04bDBhdzQ3",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzNzdGFrZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzNzdGFrZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
            "index": false
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04bDBhdzQ3",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzNzdGFrZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04bDBhdzQ3",
            "index": false
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC45OTk4ODI4MzQ0OTQ0NDMxNjE=",
            "index": false
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMjk5NDIyNzU1MjM3OTI2MjQ=",
            "index": false
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MjU5OTE1MDA0MTIwNDQ0ODI3LjczMjM3Nzc2MTY1OTc1NzQ3Mg==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzM=",
            "index": false
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzNzdGFrZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "Y3JjMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4c3A3Mm1w",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzNzdGFrZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y3JjMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4c3A3Mm1w",
            "index": false
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODEwNDczNzNzdGFrZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
            "index": false
          }
        ]
      },
      {
        "type": "proposer_reward",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA1OTA1MjM2OC42NTAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
            "index": false
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JjdmFsb3BlcjE4ejZxMzhtaHZ0c3Z5cjVtYWs4Zmo4czhnNGd3N2tqanAwZTBkaA==",
            "index": false
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA1OTA1MjM2Ljg2NTAwMDAwMDAwMDAwMDAwMHN0YWtl",
            "index": false
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JjdmFsb3BlcjE4ejZxMzhtaHZ0c3Z5cjVtYWs4Zmo4czhnNGd3N2tqanAwZTBkaA==",
            "index": false
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA1OTA1MjM2OC42NTAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
            "index": false
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JjdmFsb3BlcjE4ejZxMzhtaHZ0c3Z5cjVtYWs4Zmo4czhnNGd3N2tqanAwZTBkaA==",
            "index": false
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNDkxODcwMi44NDQ1MDAwMDAwMDAwMDAwMDBzdGFrZQ==",
            "index": false
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JjdmFsb3BlcjEybHVrdTZ1eGVoaGFrMDJweTRyY3o2NXp1MHN3aDd3ajZ1bHJsZw==",
            "index": false
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNDkxODcwMjguNDQ1MDAwMDAwMDAwMDAwMDAwc3Rha2U=",
            "index": false
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JjdmFsb3BlcjEybHVrdTZ1eGVoaGFrMDJweTRyY3o2NXp1MHN3aDd3ajZ1bHJsZw==",
            "index": false
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNDkxODcwMi44NDQ1MDAwMDAwMDAwMDAwMDBzdGFrZQ==",
            "index": false
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JjdmFsb3BlcjE4ejZxMzhtaHZ0c3Z5cjVtYWs4Zmo4czhnNGd3N2tqanAwZTBkaA==",
            "index": false
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNDkxODcwMjguNDQ1MDAwMDAwMDAwMDAwMDAwc3Rha2U=",
            "index": false
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JjdmFsb3BlcjE4ejZxMzhtaHZ0c3Z5cjVtYWs4Zmo4czhnNGd3N2tqanAwZTBkaA==",
            "index": false
          }
        ]
      },
      {
        "type": "fee_market",
        "attributes": [
          {
            "key": "YmFzZV9mZWU=",
            "value": "Nw==",
            "index": false
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "block_bloom",
        "attributes": [
          {
            "key": "Ymxvb20=",
            "value": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
            "index": false
          }
        ]
      },
      {
        "type": "block_gas",
        "attributes": [
          {
            "key": "aGVpZ2h0",
            "value": "NTY5MQ==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTcwNzk1",
            "index": false
          }
        ]
      }
    ],
    "validator_updates": null,
    "consensus_param_updates": {
      "block": {
        "max_bytes": "1048576",
        "max_gas": "81500000"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000",
        "max_bytes": "1048576"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      }
    }
  }
}`

const TX_MSG_DEPOSIT_V1_AND_START_VOTING_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/cosmos.gov.v1.MsgDeposit",
          "proposal_id": "5",
          "depositor": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
          "amount": [
            {
              "denom": "basecro",
              "amount": "100000"
            }
          ]
        }
      ],
      "memo": "",
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
            "@type": "/ethermint.crypto.v1.ethsecp256k1.PubKey",
            "key": "AkJ4WnUHRFLWKmrCInD/uPsByTddC6coh66ADcYZMV0b"
          },
          "mode_info": {
            "single": {
              "mode": "SIGN_MODE_DIRECT"
            }
          },
          "sequence": "4"
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "basecro",
            "amount": "129759063724000000"
          }
        ],
        "gas_limit": "2000000",
        "payer": "",
        "granter": ""
      },
      "tip": null
    },
    "signatures": [
      "omVYxdMPpSEcW7a//iMeZkATQpDpHHD99s1EYHj8P1VaLzcZvA0BQKmW5J/NQHEqrs5ctWU5bw2233qeygNQygA="
    ]
  },
  "tx_response": {
    "height": "5691",
    "txhash": "D41222A53DDCF690879AF8A681B0EE55991EC8EA204533223A10706DD945D7AA",
    "codespace": "",
    "code": 0,
    "data": "12230A212F636F736D6F732E676F762E76312E4D73674465706F736974526573706F6E7365",
    "raw_log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgDeposit\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"100000basecro\"},{\"key\":\"proposal_id\",\"value\":\"5\"},{\"key\":\"voting_period_start\",\"value\":\"5\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "receiver",
                "value": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd"
              },
              {
                "key": "amount",
                "value": "100000basecro"
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "spender",
                "value": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"
              },
              {
                "key": "amount",
                "value": "100000basecro"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/cosmos.gov.v1.MsgDeposit"
              },
              {
                "key": "sender",
                "value": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"
              },
              {
                "key": "module",
                "value": "governance"
              },
              {
                "key": "sender",
                "value": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"
              }
            ]
          },
          {
            "type": "proposal_deposit",
            "attributes": [
              {
                "key": "amount",
                "value": "100000basecro"
              },
              {
                "key": "proposal_id",
                "value": "5"
              },
              {
                "key": "voting_period_start",
                "value": "5"
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "recipient",
                "value": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd"
              },
              {
                "key": "sender",
                "value": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"
              },
              {
                "key": "amount",
                "value": "100000basecro"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "2000000",
    "gas_used": "170795",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.gov.v1.MsgDeposit",
            "proposal_id": "5",
            "depositor": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
            "amount": [
              {
                "denom": "basecro",
                "amount": "100000"
              }
            ]
          }
        ],
        "memo": "",
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
              "@type": "/ethermint.crypto.v1.ethsecp256k1.PubKey",
              "key": "AkJ4WnUHRFLWKmrCInD/uPsByTddC6coh66ADcYZMV0b"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "4"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "basecro",
              "amount": "129759063724000000"
            }
          ],
          "gas_limit": "2000000",
          "payer": "",
          "granter": ""
        },
        "tip": null
      },
      "signatures": [
        "omVYxdMPpSEcW7a//iMeZkATQpDpHHD99s1EYHj8P1VaLzcZvA0BQKmW5J/NQHEqrs5ctWU5bw2233qeygNQygA="
      ]
    },
    "timestamp": "2023-01-10T01:40:17Z",
    "events": [
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTI5NzU5MDYzNzI0MDAwMDAwYmFzZWNybw==",
            "index": false
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTI5NzU5MDYzNzI0MDAwMDAwYmFzZWNybw==",
            "index": false
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y3JjMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsODhla2Vy",
            "index": false
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTI5NzU5MDYzNzI0MDAwMDAwYmFzZWNybw==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTI5NzU5MDYzNzI0MDAwMDAwYmFzZWNybw==",
            "index": false
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43LzQ=",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "b21WWXhkTVBwU0VjVzdhLy9pTWVaa0FUUXBEcEhIRDk5czFFWUhqOFAxVmFMemNadkEwQlFLbVc1Si9OUUhFcXJzNWN0V1U1YncyMjMzcWV5Z05ReWdBPQ==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2Nvc21vcy5nb3YudjEuTXNnRGVwb3NpdA==",
            "index": false
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwMDAwYmFzZWNybw==",
            "index": false
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "Y3JjMTBkMDd5MjY1Z21tdXZ0NHowdzlhdzg4MGpuc3I3MDBqZHVmbnlk",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwMDAwYmFzZWNybw==",
            "index": false
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y3JjMTBkMDd5MjY1Z21tdXZ0NHowdzlhdzg4MGpuc3I3MDBqZHVmbnlk",
            "index": false
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwMDAwYmFzZWNybw==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
            "index": false
          }
        ]
      },
      {
        "type": "proposal_deposit",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAwMDAwYmFzZWNybw==",
            "index": false
          },
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "NQ==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "Z292ZXJuYW5jZQ==",
            "index": false
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43",
            "index": false
          }
        ]
      },
      {
        "type": "proposal_deposit",
        "attributes": [
          {
            "key": "dm90aW5nX3BlcmlvZF9zdGFydA==",
            "value": "NQ==",
            "index": false
          }
        ]
      }
    ]
  }
}`
