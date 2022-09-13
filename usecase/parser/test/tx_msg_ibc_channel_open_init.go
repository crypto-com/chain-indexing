package usecase_parser_test

const TX_MSG_CHANNEL_OPEN_INIT_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "A895B3B2DF1F92210B6259B345A5A5266FDA0880BE08034D3CEB7A0909603EC1",
      "parts": {
        "total": 1,
        "hash": "FBA5CA67D0435F0CEEBE415C90B0E5774AEA0FF470A477E4B1B9A4669BC1AE68"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "devnet-1",
        "height": "14",
        "time": "2021-07-04T09:35:25.884542Z",
        "last_block_id": {
          "hash": "6FB87900C9BCFB5E8C42AA74DB737946915C074B63DDE5C6C4CE4A1374711827",
          "parts": {
            "total": 1,
            "hash": "56B27C22C05708E412DE2BC03E6835CE27685AE4CBCEDB1A58126A28D6880933"
          }
        },
        "last_commit_hash": "58BA9F763CC0555F1A1FA75D261BFCF740661DC2200924909BA4A266A019EBFA",
        "data_hash": "FD7FDCF98488680DB546DB13FD8B8B93E086F10AB5AE4ED82190E92A5584C88D",
        "validators_hash": "C3BE664CFB455C44053CDFDFB7DE2A6483B6C196245455BF453D3878C580934B",
        "next_validators_hash": "C3BE664CFB455C44053CDFDFB7DE2A6483B6C196245455BF453D3878C580934B",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "05998AC1C7A51DE1A39F04787454B20C7051E873DE743E2AA63A6F375F87E4CF",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "FDB7CAC90633FDEC2964A941CC815B5A3B84F0F5"
      },
      "data": {
        "txs": [
          "Co0BCooBCicvaWJjLmNvcmUuY2hhbm5lbC52MS5Nc2dDaGFubmVsT3BlbkluaXQSXwoIdHJhbnNmZXISJwgBEAEaCgoIdHJhbnNmZXIiDGNvbm5lY3Rpb24tMCoHaWNzMjAtMRoqY3JvMTBzbmhsdmtwdWM0eGhxODJ1eWc1ZXgyZWV6bW1mNWVkNXRtcXN2EmoKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQNo3I6+deelJEhkAIkyMG6qWmXpt0W7sh9lFcjqCsy26xIECgIIARgEEhYKDwoHYmFzZWNybxIEMTAwMBDAjbcBGkDzUlxlSssHXB5+WgeEuA/wafkr9NXtGyUguIPURTcgfG3FYJGrxHR1yH7b+6rwlSFdMh/H4hDoGqqOx3PqDnzb"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "13",
        "round": 0,
        "block_id": {
          "hash": "6FB87900C9BCFB5E8C42AA74DB737946915C074B63DDE5C6C4CE4A1374711827",
          "parts": {
            "total": 1,
            "hash": "56B27C22C05708E412DE2BC03E6835CE27685AE4CBCEDB1A58126A28D6880933"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "FDB7CAC90633FDEC2964A941CC815B5A3B84F0F5",
            "timestamp": "2021-07-04T09:35:25.884542Z",
            "signature": "0RMosAoVpxGYL3wwrEEshp/Fcy7sz7fCV9hJHZJTXXvJrDfoULPLfANnw8LrI3dP8xf8quLYCNweo0r3O12zAQ=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_CHANNEL_OPEN_INIT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "14",
    "txs_results": [
      {
        "code": 0,
        "data": "ChMKEWNoYW5uZWxfb3Blbl9pbml0",
        "log": "[{\"events\":[{\"type\":\"channel_open_init\",\"attributes\":[{\"key\":\"port_id\",\"value\":\"transfer\"},{\"key\":\"channel_id\",\"value\":\"channel-0\"},{\"key\":\"counterparty_port_id\",\"value\":\"transfer\"},{\"key\":\"counterparty_channel_id\"},{\"key\":\"connection_id\",\"value\":\"connection-0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"channel_open_init\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]}]}]",
        "info": "",
        "gas_wanted": "3000000",
        "gas_used": "96471",
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
                "value": "Y3JvMTBzbmhsdmtwdWM0eGhxODJ1eWc1ZXgyZWV6bW1mNWVkNXRtcXN2",
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
                "value": "Y3JvMTBzbmhsdmtwdWM0eGhxODJ1eWc1ZXgyZWV6bW1mNWVkNXRtcXN2",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "Y2hhbm5lbF9vcGVuX2luaXQ=",
                "index": true
              }
            ]
          },
          {
            "type": "channel_open_init",
            "attributes": [
              {
                "key": "cG9ydF9pZA==",
                "value": "dHJhbnNmZXI=",
                "index": true
              },
              {
                "key": "Y2hhbm5lbF9pZA==",
                "value": "Y2hhbm5lbC0w",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X3BvcnRfaWQ=",
                "value": "dHJhbnNmZXI=",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X2NoYW5uZWxfaWQ=",
                "value": null,
                "index": true
              },
              {
                "key": "Y29ubmVjdGlvbl9pZA==",
                "value": "Y29ubmVjdGlvbi0w",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "aWJjX2NoYW5uZWw=",
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
            "value": "MjA2MTc5MDU1OTAwODdiYXNlY3Jv",
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
            "value": "MC4wMDAwMDk5OTAwMDczMTUwMzg=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMzAwMDAyODgzNTcyOTk1MDc=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTMwMTMwMzIzNDg5OTUwODc0NDczLjcxNDUyMDk0NTUzMTgwMzU3Mw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc5MDU1OTAwODc=",
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
            "value": "MjA2MTc5MDU1OTAwODdiYXNlY3Jv",
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
            "value": "MTAzMDg5NTI3OTUwNC4zNTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMDg5NTI3OTUwLjQzNTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMDg5NTI3OTUwNC4zNTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNzQ2NTIxOTg3OC4wOTEwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNzQ2NTIxOTg3ODAuOTEwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
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
}
`

const TX_MSG_CHANNEL_OPEN_INIT_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/ibc.core.channel.v1.MsgChannelOpenInit",
                    "port_id": "transfer",
                    "channel": {
                        "state": "STATE_INIT",
                        "ordering": "ORDER_UNORDERED",
                        "counterparty": {
                            "port_id": "transfer",
                            "channel_id": ""
                        },
                        "connection_hops": [
                            "connection-0"
                        ],
                        "version": "ics20-1"
                    },
                    "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
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
                        "key": "A2jcjr5156UkSGQAiTIwbqpaZem3RbuyH2UVyOoKzLbr"
                    },
                    "mode_info": {
                        "single": {
                            "mode": "SIGN_MODE_DIRECT"
                        }
                    },
                    "sequence": "4"
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
            "81JcZUrLB1wefloHhLgP8Gn5K/TV7RslILiD1EU3IHxtxWCRq8R0dch+2/uq8JUhXTIfx+IQ6Bqqjsdz6g582w=="
        ]
    },
    "tx_response": {
        "height": "14",
        "txhash": "310C7A5DE69CABB2175A2CBD417A2BFBA105C030941663F3F8809BBA2A6D810D",
        "codespace": "",
        "code": 0,
        "data": "ChMKEWNoYW5uZWxfb3Blbl9pbml0",
        "raw_log": "[{\"events\":[{\"type\":\"channel_open_init\",\"attributes\":[{\"key\":\"port_id\",\"value\":\"transfer\"},{\"key\":\"channel_id\",\"value\":\"channel-0\"},{\"key\":\"counterparty_port_id\",\"value\":\"transfer\"},{\"key\":\"counterparty_channel_id\"},{\"key\":\"connection_id\",\"value\":\"connection-0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"channel_open_init\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]}]}]",
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
                                "value": "Y3JvMTBzbmhsdmtwdWM0eGhxODJ1eWc1ZXgyZWV6bW1mNWVkNXRtcXN2",
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
                                "value": "Y3JvMTBzbmhsdmtwdWM0eGhxODJ1eWc1ZXgyZWV6bW1mNWVkNXRtcXN2",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "Y2hhbm5lbF9vcGVuX2luaXQ=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "channel_open_init",
                        "attributes": [
                            {
                                "key": "cG9ydF9pZA==",
                                "value": "dHJhbnNmZXI=",
                                "index": true
                            },
                            {
                                "key": "Y2hhbm5lbF9pZA==",
                                "value": "Y2hhbm5lbC0w",
                                "index": true
                            },
                            {
                                "key": "Y291bnRlcnBhcnR5X3BvcnRfaWQ=",
                                "value": "dHJhbnNmZXI=",
                                "index": true
                            },
                            {
                                "key": "Y291bnRlcnBhcnR5X2NoYW5uZWxfaWQ=",
                                "value": null,
                                "index": true
                            },
                            {
                                "key": "Y29ubmVjdGlvbl9pZA==",
                                "value": "Y29ubmVjdGlvbi0w",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "aWJjX2NoYW5uZWw=",
                                "index": true
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "3000000",
        "gas_used": "59471",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/ibc.core.channel.v1.MsgChannelOpenInit",
                        "port_id": "transfer",
                        "channel": {
                            "state": "STATE_INIT",
                            "ordering": "ORDER_UNORDERED",
                            "counterparty": {
                                "port_id": "transfer",
                                "channel_id": ""
                            },
                            "connection_hops": [
                                "connection-0"
                            ],
                            "version": "ics20-1"
                        },
                        "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
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
                            "key": "A2jcjr5156UkSGQAiTIwbqpaZem3RbuyH2UVyOoKzLbr"
                        },
                        "mode_info": {
                            "single": {
                                "mode": "SIGN_MODE_DIRECT"
                            }
                        },
                        "sequence": "4"
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
                "81JcZUrLB1wefloHhLgP8Gn5K/TV7RslILiD1EU3IHxtxWCRq8R0dch+2/uq8JUhXTIfx+IQ6Bqqjsdz6g582w=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
