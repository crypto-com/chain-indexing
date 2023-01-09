package usecase_parser_test

const TX_MSG_SOFTWARE_UPGRADE_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "80FA0532EE6BC83CD2D5982F8A041614CB5961B5EC4495AC92168AFAC8DB3BF2",
      "parts": {
        "total": 1,
        "hash": "2DBB779461448C3AD25C6CCF65B1632CDFDDE5763CAB7FF28C31D5C267517141"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "cronos_777-1",
        "height": "1634",
        "time": "2023-01-09T10:26:47.119113Z",
        "last_block_id": {
          "hash": "21460DC9C582FD354A26BA8FDE041D0FACB1A6D98EEED0C6338796EC0A0E2DE6",
          "parts": {
            "total": 1,
            "hash": "8CAF6286634D3139E62974B300959256A05459D48B0AC9DAE89EAC7034D6942C"
          }
        },
        "last_commit_hash": "AAFFA8C4A27B051CF98FBD6E8BECDB0402030AF2B41853FF3073B7C119CA5C4F",
        "data_hash": "76057176D58C7D20A80D912610AAA700893A3831231ED3BED6A25DC53FEAB914",
        "validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "next_validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
        "app_hash": "066710B94772F5028668583F9460F5253A38FCD0AADC36A6782EAF76DF101B8D",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "FFAF83126C47B961299A37B647A112070AC4228A"
      },
      "data": {
        "txs": [
          "CuwBCukBCiAvY29zbW9zLmdvdi52MS5Nc2dTdWJtaXRQcm9wb3NhbBLEAQp4CiovY29zbW9zLnVwZ3JhZGUudjFiZXRhMS5Nc2dTb2Z0d2FyZVVwZ3JhZGUSSgoqY3JjMTBkMDd5MjY1Z21tdXZ0NHowdzlhdzg4MGpuc3I3MDBqZHVmbnlkEhwKBG5hbWUSCwiAkrjDmP7///8BGOgHIgRpbmZvEhAKB2Jhc2Vjcm8SBTEwMDAwGipjcmMxMmx1a3U2dXhlaGhhazAycHk0cmN6NjV6dTBzd2g3d2pzcncwcHAiCmlwZnM6Ly9DSUQSgAEKWQpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQJucQpio0LeDtTXxFMty8u6+/GWUu1nsjfvq3DosgfvrBIECgIIARgDEiMKHQoHYmFzZWNybxISMTI5NzU5MDYzNzI0MDAwMDAwEICJehpBvf7G5GcU5tD4lDiRdrKh10dmy8cxbrz6XgnwHKs5cKMn3Tui/I2YGrHBXgu3sh0VOdmewUIYzZBcPo+lyNL56wE="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "1633",
        "round": 0,
        "block_id": {
          "hash": "21460DC9C582FD354A26BA8FDE041D0FACB1A6D98EEED0C6338796EC0A0E2DE6",
          "parts": {
            "total": 1,
            "hash": "8CAF6286634D3139E62974B300959256A05459D48B0AC9DAE89EAC7034D6942C"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57",
            "timestamp": "2023-01-09T10:26:47.119113Z",
            "signature": "XQONCrnnzpkCOdU0/lpVHG5d5nxAZTI4KgLwjUnbQA5Y2AvD4/Gm06uSM7XyWIXbhNEhQruW8pa/Tg9hczVNDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "FFAF83126C47B961299A37B647A112070AC4228A",
            "timestamp": "2023-01-09T10:26:47.21954Z",
            "signature": "N59PUc2wOyw0Jm08rtU7hKhwXMyS93qt+hp3nA52tga7Tm67dLtxsVafagDkhXt8eaumC2fvWbhONHWIpZqVCw=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_SOFTWARE_UPGRADE_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "1634",
    "txs_results": [
      {
        "code": 0,
        "data": "Ei4KKC9jb3Ntb3MuZ292LnYxLk1zZ1N1Ym1pdFByb3Bvc2FsUmVzcG9uc2USAggD",
        "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"amount\",\"value\":\"10000basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"amount\",\"value\":\"10000basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgSubmitProposal\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"10000basecro\"},{\"key\":\"proposal_id\",\"value\":\"3\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"3\"},{\"key\":\"proposal_messages\",\"value\":\",/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade\"},{\"key\":\"voting_period_start\",\"value\":\"3\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"amount\",\"value\":\"10000basecro\"}]}]}]",
        "info": "",
        "gas_wanted": "2000000",
        "gas_used": "204941",
        "events": [
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
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
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
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
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
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
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBwLzM=",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "dmY3RzVHY1U1dEQ0bERpUmRyS2gxMGRteThjeGJyejZYZ253SEtzNWNLTW4zVHVpL0kyWUdySEJYZ3Uzc2gwVk9kbWV3VUlZelpCY1BvK2x5Tkw1NndFPQ==",
                "index": false
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2Nvc21vcy5nb3YudjEuTXNnU3VibWl0UHJvcG9zYWw=",
                "index": false
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "cHJvcG9zYWxfaWQ=",
                "value": "Mw==",
                "index": false
              },
              {
                "key": "cHJvcG9zYWxfbWVzc2FnZXM=",
                "value": "LC9jb3Ntb3MudXBncmFkZS52MWJldGExLk1zZ1NvZnR3YXJlVXBncmFkZQ==",
                "index": false
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
                "index": false
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDBiYXNlY3Jv",
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
                "value": "MTAwMDBiYXNlY3Jv",
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
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
                "index": false
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMDBiYXNlY3Jv",
                "index": false
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
                "index": false
              }
            ]
          },
          {
            "type": "proposal_deposit",
            "attributes": [
              {
                "key": "YW1vdW50",
                "value": "MTAwMDBiYXNlY3Jv",
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
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
                "index": false
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "dm90aW5nX3BlcmlvZF9zdGFydA==",
                "value": "Mw==",
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
            "value": "NDExOTA2NDY5MDVzdGFrZQ==",
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
            "value": "NDExOTA2NDY5MDVzdGFrZQ==",
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
            "value": "NDExOTA2NDY5MDVzdGFrZQ==",
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
            "value": "NDExOTA2NDY5MDVzdGFrZQ==",
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
            "value": "NDExOTA2NDY5MDVzdGFrZQ==",
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
            "value": "MC45OTk5NjYzNjczODk1MzY2NDc=",
            "index": false
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMjk5ODM0MjQwNDkyODUzMzQ=",
            "index": false
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MjU5OTc1NTkxNzU2Mzc4MTg5LjcwMDgzNjM3MDcyMTM4OTE3OA==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExOTA2NDY5MDU=",
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
            "value": "NDExOTA2NDY5MDVzdGFrZQ==",
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
            "value": "NDExOTA2NDY5MDVzdGFrZQ==",
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
            "value": "NDExOTA2NDY5MDVzdGFrZQ==",
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
            "value": "MjA1OTUzMjM0NS4yNTAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MjA1OTUzMjM0LjUyNTAwMDAwMDAwMDAwMDAwMHN0YWtl",
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
            "value": "MjA1OTUzMjM0NS4yNTAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNTM2NTA4MS4wODI1MDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNTM2NTA4MTAuODI1MDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "MTkxNTM2NTA4MS4wODI1MDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNTM2NTA4MTAuODI1MDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "MTYzNA==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MjA0OTQx",
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

const TX_MSG_SOFTWARE_UPGRADE_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/cosmos.gov.v1.MsgSubmitProposal",
          "messages": [
            {
              "@type": "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
              "authority": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd",
              "plan": {
                "name": "name",
                "time": "0001-01-01T00:00:00Z",
                "height": "1000",
                "info": "info",
                "upgraded_client_state": null
              }
            }
          ],
          "initial_deposit": [
            {
              "denom": "basecro",
              "amount": "10000"
            }
          ],
          "proposer": "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
          "metadata": "ipfs://CID"
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
            "key": "Am5xCmKjQt4O1NfEUy3Ly7r78ZZS7WeyN++rcOiyB++s"
          },
          "mode_info": {
            "single": {
              "mode": "SIGN_MODE_DIRECT"
            }
          },
          "sequence": "3"
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
      "vf7G5GcU5tD4lDiRdrKh10dmy8cxbrz6XgnwHKs5cKMn3Tui/I2YGrHBXgu3sh0VOdmewUIYzZBcPo+lyNL56wE="
    ]
  },
  "tx_response": {
    "height": "1634",
    "txhash": "BFEDD454DED949E0CD349BBFD8F518AED187214A69630445CDDEBF924A48F83C",
    "codespace": "",
    "code": 0,
    "data": "122E0A282F636F736D6F732E676F762E76312E4D73675375626D697450726F706F73616C526573706F6E736512020803",
    "raw_log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"amount\",\"value\":\"10000basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"amount\",\"value\":\"10000basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgSubmitProposal\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"10000basecro\"},{\"key\":\"proposal_id\",\"value\":\"3\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"3\"},{\"key\":\"proposal_messages\",\"value\":\",/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade\"},{\"key\":\"voting_period_start\",\"value\":\"3\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"amount\",\"value\":\"10000basecro\"}]}]}]",
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
                "value": "10000basecro"
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "spender",
                "value": "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp"
              },
              {
                "key": "amount",
                "value": "10000basecro"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/cosmos.gov.v1.MsgSubmitProposal"
              },
              {
                "key": "sender",
                "value": "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp"
              },
              {
                "key": "module",
                "value": "governance"
              },
              {
                "key": "sender",
                "value": "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp"
              }
            ]
          },
          {
            "type": "proposal_deposit",
            "attributes": [
              {
                "key": "amount",
                "value": "10000basecro"
              },
              {
                "key": "proposal_id",
                "value": "3"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "3"
              },
              {
                "key": "proposal_messages",
                "value": ",/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade"
              },
              {
                "key": "voting_period_start",
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
                "value": "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp"
              },
              {
                "key": "amount",
                "value": "10000basecro"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "2000000",
    "gas_used": "204941",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.gov.v1.MsgSubmitProposal",
            "messages": [
              {
                "@type": "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
                "authority": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd",
                "plan": {
                  "name": "name",
                  "time": "0001-01-01T00:00:00Z",
                  "height": "1000",
                  "info": "info",
                  "upgraded_client_state": null
                }
              }
            ],
            "initial_deposit": [
              {
                "denom": "basecro",
                "amount": "10000"
              }
            ],
            "proposer": "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
            "metadata": "ipfs://CID"
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
              "key": "Am5xCmKjQt4O1NfEUy3Ly7r78ZZS7WeyN++rcOiyB++s"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "3"
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
        "vf7G5GcU5tD4lDiRdrKh10dmy8cxbrz6XgnwHKs5cKMn3Tui/I2YGrHBXgu3sh0VOdmewUIYzZBcPo+lyNL56wE="
      ]
    },
    "timestamp": "2023-01-09T10:26:47Z",
    "events": [
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
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
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
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
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
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
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBwLzM=",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "dmY3RzVHY1U1dEQ0bERpUmRyS2gxMGRteThjeGJyejZYZ253SEtzNWNLTW4zVHVpL0kyWUdySEJYZ3Uzc2gwVk9kbWV3VUlZelpCY1BvK2x5Tkw1NndFPQ==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2Nvc21vcy5nb3YudjEuTXNnU3VibWl0UHJvcG9zYWw=",
            "index": false
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "Mw==",
            "index": false
          },
          {
            "key": "cHJvcG9zYWxfbWVzc2FnZXM=",
            "value": "LC9jb3Ntb3MudXBncmFkZS52MWJldGExLk1zZ1NvZnR3YXJlVXBncmFkZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwMDBiYXNlY3Jv",
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
            "value": "MTAwMDBiYXNlY3Jv",
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
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwMDBiYXNlY3Jv",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
            "index": false
          }
        ]
      },
      {
        "type": "proposal_deposit",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAwMDBiYXNlY3Jv",
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
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBw",
            "index": false
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "dm90aW5nX3BlcmlvZF9zdGFydA==",
            "value": "Mw==",
            "index": false
          }
        ]
      }
    ]
  }
}`
