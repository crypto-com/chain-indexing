package usecase_parser_test

const TX_MSG_EXEC_LEGACY_CONTENT_V1_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "3302C826D11303C99E8015232A5E18C9B3349443AE1F89F39A5A9DEA713F064B",
      "parts": {
        "total": 1,
        "hash": "C771432140AA19318835796A5AF5280EFFE9163D476F998B57EE6C91F634651F"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "cronos_777-1",
        "height": "6580",
        "time": "2023-01-10T02:01:24.977717Z",
        "last_block_id": {
          "hash": "DB8C9E400FFEC2C9F7F222D5F1D2FB704E8E1AF72B50D77A981166C53F61BFAB",
          "parts": {
            "total": 1,
            "hash": "8E7AE8FCC15231A674F3FFF90BFD6C1C18A34C954855B26B53F0ADA01392EEA1"
          }
        },
        "last_commit_hash": "562FDDDBD3131DA773D47322F9CDE24E643313145836CA12E435BE55B332BC13",
        "data_hash": "CD3D1CCC757104ED90558189D57B7271A3C63A7AC4E3817EBE64FDD83E3AEE19",
        "validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "next_validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
        "app_hash": "FC126015763619834BA869255D81613D9C23472591EC28F1E46F2336B853FCD6",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "FFAF83126C47B961299A37B647A112070AC4228A"
      },
      "data": {
        "txs": [
          "CrICCq8CCiAvY29zbW9zLmdvdi52MS5Nc2dTdWJtaXRQcm9wb3NhbBKKAgq7AQojL2Nvc21vcy5nb3YudjEuTXNnRXhlY0xlZ2FjeUNvbnRlbnQSkwEKZQovL2Nvc21vcy51cGdyYWRlLnYxYmV0YTEuU29mdHdhcmVVcGdyYWRlUHJvcG9zYWwSMgoFdGl0bGUSC2Rlc2NyaXB0aW9uGhwKBHRlc3QSCwiAkrjDmP7///8BGJBOIgRpbmZvEipjcmMxMGQwN3kyNjVnbW11dnQ0ejB3OWF3ODgwam5zcjcwMGpkdWZueWQSEgoHYmFzZWNybxIHMTAwMDAwMBoqY3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBwIgppcGZzOi8vQ0lEEoABClkKTwooL2V0aGVybWludC5jcnlwdG8udjEuZXRoc2VjcDI1NmsxLlB1YktleRIjCiECbnEKYqNC3g7U18RTLcvLuvvxllLtZ7I376tw6LIH76wSBAoCCAEYBhIjCh0KB2Jhc2Vjcm8SEjEyOTc1OTA2MzcyNDAwMDAwMBCAiXoaQdf5VlIcGN26iWNyye7I8qpe3FPiNWnT5+tvnYuZjzzCeZM06fU2KlYEKIcBdPhLC2K7DCfge29X+HsRKtDLpHEA"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "6579",
        "round": 0,
        "block_id": {
          "hash": "DB8C9E400FFEC2C9F7F222D5F1D2FB704E8E1AF72B50D77A981166C53F61BFAB",
          "parts": {
            "total": 1,
            "hash": "8E7AE8FCC15231A674F3FFF90BFD6C1C18A34C954855B26B53F0ADA01392EEA1"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57",
            "timestamp": "2023-01-10T02:01:25.05978Z",
            "signature": "9YBOZlggcj9bGBdnhirsLCfV+k4RN4AozEaA/H1pCLGyA48c3mzUnpmy19PWI/VLLKVmPv7jKolaxVA9cUIcDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "FFAF83126C47B961299A37B647A112070AC4228A",
            "timestamp": "2023-01-10T02:01:24.977717Z",
            "signature": "SpGvPB7pfFtBljG6P72ROIBdY/R9T6CXcI/68wzaO9EbZVekyzw/fbSlhtMDO6UMfgWIHYHvatzAs/EiLTBdAA=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_EXEC_LEGACY_CONTENT_V1_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "6580",
    "txs_results": [
      {
        "code": 0,
        "data": "Ei4KKC9jb3Ntb3MuZ292LnYxLk1zZ1N1Ym1pdFByb3Bvc2FsUmVzcG9uc2USAggG",
        "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"amount\",\"value\":\"1000000basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"amount\",\"value\":\"1000000basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgSubmitProposal\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"1000000basecro\"},{\"key\":\"proposal_id\",\"value\":\"6\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"6\"},{\"key\":\"proposal_messages\",\"value\":\",/cosmos.gov.v1.MsgExecLegacyContent\"},{\"key\":\"voting_period_start\",\"value\":\"6\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"amount\",\"value\":\"1000000basecro\"}]}]}]",
        "info": "",
        "gas_wanted": "2000000",
        "gas_used": "223954",
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
                "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBwLzY=",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "MS9sV1Vod1kzYnFKWTNMSjdzanlxbDdjVStJMWFkUG42MitkaTVtUFBNSjVrelRwOVRZcVZnUW9od0YwK0VzTFlyc01KK0I3YjFmNGV4RXEwTXVrY1FBPQ==",
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
                "value": "Ng==",
                "index": false
              },
              {
                "key": "cHJvcG9zYWxfbWVzc2FnZXM=",
                "value": "LC9jb3Ntb3MuZ292LnYxLk1zZ0V4ZWNMZWdhY3lDb250ZW50",
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
                "value": "MTAwMDAwMGJhc2Vjcm8=",
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
                "value": "MTAwMDAwMGJhc2Vjcm8=",
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
                "value": "MTAwMDAwMGJhc2Vjcm8=",
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
                "value": "MTAwMDAwMGJhc2Vjcm8=",
                "index": false
              },
              {
                "key": "cHJvcG9zYWxfaWQ=",
                "value": "Ng==",
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
                "value": "Ng==",
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
            "value": "NDExNzg5NDM4OTVzdGFrZQ==",
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
            "value": "NDExNzg5NDM4OTVzdGFrZQ==",
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
            "value": "NDExNzg5NDM4OTVzdGFrZQ==",
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
            "value": "NDExNzg5NDM4OTVzdGFrZQ==",
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
            "value": "NDExNzg5NDM4OTVzdGFrZQ==",
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
            "value": "MC45OTk4NjQ1MzQ2MDk4NTIyNzA=",
            "index": false
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMjk5MzMyNjAxNDQ0NDc3MzA=",
            "index": false
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MjU5OTAxNzI3OTc3ODc2MDUyLjEyMDUzNjk1OTcxNTQ1MjY1MA==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExNzg5NDM4OTU=",
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
            "value": "NDExNzg5NDM4OTVzdGFrZQ==",
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
            "value": "NDExNzg5NDM4OTVzdGFrZQ==",
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
            "value": "NDExNzg5NDM4OTVzdGFrZQ==",
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
            "value": "MjA1ODk0NzE5NC43NTAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MjA1ODk0NzE5LjQ3NTAwMDAwMDAwMDAwMDAwMHN0YWtl",
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
            "value": "MjA1ODk0NzE5NC43NTAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDgyMDg5MS4xMTc1MDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDgyMDg5MTEuMTc1MDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "MTkxNDgyMDg5MS4xMTc1MDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDgyMDg5MTEuMTc1MDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "NjU4MA==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MjIzOTU0",
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

const TX_MSG_EXEC_LEGACY_CONTENT_V1_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/cosmos.gov.v1.MsgSubmitProposal",
          "messages": [
            {
              "@type": "/cosmos.gov.v1.MsgExecLegacyContent",
              "content": {
                "@type": "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
                "title": "title",
                "description": "description",
                "plan": {
                  "name": "test",
                  "time": "0001-01-01T00:00:00Z",
                  "height": "10000",
                  "info": "info",
                  "upgraded_client_state": null
                }
              },
              "authority": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd"
            }
          ],
          "initial_deposit": [
            {
              "denom": "basecro",
              "amount": "1000000"
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
          "sequence": "6"
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
      "1/lWUhwY3bqJY3LJ7sjyql7cU+I1adPn62+di5mPPMJ5kzTp9TYqVgQohwF0+EsLYrsMJ+B7b1f4exEq0MukcQA="
    ]
  },
  "tx_response": {
    "height": "6580",
    "txhash": "DBAFA8C7C7F3A39C8162E2463E560822554A89A79DAB550882270125902AF39C",
    "codespace": "",
    "code": 0,
    "data": "122E0A282F636F736D6F732E676F762E76312E4D73675375626D697450726F706F73616C526573706F6E736512020806",
    "raw_log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"amount\",\"value\":\"1000000basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"amount\",\"value\":\"1000000basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgSubmitProposal\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"1000000basecro\"},{\"key\":\"proposal_id\",\"value\":\"6\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"6\"},{\"key\":\"proposal_messages\",\"value\":\",/cosmos.gov.v1.MsgExecLegacyContent\"},{\"key\":\"voting_period_start\",\"value\":\"6\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd\"},{\"key\":\"sender\",\"value\":\"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp\"},{\"key\":\"amount\",\"value\":\"1000000basecro\"}]}]}]",
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
                "value": "1000000basecro"
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
                "value": "1000000basecro"
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
                "value": "1000000basecro"
              },
              {
                "key": "proposal_id",
                "value": "6"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "6"
              },
              {
                "key": "proposal_messages",
                "value": ",/cosmos.gov.v1.MsgExecLegacyContent"
              },
              {
                "key": "voting_period_start",
                "value": "6"
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
                "value": "1000000basecro"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "2000000",
    "gas_used": "223954",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.gov.v1.MsgSubmitProposal",
            "messages": [
              {
                "@type": "/cosmos.gov.v1.MsgExecLegacyContent",
                "content": {
                  "@type": "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
                  "title": "title",
                  "description": "description",
                  "plan": {
                    "name": "test",
                    "time": "0001-01-01T00:00:00Z",
                    "height": "10000",
                    "info": "info",
                    "upgraded_client_state": null
                  }
                },
                "authority": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd"
              }
            ],
            "initial_deposit": [
              {
                "denom": "basecro",
                "amount": "1000000"
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
            "sequence": "6"
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
        "1/lWUhwY3bqJY3LJ7sjyql7cU+I1adPn62+di5mPPMJ5kzTp9TYqVgQohwF0+EsLYrsMJ+B7b1f4exEq0MukcQA="
      ]
    },
    "timestamp": "2023-01-10T02:01:24Z",
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
            "value": "Y3JjMTJsdWt1NnV4ZWhoYWswMnB5NHJjejY1enUwc3doN3dqc3J3MHBwLzY=",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "MS9sV1Vod1kzYnFKWTNMSjdzanlxbDdjVStJMWFkUG42MitkaTVtUFBNSjVrelRwOVRZcVZnUW9od0YwK0VzTFlyc01KK0I3YjFmNGV4RXEwTXVrY1FBPQ==",
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
            "value": "Ng==",
            "index": false
          },
          {
            "key": "cHJvcG9zYWxfbWVzc2FnZXM=",
            "value": "LC9jb3Ntb3MuZ292LnYxLk1zZ0V4ZWNMZWdhY3lDb250ZW50",
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
            "value": "MTAwMDAwMGJhc2Vjcm8=",
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
            "value": "MTAwMDAwMGJhc2Vjcm8=",
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
            "value": "MTAwMDAwMGJhc2Vjcm8=",
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
            "value": "MTAwMDAwMGJhc2Vjcm8=",
            "index": false
          },
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "Ng==",
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
            "value": "Ng==",
            "index": false
          }
        ]
      }
    ]
  }
}`
