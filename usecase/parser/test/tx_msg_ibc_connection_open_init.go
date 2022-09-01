package usecase_parser_test

const TX_MSG_CONNECTION_OPEN_INIT_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "7691AB2EF59EC7B98495E213B69AE14D7C856A82C0C000AF41668F59EDD4E843",
      "parts": {
        "total": 1,
        "hash": "E18960251BAADEEC487BC562B0B22242F53F9BCFAB5DF9D82049B272AB1AF0EA"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "devnet-1",
        "height": "6",
        "time": "2021-05-04T18:02:48.272819Z",
        "last_block_id": {
          "hash": "121B48440580D4384DFC6CF9D17444BEB5C7D7E293902F7B3D7513DA8DC44E39",
          "parts": {
            "total": 1,
            "hash": "CBD380F934F8A87A8DCB8BB6C3F1BFD824FC168DE127EDF951F8C135A91E2560"
          }
        },
        "last_commit_hash": "45A3F2F169DFF321E916A990236F197A1D7AAD9455FD2AED10BF23E4540F4323",
        "data_hash": "51FFFD76B5509B30666167204A318E8ECCA92AC9B3150C1B79C2E523AF29324E",
        "validators_hash": "3F4AEF8603CD6582C072B7EACBBF1BDD9591996AF659617CC72683B969BE8DF5",
        "next_validators_hash": "3F4AEF8603CD6582C072B7EACBBF1BDD9591996AF659617CC72683B969BE8DF5",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "CEF5F03340FE2484D7524EF74578EEFB236F69301EAAF2D0A6C56440D6F5BEB6",
        "last_results_hash": "FF4071626D6D4053F15986DACD2B992A314D3770A495C724040B74195B857281",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "CD1AC8BFC2C01FC3C2DAF26628D1CC820FD0CF6C"
      },
      "data": {
        "txs": [
          "CrABCq0BCi0vaWJjLmNvcmUuY29ubmVjdGlvbi52MS5Nc2dDb25uZWN0aW9uT3BlbkluaXQSfAoPMDctdGVuZGVybWludC0wEhgKDzA3LXRlbmRlcm1pbnQtMBoFCgNpYmMaIwoBMRINT1JERVJfT1JERVJFRBIPT1JERVJfVU5PUkRFUkVEKipjcm8xZ2Rzd3Jtd3R6Z3Yza3ZmMjhsdnR0N3F2N3E3bXl6bW40NjZyM2YSagpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAy3F5FOQif8MrhCypdaMnizCw4zKh3WZMB6OYdD04+eIEgQKAggBGAESFgoPCgdiYXNlY3JvEgQxMDAwEMCNtwEaQGFlO44JqCt7T4sh0EQLIC8WUKMCqrp+D7KKLjY1RH3TVdD4DhIx2n1Zn9rHNrN2XIv13l7JkSQs1bfck8GUCN0="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "5",
        "round": 0,
        "block_id": {
          "hash": "121B48440580D4384DFC6CF9D17444BEB5C7D7E293902F7B3D7513DA8DC44E39",
          "parts": {
            "total": 1,
            "hash": "CBD380F934F8A87A8DCB8BB6C3F1BFD824FC168DE127EDF951F8C135A91E2560"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "CD1AC8BFC2C01FC3C2DAF26628D1CC820FD0CF6C",
            "timestamp": "2021-05-04T18:02:48.272819Z",
            "signature": "oc6ZVo4e7ETD0eLWok1egYo9eqPx+zpApzuO2TG2lnlJ+HEBXlGaKtR8lBd0jYEmGGdXU+Iv7imvFL+IeQQ3Aw=="
          }
        ]
      }
    }
  }
} `

const TX_MSG_CONNECTION_OPEN_INIT_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "6",
    "txs_results": [
      {
        "code": 0,
        "data": "ChYKFGNvbm5lY3Rpb25fb3Blbl9pbml0",
        "log": "[{\"events\":[{\"type\":\"connection_open_init\",\"attributes\":[{\"key\":\"connection_id\",\"value\":\"connection-0\"},{\"key\":\"client_id\",\"value\":\"07-tendermint-0\"},{\"key\":\"counterparty_client_id\",\"value\":\"07-tendermint-0\"},{\"key\":\"counterparty_connection_id\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"connection_open_init\"},{\"key\":\"module\",\"value\":\"ibc_connection\"}]}]}]",
        "info": "",
        "gas_wanted": "3000000",
        "gas_used": "62991",
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
                "value": "Y3JvMWdkc3dybXd0emd2M2t2ZjI4bHZ0dDdxdjdxN215em1uNDY2cjNm",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MTAwMGJhc2Vjcm8=",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMWdkc3dybXd0emd2M2t2ZjI4bHZ0dDdxdjdxN215em1uNDY2cjNm",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "Y29ubmVjdGlvbl9vcGVuX2luaXQ=",
                "index": true
              }
            ]
          },
          {
            "type": "connection_open_init",
            "attributes": [
              {
                "key": "Y29ubmVjdGlvbl9pZA==",
                "value": "Y29ubmVjdGlvbi0w",
                "index": true
              },
              {
                "key": "Y2xpZW50X2lk",
                "value": "MDctdGVuZGVybWludC0w",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X2NsaWVudF9pZA==",
                "value": "MDctdGVuZGVybWludC0w",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X2Nvbm5lY3Rpb25faWQ=",
                "value": null,
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "aWJjX2Nvbm5lY3Rpb24=",
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
            "value": "MjA2MTc4NzYwNTk0NzZiYXNlY3Jv",
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
            "value": "MC4wMDAwMDk5OTAwMDg5NjExNzU=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMzAwMDAxMjM1ODE2OTk3ODc=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTMwMTMwMTM3MTA2OTA2NDY3MzE2LjEzNDQzNDkyNDEyMjE0MjI4Nw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc4NzYwNTk0NzY=",
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
            "value": "MjA2MTc4NzYwNjA0NzZiYXNlY3Jv",
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
            "value": "MTAzMDg5MzgwMzAyMy44MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMDg5MzgwMzAyLjM4MDAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMDg5MzgwMzAyMy44MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNzQ2MjQ3MzYyNC4yNjgwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNzQ2MjQ3MzYyNDIuNjgwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": null,
    "validator_updates": null,
    "consensus_param_updates": {
      "block": {
        "max_bytes": "22020096",
        "max_gas": "-1"
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

const TX_MSG_CONNECTION_OPEN_INIT_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/ibc.core.connection.v1.MsgConnectionOpenInit",
                    "client_id": "07-tendermint-0",
                    "counterparty": {
                        "client_id": "07-tendermint-0",
                        "connection_id": "",
                        "prefix": {
                            "key_prefix": "aWJj"
                        }
                    },
                    "version": {
                        "identifier": "1",
                        "features": [
                            "ORDER_ORDERED",
                            "ORDER_UNORDERED"
                        ]
                    },
                    "delay_period": "0",
                    "signer": "cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f"
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
                        "key": "Ay3F5FOQif8MrhCypdaMnizCw4zKh3WZMB6OYdD04+eI"
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
                        "amount": "1000"
                    }
                ],
                "gas_limit": "3000000",
                "payer": "",
                "granter": ""
            }
        },
        "signatures": [
            "YWU7jgmoK3tPiyHQRAsgLxZQowKqun4PsoouNjVEfdNV0PgOEjHafVmf2sc2s3Zci/XeXsmRJCzVt9yTwZQI3Q=="
        ]
    },
    "tx_response": {
        "height": "6",
        "txhash": "F2B7D61BA783E6CDD9FE5825EBF7770688F6F45C482CB78ACB51E84B06FC643E",
        "codespace": "",
        "code": 0,
        "data": "ChYKFGNvbm5lY3Rpb25fb3Blbl9pbml0",
        "raw_log": "[{\"events\":[{\"type\":\"connection_open_init\",\"attributes\":[{\"key\":\"connection_id\",\"value\":\"connection-0\"},{\"key\":\"client_id\",\"value\":\"07-tendermint-0\"},{\"key\":\"counterparty_client_id\",\"value\":\"07-tendermint-0\"},{\"key\":\"counterparty_connection_id\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"connection_open_init\"},{\"key\":\"module\",\"value\":\"ibc_connection\"}]}]}]",
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
                                "value": "Y3JvMWdkc3dybXd0emd2M2t2ZjI4bHZ0dDdxdjdxN215em1uNDY2cjNm",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MTAwMGJhc2Vjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMWdkc3dybXd0emd2M2t2ZjI4bHZ0dDdxdjdxN215em1uNDY2cjNm",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "Y29ubmVjdGlvbl9vcGVuX2luaXQ=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "connection_open_init",
                        "attributes": [
                            {
                                "key": "Y29ubmVjdGlvbl9pZA==",
                                "value": "Y29ubmVjdGlvbi0w",
                                "index": true
                            },
                            {
                                "key": "Y2xpZW50X2lk",
                                "value": "MDctdGVuZGVybWludC0w",
                                "index": true
                            },
                            {
                                "key": "Y291bnRlcnBhcnR5X2NsaWVudF9pZA==",
                                "value": "MDctdGVuZGVybWludC0w",
                                "index": true
                            },
                            {
                                "key": "Y291bnRlcnBhcnR5X2Nvbm5lY3Rpb25faWQ=",
                                "value": null,
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "aWJjX2Nvbm5lY3Rpb24=",
                                "index": true
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "62991",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/ibc.core.connection.v1.MsgConnectionOpenInit",
                        "client_id": "07-tendermint-0",
                        "counterparty": {
                            "client_id": "07-tendermint-0",
                            "connection_id": "",
                            "prefix": {
                                "key_prefix": "aWJj"
                            }
                        },
                        "version": {
                            "identifier": "1",
                            "features": [
                                "ORDER_ORDERED",
                                "ORDER_UNORDERED"
                            ]
                        },
                        "delay_period": "0",
                        "signer": "cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f"
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
                            "key": "Ay3F5FOQif8MrhCypdaMnizCw4zKh3WZMB6OYdD04+eI"
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
                            "amount": "1000"
                        }
                    ],
                    "gas_limit": "3000000",
                    "payer": "",
                    "granter": ""
                }
            },
            "signatures": [
                "YWU7jgmoK3tPiyHQRAsgLxZQowKqun4PsoouNjVEfdNV0PgOEjHafVmf2sc2s3Zci/XeXsmRJCzVt9yTwZQI3Q=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
