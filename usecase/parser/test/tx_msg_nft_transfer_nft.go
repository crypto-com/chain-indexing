package usecase_parser_test

const TX_MSG_NFT_TRANSFER_NFT_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "A080FE35E2427EC312065A1C5B79AAA66CC404B9F58E582C192F96FCB7F842E7",
      "parts": {
        "total": 1,
        "hash": "7868503F822EEFB5F7A2923F698100E3F4BA64BBE28E97C3757F272836BF16DC"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "crypto-org-chain-mainnet-1",
        "height": "13103",
        "time": "2021-05-12T16:10:46.840382Z",
        "last_block_id": {
          "hash": "68BD5BF955D699B4F49EE399AA2A5EF918C03FC8DB4D3B91E015004C2231295E",
          "parts": {
            "total": 1,
            "hash": "2E3FE3872B10BBC41B9FBF57928BDA2156CC5FAAA34A47C6BAF928EFE821A6EE"
          }
        },
        "last_commit_hash": "DDBFD60E26D261B1127C49EDC8ABFD1FBE4F7A06452DFF5065F83DAC438F66E2",
        "data_hash": "A1342832C28738F34E761412B14B68AF712706875E311A70571A45924C5B9106",
        "validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "next_validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "AFA3CF297D770CCA8C8F0E03D2C017FCEA31844D5225186680D6EF798E6B5FBF",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "A996B7F7FF522DAB36127806F8570929AFE8404E"
      },
      "data": {
        "txs": [
          "CpIBCo8BCiAvY2hhaW5tYWluLm5mdC52MS5Nc2dUcmFuc2Zlck5GVBJrCgh0b2tlbmlkMhIHZGVub21pZBoqY3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2NkIipjcm8xMGFyN2RwdjRnOTVyM2dqMDZwbjUyZjY0dndxN3Y5a2RoMGc1OGESaQpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAiLBhn2Jb4CLU5dYpKB3LHDpjIFldrsQWD6LWDHyjpWMEgQKAggBGBQSFQoPCgdiYXNlY3JvEgQ1MDAwEMCaDBpANyfB8PAGQ942BHv7X5dqFGok+1eOBT+m5vAwDy+aA0R0p1oxpgR5dq+RA3kYUPv1tmUKhXihTMb+0oUtmv4b5g=="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "13102",
        "round": 0,
        "block_id": {
          "hash": "68BD5BF955D699B4F49EE399AA2A5EF918C03FC8DB4D3B91E015004C2231295E",
          "parts": {
            "total": 1,
            "hash": "2E3FE3872B10BBC41B9FBF57928BDA2156CC5FAAA34A47C6BAF928EFE821A6EE"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A996B7F7FF522DAB36127806F8570929AFE8404E",
            "timestamp": "2021-05-12T16:10:46.840382Z",
            "signature": "PttLn/dVRLehdwrNjvgDT91wdYuoFHopAH4TYBCYDHd+f2Y7NxTsPnMXD0AmtSe6fw3ajMjX9Y0amzLnGWP8BQ=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_NFT_TRANSFER_NFT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "13103",
    "txs_results": [
      {
        "code": 0,
        "data": "Cg4KDHRyYW5zZmVyX25mdA==",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"transfer_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"transfer_nft\",\"attributes\":[{\"key\":\"token_id\",\"value\":\"tokenid2\"},{\"key\":\"denom_id\",\"value\":\"denomid\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"recipient\",\"value\":\"cro10ar7dpv4g95r3gj06pn52f64vwq7v9kdh0g58a\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "61860",
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
                "value": "dHJhbnNmZXJfbmZ0",
                "index": true
              }
            ]
          },
          {
            "type": "transfer_nft",
            "attributes": [
              {
                "key": "dG9rZW5faWQ=",
                "value": "dG9rZW5pZDI=",
                "index": true
              },
              {
                "key": "ZGVub21faWQ=",
                "value": "ZGVub21pZA==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                "index": true
              },
              {
                "key": "cmVjaXBpZW50",
                "value": "Y3JvMTBhcjdkcHY0Zzk1cjNnajA2cG41MmY2NHZ3cTd2OWtkaDBnNThh",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "bmZ0",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
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
            "value": "NzEyNzk3Nzc4NzViYXNlY3Jv",
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
            "value": "MC40NDQ0MjU5OTIzMDU0OTk5OTQ=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTk5OTQwMDI5MTk2MTMzOTA=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "NDQ5ODgzNzQzNjU1MzMxNTI0LjczMTEzMTY0NzAxNjk5MjM5MA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NzEyNzk3Nzc4NzU=",
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
            "value": "NzEyNzk3Nzc4NzViYXNlY3Jv",
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
            "value": "MzU2Mzk4ODg5My43NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "MzU2Mzk4ODg5LjM3NTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
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
            "value": "MzU2Mzk4ODg5My43NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "Njc3MTU3ODg5OC4xMjUwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "Njc3MTU3ODg5ODEuMjUwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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
}
`

const TX_MSG_NFT_TRANSFER_NFT_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/chainmain.nft.v1.MsgTransferNFT",
                    "id": "tokenid2",
                    "denom_id": "denomid",
                    "sender": "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
                    "recipient": "cro10ar7dpv4g95r3gj06pn52f64vwq7v9kdh0g58a"
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
                    "sequence": "20"
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
            "NyfB8PAGQ942BHv7X5dqFGok+1eOBT+m5vAwDy+aA0R0p1oxpgR5dq+RA3kYUPv1tmUKhXihTMb+0oUtmv4b5g=="
        ]
    },
    "tx_response": {
        "height": "13103",
        "txhash": "8CF41CCC69DCE6B784B4A37B12017EE5A18A2018E17D2B6CEC3E06F4DFD7DFB1",
        "codespace": "",
        "code": 0,
        "data": "",
        "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"transfer_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"transfer_nft\",\"attributes\":[{\"key\":\"token_id\",\"value\":\"tokenid2\"},{\"key\":\"denom_id\",\"value\":\"denomid\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"},{\"key\":\"recipient\",\"value\":\"cro10ar7dpv4g95r3gj06pn52f64vwq7v9kdh0g58a\"}]}]}]",
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
                                "value": "dHJhbnNmZXJfbmZ0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer_nft",
                        "attributes": [
                            {
                                "key": "dG9rZW5faWQ=",
                                "value": "dG9rZW5pZDI=",
                                "index": true
                            },
                            {
                                "key": "ZGVub21faWQ=",
                                "value": "ZGVub21pZA==",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                                "index": true
                            },
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "Y3JvMTBhcjdkcHY0Zzk1cjNnajA2cG41MmY2NHZ3cTd2OWtkaDBnNThh",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "bmZ0",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
                                "index": true
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "61860",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/chainmain.nft.v1.MsgTransferNFT",
                        "id": "tokenid2",
                        "denom_id": "denomid",
                        "sender": "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
                        "recipient": "cro10ar7dpv4g95r3gj06pn52f64vwq7v9kdh0g58a"
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
                        "sequence": "20"
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
                "NyfB8PAGQ942BHv7X5dqFGok+1eOBT+m5vAwDy+aA0R0p1oxpgR5dq+RA3kYUPv1tmUKhXihTMb+0oUtmv4b5g=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
