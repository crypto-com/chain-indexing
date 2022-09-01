package usecase_parser_test

const TX_MSG_NFT_BURN_NFT_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "3BCBDB18DF19C266C872CDFF4D34A52F6199BEC917D2AF7DDEBD5B6160B5A235",
      "parts": {
        "total": 1,
        "hash": "D50E16E89B31DB460AFEDC18B6A5E3FE9DABD81BCD637D990902744E3452B2B5"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "crypto-org-chain-mainnet-1",
        "height": "17699",
        "time": "2021-05-13T06:50:01.780259Z",
        "last_block_id": {
          "hash": "A1AF84D5092A51A3A943F2464F038171E10BF5619AAC4A98DF4B6D9A3CF4BCAF",
          "parts": {
            "total": 1,
            "hash": "173FA9B81CFE891451A5E91D0547DFCA8AA0A474F17C1D4D4CF48AB5EC6C0B58"
          }
        },
        "last_commit_hash": "6ECA67C24E06880433B44465224010AD9B5EDEE14BF11CFBF0424A762E4BF7A6",
        "data_hash": "7FB8806A045CD900FC9065169359F6EAA2E6A9F83F19AB87D54735A10A9BF3B1",
        "validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "next_validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "7A3406C77D0CF9772C5AEA8140EE24E1644B23956C63D0D2B3CD286616A2A429",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "A996B7F7FF522DAB36127806F8570929AFE8404E"
      },
      "data": {
        "txs": [
          "CmEKXwocL2NoYWlubWFpbi5uZnQudjEuTXNnQnVybk5GVBI/Cgh0b2tlbmlkNBIHZGVub21pZBoqY3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2NkEmkKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQIiwYZ9iW+Ai1OXWKSgdyxw6YyBZXa7EFg+i1gx8o6VjBIECgIIARgWEhUKDwoHYmFzZWNybxIENTAwMBDAmgwaQIccLPZgG69mz6KGkmh3riPciB3Hk4CnGLHVWWgFLks5SA6PmAUOLEEx+miFVss01/ZEaLuqaqs3w44CQf60O5Q="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "17698",
        "round": 0,
        "block_id": {
          "hash": "A1AF84D5092A51A3A943F2464F038171E10BF5619AAC4A98DF4B6D9A3CF4BCAF",
          "parts": {
            "total": 1,
            "hash": "173FA9B81CFE891451A5E91D0547DFCA8AA0A474F17C1D4D4CF48AB5EC6C0B58"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A996B7F7FF522DAB36127806F8570929AFE8404E",
            "timestamp": "2021-05-13T06:50:01.780259Z",
            "signature": "NYXDbmjn/39UfOgdFjGRvi8ZAoPBOoz4A3jjeIE/KXYbOBHJmklQ3YDe/zRFBXYPFujq5uZb58a63wS4guoiBg=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_NFT_BURN_NFT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "17699",
    "txs_results": [
      {
        "code": 0,
        "data": "CgoKCGJ1cm5fbmZ0",
        "log": "[{\"events\":[{\"type\":\"burn_nft\",\"attributes\":[{\"key\":\"denom_id\",\"value\":\"denomid\"},{\"key\":\"token_id\",\"value\":\"tokenid4\"},{\"key\":\"owner\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"burn_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "59471",
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
                "value": "YnVybl9uZnQ=",
                "index": true
              }
            ]
          },
          {
            "type": "burn_nft",
            "attributes": [
              {
                "key": "ZGVub21faWQ=",
                "value": "ZGVub21pZA==",
                "index": true
              },
              {
                "key": "dG9rZW5faWQ=",
                "value": "dG9rZW5pZDQ=",
                "index": true
              },
              {
                "key": "b3duZXI=",
                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
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
            "value": "NzEyNzMzMTY5NDhiYXNlY3Jv",
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
            "value": "MC40NDQ0MTk1MjIwODIzNzI4OTI=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTk5OTE4OTk1Njg5NTEzMjc=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "NDQ5ODQyOTY1Mzg3MjM2OTgwLjU4NzY3MzA1OTgwNzA1NTIwNw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NzEyNzMzMTY5NDg=",
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
            "value": "NzEyNzMzMTY5NDhiYXNlY3Jv",
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
            "value": "MzU2MzY2NTg0Ny40MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "MzU2MzY2NTg0Ljc0MDAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
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
            "value": "MzU2MzY2NTg0Ny40MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "Njc3MDk2NTExMC4wNjAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "Njc3MDk2NTExMDAuNjAwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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

const TX_MSG_NFT_BURN_NFT_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/chainmain.nft.v1.MsgBurnNFT",
                    "id": "tokenid4",
                    "denom_id": "denomid",
                    "sender": "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"
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
                    "sequence": "22"
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
            "hxws9mAbr2bPooaSaHeuI9yIHceTgKcYsdVZaAUuSzlIDo+YBQ4sQTH6aIVWyzTX9kRou6pqqzfDjgJB/rQ7lA=="
        ]
    },
    "tx_response": {
        "height": "17699",
        "txhash": "63B42F5AC39D758E5590E7D54A0F811968D1C5C0420EA5162CE83CA6D6818AA5",
        "codespace": "",
        "code": 0,
        "data": "CgoKCGJ1cm5fbmZ0",
        "raw_log": "[{\"events\":[{\"type\":\"burn_nft\",\"attributes\":[{\"key\":\"denom_id\",\"value\":\"denomid\"},{\"key\":\"token_id\",\"value\":\"tokenid4\"},{\"key\":\"owner\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"burn_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]}]}]",
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
                                "value": "YnVybl9uZnQ=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "burn_nft",
                        "attributes": [
                            {
                                "key": "ZGVub21faWQ=",
                                "value": "ZGVub21pZA==",
                                "index": true
                            },
                            {
                                "key": "dG9rZW5faWQ=",
                                "value": "dG9rZW5pZDQ=",
                                "index": true
                            },
                            {
                                "key": "b3duZXI=",
                                "value": "Y3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2Nk",
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
        "gas_used": "59471",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/chainmain.nft.v1.MsgBurnNFT",
                        "id": "tokenid4",
                        "denom_id": "denomid",
                        "sender": "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"
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
                        "sequence": "22"
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
                "hxws9mAbr2bPooaSaHeuI9yIHceTgKcYsdVZaAUuSzlIDo+YBQ4sQTH6aIVWyzTX9kRou6pqqzfDjgJB/rQ7lA=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
