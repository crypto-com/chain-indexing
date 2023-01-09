package usecase_parser_test

const TX_MSG_DEPOSIT_V1_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "3CE06614F79111764FBB4AD88CC6FF6448D16EBDB86BC3EB084D8CAAF5851A0D",
      "parts": {
        "total": 1,
        "hash": "C17054B8552A39CA9C31588CD1AB15B4E1C6DFF76496426518811918820967DC"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "cronos_777-1",
        "height": "5066",
        "time": "2023-01-10T01:25:19.002708Z",
        "last_block_id": {
          "hash": "1669996F0C09C972C535C46B27023F8512F2BD1E01FA78BDDE18AC517A75D967",
          "parts": {
            "total": 1,
            "hash": "3ACD950796AC8FBE200978B8F6BEBAD0291687868596AF1F26CD3DAA66AD106D"
          }
        },
        "last_commit_hash": "B91FB539DD91BC5DB7260F531CC7E54F279502398E9FB70821A61BA76B164AE2",
        "data_hash": "CBA30AC2B1918ECD1002C1B5BC031E4643B178D87ADC69B783BF241F8568263B",
        "validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "next_validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
        "app_hash": "F1A5C4B2894785B6F5D6D891974AC1709130CFA51C73E5C8C0A9421CECE5EE0E",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "FFAF83126C47B961299A37B647A112070AC4228A"
      },
      "data": {
        "txs": [
          "CmAKXgoZL2Nvc21vcy5nb3YudjEuTXNnRGVwb3NpdBJBCAMSKmNyYzE4ejZxMzhtaHZ0c3Z5cjVtYWs4Zmo4czhnNGd3N2tqanRzZ3JuNxoRCgdiYXNlY3JvEgYxMDAwMDASgAEKWQpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQJCeFp1B0RS1ipqwiJw/7j7Ack3XQunKIeugA3GGTFdGxIECgIIARgBEiMKHQoHYmFzZWNybxISMTI5NzU5MDYzNzI0MDAwMDAwEICJehpBj5YLxrlLfz8joXuhfysNJg0qhoeXEe/HSewLiIttPnYPZf4nlv2TAIBRm3iA65u2Q1nNZAGz9TGC2KWDEV3zPQA="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "5065",
        "round": 0,
        "block_id": {
          "hash": "1669996F0C09C972C535C46B27023F8512F2BD1E01FA78BDDE18AC517A75D967",
          "parts": {
            "total": 1,
            "hash": "3ACD950796AC8FBE200978B8F6BEBAD0291687868596AF1F26CD3DAA66AD106D"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57",
            "timestamp": "2023-01-10T01:25:19.176842Z",
            "signature": "+1i7kXriQDURx/1x9dR08BM12TrT2QA1kIuu0Bre0zBcUCYsbphWUdSsQ4vtaFCqg6TLJWEHbPjalh4ZLfSoBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "FFAF83126C47B961299A37B647A112070AC4228A",
            "timestamp": "2023-01-10T01:25:19.002708Z",
            "signature": "+dHZnKOdSfo00EtmAel015aQuubP098fSxhwlzQw6cH6plnE6LMTvmhWJeXiSUlSp1Rr/0W3SlirNdO8d5T/AQ=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_DEPOSIT_V1_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "5066",
    "txs_results": [
      {
        "code": 0,
        "data": "EiMKIS9jb3Ntb3MuZ292LnYxLk1zZ0RlcG9zaXRSZXNwb25zZQ==",
        "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgDeposit\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"100000basecro\"},{\"key\":\"proposal_id\",\"value\":\"3\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]}]}]",
        "info": "",
        "gas_wanted": "2000000",
        "gas_used": "155908",
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
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43LzE=",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "ajVZTHhybExmejhqb1h1aGZ5c05KZzBxaG9lWEVlL0hTZXdMaUl0dFBuWVBaZjRubHYyVEFJQlJtM2lBNjV1MlExbk5aQUd6OVRHQzJLV0RFVjN6UFFBPQ==",
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
                "value": "Mw==",
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
            "value": "NDExODI1MjYyMDZzdGFrZQ==",
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
            "value": "NDExODI1MjYyMDZzdGFrZQ==",
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
            "value": "NDExODI1MjYyMDZzdGFrZQ==",
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
            "value": "NDExODI1MjYyMDZzdGFrZQ==",
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
            "value": "NDExODI1MjYyMDZzdGFrZQ==",
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
            "value": "MC45OTk4OTU3MDA5NTMyNDU5MjI=",
            "index": false
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMjk5NDg2MTM5Njk1NDAzNjE=",
            "index": false
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MjU5OTI0MzM3Nzk5NzQxNDE5Ljk3OTI0OTk1ODQyMjA2MDY1MQ==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODI1MjYyMDY=",
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
            "value": "NDExODI1MjYyMDZzdGFrZQ==",
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
            "value": "NDExODI1MjYyMDZzdGFrZQ==",
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
            "value": "NDExODI1MjYyMDZzdGFrZQ==",
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
            "value": "MjA1OTEyNjMxMC4zMDAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MjA1OTEyNjMxLjAzMDAwMDAwMDAwMDAwMDAwMHN0YWtl",
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
            "value": "MjA1OTEyNjMxMC4zMDAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk4NzQ2OC41NzkwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk4NzQ2ODUuNzkwMDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "MTkxNDk4NzQ2OC41NzkwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk4NzQ2ODUuNzkwMDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "NTA2Ng==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTU1OTA4",
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

const TX_MSG_DEPOSIT_V1_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/cosmos.gov.v1.MsgDeposit",
          "proposal_id": "3",
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
          "sequence": "1"
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
      "j5YLxrlLfz8joXuhfysNJg0qhoeXEe/HSewLiIttPnYPZf4nlv2TAIBRm3iA65u2Q1nNZAGz9TGC2KWDEV3zPQA="
    ]
  },
  "tx_response": {
    "height": "5066",
    "txhash": "8F549DCBED33FF8909C83C6FE957D262F833962ACFE70C53E969DE8F1D657B7C",
    "codespace": "",
    "code": 0,
    "data": "12230A212F636F736D6F732E676F762E76312E4D73674465706F736974526573706F6E7365",
    "raw_log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgDeposit\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"100000basecro\"},{\"key\":\"proposal_id\",\"value\":\"3\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"},{\"key\":\"amount\",\"value\":\"100000basecro\"}]}]}]",
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
                "value": "3"
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
    "gas_used": "155908",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.gov.v1.MsgDeposit",
            "proposal_id": "3",
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
            "sequence": "1"
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
        "j5YLxrlLfz8joXuhfysNJg0qhoeXEe/HSewLiIttPnYPZf4nlv2TAIBRm3iA65u2Q1nNZAGz9TGC2KWDEV3zPQA="
      ]
    },
    "timestamp": "2023-01-10T01:25:19Z",
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
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43LzE=",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "ajVZTHhybExmejhqb1h1aGZ5c05KZzBxaG9lWEVlL0hTZXdMaUl0dFBuWVBaZjRubHYyVEFJQlJtM2lBNjV1MlExbk5aQUd6OVRHQzJLV0RFVjN6UFFBPQ==",
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
            "value": "Mw==",
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
      }
    ]
  }
}`
