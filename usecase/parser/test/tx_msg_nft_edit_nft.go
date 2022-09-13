package usecase_parser_test

const TX_MSG_NFT_EDIT_NFT_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "5095AC342A6862BA01615E97F6037CA21FFF6D75853DA98269978F5E44F319C1",
      "parts": {
        "total": 1,
        "hash": "DDEF5D34F1888535CAF13B42D9DC0755DFF19E1976E76BC3D931749664ED4C04"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "crypto-org-chain-mainnet-1",
        "height": "16754",
        "time": "2021-05-13T05:18:04.485445Z",
        "last_block_id": {
          "hash": "D78A678683D8F1F42FD637EC84172E2DC9E049C9F07E425638EFD58BCC30B678",
          "parts": {
            "total": 1,
            "hash": "D2D7DAD44054742D1F658B65EEB0C132CAAEE874E4C805717C0067AB2184D494"
          }
        },
        "last_commit_hash": "8B0B62AFFB933488C0E5FC944DD8155E106FC527C7C4855CF45C6DB99C1C5CF5",
        "data_hash": "9067E458031E78F4F0252E0F68E828DD3F5D287E0FD12ED055696CB871045677",
        "validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "next_validators_hash": "243F22D7662AF9DDF28B5D56CB4DF810E6B2F7924632C8FCA13C2494FFB93949",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "DFD46CD395B0E8CDF154392428618F84F06DBF06917199F214B00AEFDA4D4679",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "A996B7F7FF522DAB36127806F8570929AFE8404E"
      },
      "data": {
        "txs": [
          "CnsKeQocL2NoYWlubWFpbi5uZnQudjEuTXNnRWRpdE5GVBJZCgh0b2tlbmlkOBIHZGVub21pZBoHbmV3TmFtZSIGbmV3VXJpKgduZXdEYXRhMipjcm8xbms0cnEzcTQ2bHRnamdoeHo4MGh5Mzg1cDl1ajB0ZjU4YXBrY2QSaQpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAiLBhn2Jb4CLU5dYpKB3LHDpjIFldrsQWD6LWDHyjpWMEgQKAggBGBUSFQoPCgdiYXNlY3JvEgQ1MDAwEMCaDBpAvW3ZCLb++VSnkojQIRyAds0+tk5glGRpapMfOyXkoGg/9+hK2OoOf1q5mcPsA0l8dodG1UHWpmLOf0G6YdcPyQ=="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "16753",
        "round": 0,
        "block_id": {
          "hash": "D78A678683D8F1F42FD637EC84172E2DC9E049C9F07E425638EFD58BCC30B678",
          "parts": {
            "total": 1,
            "hash": "D2D7DAD44054742D1F658B65EEB0C132CAAEE874E4C805717C0067AB2184D494"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A996B7F7FF522DAB36127806F8570929AFE8404E",
            "timestamp": "2021-05-13T05:18:04.485445Z",
            "signature": "huxVPUfZZYy9z4x0XvCTI3iIG5uDzu42TH0l0XV0kH3jGNn2fmu4jqt7CuiZSeNGngjSostJhIwbd918KFV6AQ=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_NFT_EDIT_NFT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "16754",
    "txs_results": [
      {
        "code": 0,
        "data": "CgoKCGVkaXRfbmZ0",
        "log": "[{\"events\":[{\"type\":\"edit_nft\",\"attributes\":[{\"key\":\"token_id\",\"value\":\"tokenid8\"},{\"key\":\"denom_id\",\"value\":\"denomid\"},{\"key\":\"token_uri\",\"value\":\"newUri\"},{\"key\":\"owner\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"edit_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "60439",
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
                "value": "ZWRpdF9uZnQ=",
                "index": true
              }
            ]
          },
          {
            "type": "edit_nft",
            "attributes": [
              {
                "key": "dG9rZW5faWQ=",
                "value": "dG9rZW5pZDg=",
                "index": true
              },
              {
                "key": "ZGVub21faWQ=",
                "value": "ZGVub21pZA==",
                "index": true
              },
              {
                "key": "dG9rZW5fdXJp",
                "value": "bmV3VXJp",
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
            "value": "NzEyNzQ2NDU0MDBiYXNlY3Jv",
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
            "value": "MC40NDQ0MjA4NTIzODQ4MTAyODQ=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTk5OTIzMzIwMzg3MDkwNjY=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "NDQ5ODUxMzQ5OTQwNjcxMTc3LjAzMTc1MTM5ODgwMzY5NjM3Mg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NzEyNzQ2NDU0MDA=",
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
            "value": "NzEyNzQ2NDU0MDBiYXNlY3Jv",
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
            "value": "MzU2MzczMjI3MC4wMDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "MzU2MzczMjI3LjAwMDAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
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
            "value": "MzU2MzczMjI3MC4wMDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "Njc3MTA5MTMxMy4wMDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
            "value": "Njc3MTA5MTMxMzAuMDAwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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

const TX_MSG_NFT_EDIT_NFT_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/chainmain.nft.v1.MsgEditNFT",
                    "id": "tokenid8",
                    "denom_id": "denomid",
                    "name": "newName",
                    "uri": "newUri",
                    "data": "newData",
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
                    "sequence": "21"
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
            "vW3ZCLb++VSnkojQIRyAds0+tk5glGRpapMfOyXkoGg/9+hK2OoOf1q5mcPsA0l8dodG1UHWpmLOf0G6YdcPyQ=="
        ]
    },
    "tx_response": {
        "height": "16754",
        "txhash": "652358F0C39307A88C347013DF65317498DB81FE847821066559E401AFF8F632",
        "codespace": "",
        "code": 0,
        "data": "CgoKCGVkaXRfbmZ0",
        "raw_log": "[{\"events\":[{\"type\":\"edit_nft\",\"attributes\":[{\"key\":\"token_id\",\"value\":\"tokenid8\"},{\"key\":\"denom_id\",\"value\":\"denomid\"},{\"key\":\"token_uri\",\"value\":\"newUri\"},{\"key\":\"owner\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"edit_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd\"}]}]}]",
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
                                "value": "ZWRpdF9uZnQ=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "edit_nft",
                        "attributes": [
                            {
                                "key": "dG9rZW5faWQ=",
                                "value": "dG9rZW5pZDg=",
                                "index": true
                            },
                            {
                                "key": "ZGVub21faWQ=",
                                "value": "ZGVub21pZA==",
                                "index": true
                            },
                            {
                                "key": "dG9rZW5fdXJp",
                                "value": "bmV3VXJp",
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
        "gas_used": "60439",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/chainmain.nft.v1.MsgEditNFT",
                        "id": "tokenid8",
                        "denom_id": "denomid",
                        "name": "newName",
                        "uri": "newUri",
                        "data": "newData",
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
                        "sequence": "21"
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
                "vW3ZCLb++VSnkojQIRyAds0+tk5glGRpapMfOyXkoGg/9+hK2OoOf1q5mcPsA0l8dodG1UHWpmLOf0G6YdcPyQ=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
