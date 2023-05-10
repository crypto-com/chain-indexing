package usecase_parser_test

const TX_MSG_VOTE_WEIGHTED_V1_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "2E2623A467973C37F54C048A140C43CDFF8795B52671CD0F34CB3682BD91DE2B",
      "parts": {
        "total": 1,
        "hash": "0DE0AC94D90253D5EB7237821816F3E706A48330A6317246090974D9556D1497"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "cronos_777-1",
        "height": "5284",
        "time": "2023-01-10T01:30:33.948617Z",
        "last_block_id": {
          "hash": "FC5E7CB1B721478AE14BC70B70C115CA302C5DA38D10E1E26DAF1946DEDABD15",
          "parts": {
            "total": 1,
            "hash": "FA6CD47E138597264404E411BBEC1E8ABE2211F3BDEF54C29DBF25F02D70FC19"
          }
        },
        "last_commit_hash": "EE996EA6AEB531F8F0419192845854AF0C76011A6CC62D527EBCBE455DCC0C8D",
        "data_hash": "6C1C71B1D04450ED10596722BC913ED462E031117CA1ECBDCF2AB9BCC51313D3",
        "validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "next_validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
        "app_hash": "2C62B6709CD2446C4C1BDEC6D44490D6696CE34F17C9B23A65DB9096EE77C3D0",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "FFAF83126C47B961299A37B647A112070AC4228A"
      },
      "data": {
        "txs": [
          "CocBCoQBCh4vY29zbW9zLmdvdi52MS5Nc2dWb3RlV2VpZ2h0ZWQSYggBEipjcmMxOHo2cTM4bWh2dHN2eXI1bWFrOGZqOHM4ZzRndzdramp0c2dybjcaGAgBEhQwLjIwMDAwMDAwMDAwMDAwMDAwMBoYCAMSFDAuODAwMDAwMDAwMDAwMDAwMDAwEoABClkKTwooL2V0aGVybWludC5jcnlwdG8udjEuZXRoc2VjcDI1NmsxLlB1YktleRIjCiECQnhadQdEUtYqasIicP+4+wHJN10LpyiHroANxhkxXRsSBAoCCAEYAxIjCh0KB2Jhc2Vjcm8SEjEyOTc1OTA2MzcyNDAwMDAwMBCAiXoaQZ9P6wIIgjDoPUs6ASV3Sgabj5Yo4h7NUN3Wxt3DKeqhM7oO2z9WB1jQ4AtPy/VkpCi0Rd5InlzCgwUaKP3ZB+cA"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "5283",
        "round": 0,
        "block_id": {
          "hash": "FC5E7CB1B721478AE14BC70B70C115CA302C5DA38D10E1E26DAF1946DEDABD15",
          "parts": {
            "total": 1,
            "hash": "FA6CD47E138597264404E411BBEC1E8ABE2211F3BDEF54C29DBF25F02D70FC19"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57",
            "timestamp": "2023-01-10T01:30:34.0142Z",
            "signature": "UBuyxcvV3hunucHBNutzTY8WXQk45Z1ZFP02lCTYzuIly5byAo+9ni/bdBwvkdaAQPoPBno70kKR+xBgNUJsBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "FFAF83126C47B961299A37B647A112070AC4228A",
            "timestamp": "2023-01-10T01:30:33.948617Z",
            "signature": "GnG/plujdqWpGfZo9hLhus2A2B0+3EVdh7QRG9Tdl6oaEdIQRF/AX5uOBb2cKBBEBo4NRcy8XDAR3YyCG9k0DQ=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_VOTE_WEIGHTED_V1_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "5284",
    "txs_results": [
      {
        "code": 0,
        "data": "EigKJi9jb3Ntb3MuZ292LnYxLk1zZ1ZvdGVXZWlnaHRlZFJlc3BvbnNl",
        "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgVoteWeighted\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"}]},{\"type\":\"proposal_vote\",\"attributes\":[{\"key\":\"option\",\"value\":\"option:VOTE_OPTION_YES weight:\\\"0.200000000000000000\\\" \\noption:VOTE_OPTION_NO weight:\\\"0.800000000000000000\\\"\"},{\"key\":\"proposal_id\",\"value\":\"1\"}]}]}]",
        "info": "",
        "gas_wanted": "2000000",
        "gas_used": "132474",
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
                "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43LzM=",
                "index": false
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "bjAvckFnaUNNT2c5U3pvQkpYZEtCcHVQbGlqaUhzMVEzZGJHM2NNcDZxRXp1ZzdiUDFZSFdORGdDMC9MOVdTa0tMUkYza2llWE1LREJSb28vZGtINXdBPQ==",
                "index": false
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2Nvc21vcy5nb3YudjEuTXNnVm90ZVdlaWdodGVk",
                "index": false
              }
            ]
          },
          {
            "type": "proposal_vote",
            "attributes": [
              {
                "key": "b3B0aW9u",
                "value": "b3B0aW9uOlZPVEVfT1BUSU9OX1lFUyB3ZWlnaHQ6IjAuMjAwMDAwMDAwMDAwMDAwMDAwIiAKb3B0aW9uOlZPVEVfT1BUSU9OX05PIHdlaWdodDoiMC44MDAwMDAwMDAwMDAwMDAwMDAi",
                "index": false
              },
              {
                "key": "cHJvcG9zYWxfaWQ=",
                "value": "MQ==",
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
            "value": "NDExODIwMTAzODhzdGFrZQ==",
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
            "value": "NDExODIwMTAzODhzdGFrZQ==",
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
            "value": "NDExODIwMTAzODhzdGFrZQ==",
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
            "value": "NDExODIwMTAzODhzdGFrZQ==",
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
            "value": "NDExODIwMTAzODhzdGFrZQ==",
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
            "value": "MC45OTk4OTEyMTMwNDIzMzY1MTI=",
            "index": false
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMjk5NDY0MDMwOTE1ODc1MTA=",
            "index": false
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MjU5OTIxMDgyMjA2OTQwOTE0LjMwNTY1ODkwMjc4Mzc3NzIxMA==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "NDExODIwMTAzODg=",
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
            "value": "NDExODIwMTAzODhzdGFrZQ==",
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
            "value": "NDExODIwMTAzODhzdGFrZQ==",
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
            "value": "NDExODIwMTAzODhzdGFrZQ==",
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
            "value": "MjA1OTEwMDUxOS40MDAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MjA1OTEwMDUxLjk0MDAwMDAwMDAwMDAwMDAwMHN0YWtl",
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
            "value": "MjA1OTEwMDUxOS40MDAwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk2MzQ4My4wNDIwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk2MzQ4MzAuNDIwMDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "MTkxNDk2MzQ4My4wNDIwMDAwMDAwMDAwMDAwMDBzdGFrZQ==",
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
            "value": "MTkxNDk2MzQ4MzAuNDIwMDAwMDAwMDAwMDAwMDAwc3Rha2U=",
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
            "value": "NTI4NA==",
            "index": false
          },
          {
            "key": "YW1vdW50",
            "value": "MTMyNDc0",
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

const TX_MSG_VOTE_WEIGHTED_V1_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/cosmos.gov.v1.MsgVoteWeighted",
          "proposal_id": "1",
          "voter": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
          "options": [
            {
              "option": "VOTE_OPTION_YES",
              "weight": "0.200000000000000000"
            },
            {
              "option": "VOTE_OPTION_NO",
              "weight": "0.800000000000000000"
            }
          ],
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
      "n0/rAgiCMOg9SzoBJXdKBpuPlijiHs1Q3dbG3cMp6qEzug7bP1YHWNDgC0/L9WSkKLRF3kieXMKDBRoo/dkH5wA="
    ]
  },
  "tx_response": {
    "height": "5284",
    "txhash": "F258B015674DAFBE5D1788825690BC31D1F6254F599944AD4CDEE9A461179FF7",
    "codespace": "",
    "code": 0,
    "data": "12280A262F636F736D6F732E676F762E76312E4D7367566F74655765696768746564526573706F6E7365",
    "raw_log": "[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.gov.v1.MsgVoteWeighted\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7\"}]},{\"type\":\"proposal_vote\",\"attributes\":[{\"key\":\"option\",\"value\":\"option:VOTE_OPTION_YES weight:\\\"0.200000000000000000\\\" \\noption:VOTE_OPTION_NO weight:\\\"0.800000000000000000\\\"\"},{\"key\":\"proposal_id\",\"value\":\"1\"}]}]}]",
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
                "value": "/cosmos.gov.v1.MsgVoteWeighted"
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
                "value": "option:VOTE_OPTION_YES weight:\"0.200000000000000000\" \noption:VOTE_OPTION_NO weight:\"0.800000000000000000\""
              },
              {
                "key": "proposal_id",
                "value": "1"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "2000000",
    "gas_used": "132474",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.gov.v1.MsgVoteWeighted",
            "proposal_id": "1",
            "voter": "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
            "options": [
              {
                "option": "VOTE_OPTION_YES",
                "weight": "0.200000000000000000"
              },
              {
                "option": "VOTE_OPTION_NO",
                "weight": "0.800000000000000000"
              }
            ],
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
        "n0/rAgiCMOg9SzoBJXdKBpuPlijiHs1Q3dbG3cMp6qEzug7bP1YHWNDgC0/L9WSkKLRF3kieXMKDBRoo/dkH5wA="
      ]
    },
    "timestamp": "2023-01-10T01:30:33Z",
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
            "value": "Y3JjMTh6NnEzOG1odnRzdnlyNW1hazhmajhzOGc0Z3c3a2pqdHNncm43LzM=",
            "index": false
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "bjAvckFnaUNNT2c5U3pvQkpYZEtCcHVQbGlqaUhzMVEzZGJHM2NNcDZxRXp1ZzdiUDFZSFdORGdDMC9MOVdTa0tMUkYza2llWE1LREJSb28vZGtINXdBPQ==",
            "index": false
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2Nvc21vcy5nb3YudjEuTXNnVm90ZVdlaWdodGVk",
            "index": false
          }
        ]
      },
      {
        "type": "proposal_vote",
        "attributes": [
          {
            "key": "b3B0aW9u",
            "value": "b3B0aW9uOlZPVEVfT1BUSU9OX1lFUyB3ZWlnaHQ6IjAuMjAwMDAwMDAwMDAwMDAwMDAwIiAKb3B0aW9uOlZPVEVfT1BUSU9OX05PIHdlaWdodDoiMC44MDAwMDAwMDAwMDAwMDAwMDAi",
            "index": false
          },
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MQ==",
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
