package usecase_parser_test

const TX_MSG_VOTE_V1_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "E59B6DF081C72D54321EED4434493A2B430E289A1AA5E8C778713557E11298A5",
      "parts": {
        "total": 1,
        "hash": "877FA94EAE325E00911617C7872FA12BAA524A975935FE53344021E372EC67E8"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "cronos_777-1",
        "height": "5183",
        "time": "2023-01-10T01:28:09.630798Z",
        "last_block_id": {
          "hash": "B1CCFFDC3AE025E792C75B6E7CED2EFF5E07C4030E4FADF146217CC0F18F798A",
          "parts": {
            "total": 1,
            "hash": "BD5A4145D6659E005E6185F887A4958E9A969DA474BE609B072C4C67AEF81222"
          }
        },
        "last_commit_hash": "DA7044A81738691B3A13D92DC7A7E9529736866804EF43D10CDC61C167757496",
        "data_hash": "0D3ADDA53466B5AF4F0BD7F1C15D0A736899C8CCBEBD6255790FBD9BC07FEA2D",
        "validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "next_validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
        "app_hash": "9794111060A1EB15780166CB2ED8565B7D371FE2624162A14926630FB6C7722E",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57"
      },
      "data": {
        "txs": [
          "CkwKSgoWL2Nvc21vcy5nb3YudjEuTXNnVm90ZRIwCAMSKmNyYzE4ejZxMzhtaHZ0c3Z5cjVtYWs4Zmo4czhnNGd3N2tqanRzZ3JuNxgDEoABClkKTwooL2V0aGVybWludC5jcnlwdG8udjEuZXRoc2VjcDI1NmsxLlB1YktleRIjCiECQnhadQdEUtYqasIicP+4+wHJN10LpyiHroANxhkxXRsSBAoCCAEYAhIjCh0KB2Jhc2Vjcm8SEjEyOTc1OTA2MzcyNDAwMDAwMBCAiXoaQVj+zgRkRLSKYdB/obVDzHQA1ZtSPmjTEQO7ngCNkMZETm3MqplCW9XdM2lAQL+lOPyTl+0HkiikdA0Rc9qfLZ4A"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "5182",
        "round": 0,
        "block_id": {
          "hash": "B1CCFFDC3AE025E792C75B6E7CED2EFF5E07C4030E4FADF146217CC0F18F798A",
          "parts": {
            "total": 1,
            "hash": "BD5A4145D6659E005E6185F887A4958E9A969DA474BE609B072C4C67AEF81222"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57",
            "timestamp": "2023-01-10T01:28:09.630798Z",
            "signature": "ofYd84VyNZ/RicwTfVF6+/7IakDMr718b5bU7D5UabHCRiD4x3hEjM13ozrQdntHHNEjpPMV4C65LNx8P1NQDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "FFAF83126C47B961299A37B647A112070AC4228A",
            "timestamp": "2023-01-10T01:28:09.825914Z",
            "signature": "ZCDN4f0tvKSl/qts9kR3ZnfJO+c3f4DRymuE6KyiW+Ysb7l+2UXjIGsI0yVonmcWmsurilno331Zl47AEuhpBg=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_VOTE_V1_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "5183",
    "txs_results": [
      {
        "code": 0,
        "data": "EiAKHi9jb3Ntb3MuZ292LnYxLk1zZ1ZvdGVSZXNwb25zZQ==",
        "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgVote\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"}]},{\"type\":\"proposal_vote\",\"attributes\":[{\"key\":\"option\",\"value\":\"option:VOTE_OPTION_NO weight:\\\"1.000000000000000000\\\"\"},{\"key\":\"proposal_id\",\"value\":\"3\"}]}]}]",
        "info": "",
        "gas_wanted": "2000000",
        "gas_used": "130869",
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
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43LzI=",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "V1A3T0JHUkV0SXBoMEgraHRVUE1kQURWbTFJK2FOTVJBN3VlQUkyUXhrUk9iY3lxbVVKYjFkMHphVUJBdjZVNC9KT1g3UWVTS0tSMERSRnoycDh0bmdBPQ==",
                "index": false
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2Nvc21vcy5nb3YudjEuTXNnVm90ZQ==",
                "index": false
              }
            ]
          },
          {
            "type": "proposal_vote",
            "attributes": [
              {
                "key": "b3B0aW9u",
                "value": "b3B0aW9uOlZPVEVfT1BUSU9OX05PIHdlaWdodDoiMS4wMDAwMDAwMDAwMDAwMDAwMDAi",
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
            "value": "NDExODIyNDkzNjhzdGFrZQ==",
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
            "value": "NDExODIyNDkzNjhzdGFrZQ==",
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
            "value": "NDExODIyNDkzNjhzdGFrZQ==",
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
            "value": "NDExODIyNDkzNjhzdGFrZQ==",
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
            "value": "NDExODIyNDkzNjhzdGFrZQ==",
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
            "value": "MC45OTk4OTMyOTIyOTE4MTU0ODE=",
            "index": false
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMjk5NDc0MjczOTM2ODg2NzI=",
            "index": false
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MjU5OTIyNTkwNTMxMzE4MzI0LjA4MTk3MjQ1NTExODkxODQzMg==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODIyNDkzNjg=",
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
            "value": "NDExODIyNDkzNjhzdGFrZQ==",
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
            "value": "NDExODIyNDkzNjhzdGFrZQ==",
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
            "value": "NDExODIyNDkzNjhzdGFrZQ==",
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
            "value": "MjA1OTExMjQ2OC40MDAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MjA1OTExMjQ2Ljg0MDAwMDAwMDAwMDAwMDAwMHN0YWtl",
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
            "value": "MjA1OTExMjQ2OC40MDAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk3NDU5NS42MTIwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk3NDU5NTYuMTIwMDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "MTkxNDk3NDU5NS42MTIwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk3NDU5NTYuMTIwMDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "NTE4Mw==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTMwODY5",
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

const TX_MSG_VOTE_V1_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/cosmos.gov.v1.MsgVote",
          "proposal_id": "3",
          "voter": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
          "option": "VOTE_OPTION_NO",
          "metadata": ""
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
          "sequence": "2"
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
      "WP7OBGREtIph0H+htUPMdADVm1I+aNMRA7ueAI2QxkRObcyqmUJb1d0zaUBAv6U4/JOX7QeSKKR0DRFz2p8tngA="
    ]
  },
  "tx_response": {
    "height": "5183",
    "txhash": "D2711F0542407D7D7F4A2E34184D122D68E8E7E207E329E4354F96171793B16F",
    "codespace": "",
    "code": 0,
    "data": "12200A1E2F636F736D6F732E676F762E76312E4D7367566F7465526573706F6E7365",
    "raw_log": "[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgVote\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"}]},{\"type\":\"proposal_vote\",\"attributes\":[{\"key\":\"option\",\"value\":\"option:VOTE_OPTION_NO weight:\\\"1.000000000000000000\\\"\"},{\"key\":\"proposal_id\",\"value\":\"3\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/cosmos.gov.v1.MsgVote"
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
            "type": "proposal_vote",
            "attributes": [
              {
                "key": "option",
                "value": "option:VOTE_OPTION_NO weight:\"1.000000000000000000\""
              },
              {
                "key": "proposal_id",
                "value": "3"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "2000000",
    "gas_used": "130869",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.gov.v1.MsgVote",
            "proposal_id": "3",
            "voter": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
            "option": "VOTE_OPTION_NO",
            "metadata": ""
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
            "sequence": "2"
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
        "WP7OBGREtIph0H+htUPMdADVm1I+aNMRA7ueAI2QxkRObcyqmUJb1d0zaUBAv6U4/JOX7QeSKKR0DRFz2p8tngA="
      ]
    },
    "timestamp": "2023-01-10T01:28:09Z",
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
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43LzI=",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "V1A3T0JHUkV0SXBoMEgraHRVUE1kQURWbTFJK2FOTVJBN3VlQUkyUXhrUk9iY3lxbVVKYjFkMHphVUJBdjZVNC9KT1g3UWVTS0tSMERSRnoycDh0bmdBPQ==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2Nvc21vcy5nb3YudjEuTXNnVm90ZQ==",
            "index": false
          }
        ]
      },
      {
        "type": "proposal_vote",
        "attributes": [
          {
            "key": "b3B0aW9u",
            "value": "b3B0aW9uOlZPVEVfT1BUSU9OX05PIHdlaWdodDoiMS4wMDAwMDAwMDAwMDAwMDAwMDAi",
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
