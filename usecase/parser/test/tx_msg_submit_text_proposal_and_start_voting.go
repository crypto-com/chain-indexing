package usecase_parser_test

const TX_MSG_SUBMIT_TEXT_PROPOSAL_AND_START_VOTING_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "AD0F45F1348759E50119CD8E4EECE06FDB23902F0D6DA7D91E739BB829B1FA74",
      "parts": {
        "total": 1,
        "hash": "BB0C7948619C10ED78B89C05DAEA30B8093E74B35E7DBEED910D9C7FA030EA35"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "crypto-org-chain-mainnet-1",
        "height": "2036",
        "time": "2021-04-27T10:24:44.361437Z",
        "last_block_id": {
          "hash": "E053BF977E4670ECBE96D6EE9312B5079E6A5EDC112B0AE6E969C37B8E362471",
          "parts": {
            "total": 1,
            "hash": "9073B1487DEF3EF827BEE805404DD3522797C1B9E26124001379A546C3A4AA50"
          }
        },
        "last_commit_hash": "4B264FB2F3C208E1AB337D15E3E7488A033F46568410FFFA2FC38889C59C4D78",
        "data_hash": "C10D72A62E327CE1A211175E239D4778643D8A6F3F838B0A4463E17F63A03F07",
        "validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "next_validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "A8592B1C170AAF65C94F4C11054663E4E597420D3BF7A6F589300F6C67EF29EB",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "A996B7F7FF522DAB36127806F8570929AFE8404E"
      },
      "data": {
        "txs": [
          "CrUBCrIBCiUvY29zbW9zLmdvdi52MWJldGExLk1zZ1N1Ym1pdFByb3Bvc2FsEogBCkgKIC9jb3Ntb3MuZ292LnYxYmV0YTEuVGV4dFByb3Bvc2FsEiQKDVRlc3QgUHJvcG9zYWwSE015IGF3ZXNvbWUgcHJvcG9zYWwSEAoHYmFzZWNybxIFMjAwMDAaKmNybzFuazRycTNxNDZsdGdqZ2h4ejgwaHkzODVwOXVqMHRmNThhcGtjZBJpClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiECIsGGfYlvgItTl1ikoHcscOmMgWV2uxBYPotYMfKOlYwSBAoCCAEYAhIVCg8KB2Jhc2Vjcm8SBDUwMDAQwJoMGkDq37hntFruFNK05IIcv1my9Hjq0lr5x61e5+GlTBEJ31gWZXeZ5Cuhq5SRJ3FGn/VnDcoGb9a3Bezvevx+DPkA"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "2035",
        "round": 0,
        "block_id": {
          "hash": "E053BF977E4670ECBE96D6EE9312B5079E6A5EDC112B0AE6E969C37B8E362471",
          "parts": {
            "total": 1,
            "hash": "9073B1487DEF3EF827BEE805404DD3522797C1B9E26124001379A546C3A4AA50"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A996B7F7FF522DAB36127806F8570929AFE8404E",
            "timestamp": "2021-04-27T10:24:44.361437Z",
            "signature": "0rmFFM3+dYf77eYlqUPyP/yvdX1fzSkMvWFNgAun5+EDpw8r5PnX7qqRyQQIXdmcdpgx5OjrMlZvpYoSure7Cg=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_SUBMIT_TEXT_PROPOSAL_AND_START_VOTING_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "2036",
    "txs_results": [
      {
        "code": 0,
        "data": "ChUKD3N1Ym1pdF9wcm9wb3NhbBICCAI=",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"submit_proposal\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"20000basecro\"},{\"key\":\"proposal_id\",\"value\":\"2\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"2\"},{\"key\":\"proposal_type\",\"value\":\"Text\"},{\"key\":\"voting_period_start\",\"value\":\"2\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro10d07y265gmmuvt4z0w9aw880jnsr700jzemu2z\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"amount\",\"value\":\"20000basecro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "100600",
        "events": [
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMGJhc2Vjcm8=",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "c3VibWl0X3Byb3Bvc2Fs",
                "index": true
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "cHJvcG9zYWxfaWQ=",
                "value": "Mg==",
                "index": true
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "Y3JvMTBkMDd5MjY1Z21tdXZ0NHowdzlhdzg4MGpuc3I3MDBqemVtdTJ6",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MjAwMDBiYXNlY3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                "index": true
              }
            ]
          },
          {
            "type": "proposal_deposit",
            "attributes": [
              {
                "key": "YW1vdW50",
                "value": "MjAwMDBiYXNlY3Jv",
                "index": true
              },
              {
                "key": "cHJvcG9zYWxfaWQ=",
                "value": "Mg==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "Z292ZXJuYW5jZQ==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                "index": true
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "cHJvcG9zYWxfdHlwZQ==",
                "value": "VGV4dA==",
                "index": true
              },
              {
                "key": "dm90aW5nX3BlcmlvZF9zdGFydA==",
                "value": "Mg==",
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
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "Mzk2MDM5NzY4MzNiYXNlY3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
            "index": true
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC43OTk5OTQ4MzIwMTIwNDI1Njc=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTk5OTY3NzQxNjkzNjMxMDY=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MjQ5OTYxMjkxODYzMjIxNTQxLjY5ODE2Njc1MzI5NzI2MzgxNA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "Mzk2MDM5NzY4MzM=",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "Mzk2MDM5NzY4MzNiYXNlY3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
            "index": true
          }
        ]
      },
      {
        "type": "proposer_reward",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTk4MDE5ODg0MS42NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDEwYXI3ZHB2NGc5NXIzZ2owNnBuNTJmNjR2d3E3djlrZDV6dGE5cA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTk4MDE5ODg0LjE2NTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDEwYXI3ZHB2NGc5NXIzZ2owNnBuNTJmNjR2d3E3djlrZDV6dGE5cA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTk4MDE5ODg0MS42NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDEwYXI3ZHB2NGc5NXIzZ2owNnBuNTJmNjR2d3E3djlrZDV6dGE5cA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mzc2MjM3Nzc5OS4xMzUwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDEwYXI3ZHB2NGc5NXIzZ2owNnBuNTJmNjR2d3E3djlrZDV6dGE5cA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mzc2MjM3Nzc5OTEuMzUwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDEwYXI3ZHB2NGc5NXIzZ2owNnBuNTJmNjR2d3E3djlrZDV6dGE5cA==",
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
}`

const TX_MSG_SUBMIT_TEXT_PROPOSAL_AND_START_VOTING_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.gov.v1beta1.MsgSubmitProposal",
                    "content": {
                        "@type": "/cosmos.gov.v1beta1.TextProposal",
                        "title": "Test Proposal",
                        "description": "My awesome proposal"
                    },
                    "initial_deposit": [
                        {
                            "denom": "basecro",
                            "amount": "20000"
                        }
                    ],
                    "proposer": "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"
                }
            ],
            "memo": "",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
        },
        "auth_info": {
            "signer_infos": [
                {
                    "public_key": {
                        "@type": "/cosmos.crypto.secp256k1.PubKey",
                        "key": "AiLBhn2Jb4CLU5dYpKB3LHDpjIFldrsQWD6LWDHyjpWM"
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
                        "amount": "5000"
                    }
                ],
                "gas_limit": "200000",
                "payer": "",
                "granter": ""
            }
        },
        "signatures": [
            "6t+4Z7Ra7hTStOSCHL9ZsvR46tJa+cetXufhpUwRCd9YFmV3meQroauUkSdxRp/1Zw3KBm/WtwXs73r8fgz5AA=="
        ]
    },
    "tx_response": {
        "height": "2036",
        "txhash": "AAE71A09FFCDC3DD9D26CF2580CF26C38DDE2C6D7CBBF4264D295086B7E24148",
        "codespace": "",
        "code": 0,
        "data": "",
        "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"submit_proposal\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"20000basecro\"},{\"key\":\"proposal_id\",\"value\":\"2\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"2\"},{\"key\":\"proposal_type\",\"value\":\"Text\"},{\"key\":\"voting_period_start\",\"value\":\"2\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro10d07y265gmmuvt4z0w9aw880jnsr700jzemu2z\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"amount\",\"value\":\"20000basecro\"}]}]}]",
        "logs": [
            {
                "msg_index": 0,
                "log": "",
                "events": [
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "NTAwMGJhc2Vjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "c3VibWl0X3Byb3Bvc2Fs",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "submit_proposal",
                        "attributes": [
                            {
                                "key": "cHJvcG9zYWxfaWQ=",
                                "value": "Mg==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "Y3JvMTBkMDd5MjY1Z21tdXZ0NHowdzlhdzg4MGpuc3I3MDBqemVtdTJ6",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MjAwMDBiYXNlY3Jv",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "proposal_deposit",
                        "attributes": [
                            {
                                "key": "YW1vdW50",
                                "value": "MjAwMDBiYXNlY3Jv",
                                "index": true
                            },
                            {
                                "key": "cHJvcG9zYWxfaWQ=",
                                "value": "Mg==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "Z292ZXJuYW5jZQ==",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "submit_proposal",
                        "attributes": [
                            {
                                "key": "cHJvcG9zYWxfdHlwZQ==",
                                "value": "VGV4dA==",
                                "index": true
                            },
                            {
                                "key": "dm90aW5nX3BlcmlvZF9zdGFydA==",
                                "value": "Mg==",
                                "index": true
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "100600",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.gov.v1beta1.MsgSubmitProposal",
                        "content": {
                            "@type": "/cosmos.gov.v1beta1.TextProposal",
                            "title": "Test Proposal",
                            "description": "My awesome proposal"
                        },
                        "initial_deposit": [
                            {
                                "denom": "basecro",
                                "amount": "20000"
                            }
                        ],
                        "proposer": "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"
                    }
                ],
                "memo": "",
                "timeout_height": "0",
                "extension_options": [],
                "non_critical_extension_options": []
            },
            "auth_info": {
                "signer_infos": [
                    {
                        "public_key": {
                            "@type": "/cosmos.crypto.secp256k1.PubKey",
                            "key": "AiLBhn2Jb4CLU5dYpKB3LHDpjIFldrsQWD6LWDHyjpWM"
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
                            "amount": "5000"
                        }
                    ],
                    "gas_limit": "200000",
                    "payer": "",
                    "granter": ""
                }
            },
            "signatures": [
                "6t+4Z7Ra7hTStOSCHL9ZsvR46tJa+cetXufhpUwRCd9YFmV3meQroauUkSdxRp/1Zw3KBm/WtwXs73r8fgz5AA=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
