package usecase_parser_test

const TX_MSG_DEPOSIT_AND_START_VOTING_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "A65977CD5DE91116F6B0A2C04D8AB30D31C8997DABF2A7F61DDB287B4B3703DB",
      "parts": {
        "total": 1,
        "hash": "A929D5ED1C523D03B3D0782D46D4BBBF8A2EBD0FCF2D485E159D0FFD8A8D7074"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "crypto-org-chain-mainnet-1",
        "height": "439",
        "time": "2021-04-27T08:05:55.515022Z",
        "last_block_id": {
          "hash": "40AC2F47BB73D4BE076DD0E5C151EADD332A1D867A7C1F38F58B3C23FD62861E",
          "parts": {
            "total": 1,
            "hash": "7A684555BECBDCCC5ED732E12053ACED72A1537FA253CDB2142FBD37F777BC5C"
          }
        },
        "last_commit_hash": "E2ECF3721C82075047613E9A08C61F6DC0C445B41ECF25D155569FC6BDFAD233",
        "data_hash": "EBCE4FB09C16C0F776F4D428AA59B335EAF1E4925F3D45D6FE5365D650D33CB5",
        "validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "next_validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "565CFCF33BE823055CEE34FD272DD821786BE3986B113C15E1F3368560EDB761",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "A996B7F7FF522DAB36127806F8570929AFE8404E"
      },
      "data": {
        "txs": [
          "CmQKYgoeL2Nvc21vcy5nb3YudjFiZXRhMS5Nc2dEZXBvc2l0EkAIARIqY3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2NkGhAKB2Jhc2Vjcm8SBTEwMDAwEmkKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQIiwYZ9iW+Ai1OXWKSgdyxw6YyBZXa7EFg+i1gx8o6VjBIECgIIARgBEhUKDwoHYmFzZWNybxIENTAwMBDAmgwaQH8Zxq6K1REakNhcyF+MtBsFRtJX91LQz+puWgNNPWimdymj/TUc4UnFZIF3unUwL5yBsy0bgW8eszlDx5WOxYg="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "438",
        "round": 0,
        "block_id": {
          "hash": "40AC2F47BB73D4BE076DD0E5C151EADD332A1D867A7C1F38F58B3C23FD62861E",
          "parts": {
            "total": 1,
            "hash": "7A684555BECBDCCC5ED732E12053ACED72A1537FA253CDB2142FBD37F777BC5C"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A996B7F7FF522DAB36127806F8570929AFE8404E",
            "timestamp": "2021-04-27T08:05:55.515022Z",
            "signature": "VC1Gx+w7KQl0EAojqGGIQk7E5lwqXimrCuLbL1ax5tb+c71nBa+fehQp9rOMI6c6KJ8RHPUAx66d426wGnw1BA=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_DEPOSIT_AND_START_VOTING_BLOCK_RESULT_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "439",
    "txs_results": [
      {
        "code": 0,
        "data": "CgkKB2RlcG9zaXQ=",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"deposit\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"10000basecro\"},{\"key\":\"proposal_id\",\"value\":\"1\"},{\"key\":\"voting_period_start\",\"value\":\"1\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro10d07y265gmmuvt4z0w9aw880jnsr700jzemu2z\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"amount\",\"value\":\"10000basecro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "86860",
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
                "value": "ZGVwb3NpdA==",
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
                "value": "MTAwMDBiYXNlY3Jv",
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
                "value": "MTAwMDBiYXNlY3Jv",
                "index": true
              },
              {
                "key": "cHJvcG9zYWxfaWQ=",
                "value": "MQ==",
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
            "type": "proposal_deposit",
            "attributes": [
              {
                "key": "dm90aW5nX3BlcmlvZF9zdGFydA==",
                "value": "MQ==",
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
            "value": "Mzk2MDg3ODc2NzBiYXNlY3Jv",
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
            "value": "MC43OTk5OTg4ODAwNjk0OTY0MjI=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTk5OTkzMDQ0NDcyNTgxMzc=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MjQ5OTkxNjU1NTU3MzI3NTg1LjUyNjM3NDU0Nzc5NjQxMjE1Ng==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "Mzk2MDg3ODc2NzA=",
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
            "value": "Mzk2MDg3ODc2NzBiYXNlY3Jv",
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
            "value": "MTk4MDQzOTM4My41MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "MTk4MDQzOTM4LjM1MDAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
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
            "value": "MTk4MDQzOTM4My41MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "Mzc2MjgzNDgyOC42NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "Mzc2MjgzNDgyODYuNTAwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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

const TX_MSG_DEPOSIT_AND_START_VOTING_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.gov.v1beta1.MsgDeposit",
                    "proposal_id": "1",
                    "depositor": "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
                    "amount": [
                        {
                            "denom": "basecro",
                            "amount": "10000"
                        }
                    ]
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
                    "sequence": "1"
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
            "fxnGrorVERqQ2FzIX4y0GwVG0lf3UtDP6m5aA009aKZ3KaP9NRzhScVkgXe6dTAvnIGzLRuBbx6zOUPHlY7FiA=="
        ]
    },
    "tx_response": {
        "height": "439",
        "txhash": "3EB28276333878ABCBB0D0ACB942A6F94BC23BFFE3E972B9050509D342C7F747",
        "codespace": "",
        "code": 0,
        "data": "CgkKB2RlcG9zaXQ=",
        "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"deposit\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"proposal_deposit\",\"attributes\":[{\"key\":\"amount\",\"value\":\"10000basecro\"},{\"key\":\"proposal_id\",\"value\":\"1\"},{\"key\":\"voting_period_start\",\"value\":\"1\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro10d07y265gmmuvt4z0w9aw880jnsr700jzemu2z\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"amount\",\"value\":\"10000basecro\"}]}]}]",
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
                                "value": "ZGVwb3NpdA==",
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
                                "value": "MTAwMDBiYXNlY3Jv",
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
                                "value": "MTAwMDBiYXNlY3Jv",
                                "index": true
                            },
                            {
                                "key": "cHJvcG9zYWxfaWQ=",
                                "value": "MQ==",
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
                        "type": "proposal_deposit",
                        "attributes": [
                            {
                                "key": "dm90aW5nX3BlcmlvZF9zdGFydA==",
                                "value": "MQ==",
                                "index": true
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "86860",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.gov.v1beta1.MsgDeposit",
                        "proposal_id": "1",
                        "depositor": "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
                        "amount": [
                            {
                                "denom": "basecro",
                                "amount": "10000"
                            }
                        ]
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
                        "sequence": "1"
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
                "fxnGrorVERqQ2FzIX4y0GwVG0lf3UtDP6m5aA009aKZ3KaP9NRzhScVkgXe6dTAvnIGzLRuBbx6zOUPHlY7FiA=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
