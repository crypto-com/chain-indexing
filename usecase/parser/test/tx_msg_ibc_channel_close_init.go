package usecase_parser_test

const TX_MSG_CHANNEL_CLOSE_INIT_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "F6EA8DBCF90CE918F09777E4D05B0333A1F041D2B4CF4033B9170D342EEBC678",
      "parts": {
        "total": 1,
        "hash": "191CB264DE0B49A9CB7562ADE2964A4DFBE91140EFDCE1C753CF0D218D05E053"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "devnet-2",
        "height": "23",
        "time": "2021-10-13T04:46:43.53024Z",
        "last_block_id": {
          "hash": "93A0B76E3B810E110A25E5D4B361661CC1A6DD3FDF34D9084EF01CE36AF72E12",
          "parts": {
            "total": 1,
            "hash": "FD17A5DC0B9E9BC36C82E419C0A92B78F99B1FEED1A4ACBD8739AEB66823ECEC"
          }
        },
        "last_commit_hash": "8A6325DA51A6AB5D2744F3D3FC7D157BEF652DF85F7D54810F45AE3F7880AE97",
        "data_hash": "0DBDBBBBF9AE56A0C87D7264124B8771AF0DE5AADA415F23382124DB9986A2AC",
        "validators_hash": "C3339989D95E5D67B7F93A541B5D65259B44DC274212D685CB327E5FAD17968F",
        "next_validators_hash": "C3339989D95E5D67B7F93A541B5D65259B44DC274212D685CB327E5FAD17968F",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "BA05E8ABF004D0464611E1BB9A8D556EBA64998AC2BB1A3039FADAE9C8F93875",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "E1BCA6465808844C534471E7D786FCFE1526C09E"
      },
      "data": {
        "txs": [
          "Cm8KbQooL2liYy5jb3JlLmNoYW5uZWwudjEuTXNnQ2hhbm5lbENsb3NlSW5pdBJBCgh0cmFuc2ZlchIJY2hhbm5lbC0xGipjcm8xdDd5azNkNG1lZWFxZjV6ZmVndjhwOTR3bGZocGNuc2Z0ejU1ZjcSbQpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA21jVocfixpkt0btHUSMNYeU7z07Lv7BxLQlIprpnHTCEgQKAggBGAkSGQoTCgdiYXNlY3JvEgg2NjM5MjAwMBDYhgQaQLv/DtPhoeLlRZKq2qoCG2TTu+0L3AAq4O4H1R5qb1MLCvoBppzA0Sp44fNl3GDAk3Q+uWHWxXTeCbEybS41dOc="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "22",
        "round": 0,
        "block_id": {
          "hash": "93A0B76E3B810E110A25E5D4B361661CC1A6DD3FDF34D9084EF01CE36AF72E12",
          "parts": {
            "total": 1,
            "hash": "FD17A5DC0B9E9BC36C82E419C0A92B78F99B1FEED1A4ACBD8739AEB66823ECEC"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "E1BCA6465808844C534471E7D786FCFE1526C09E",
            "timestamp": "2021-10-13T04:46:43.53024Z",
            "signature": "Y+76aQ4ut1MAaf00D1LfwXys/xfmIBnigU8f/Am46aXdpyaCrYIwmZiCNIoUCMHGvQcOVcfGDFGhOrOHl2XFDQ=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_CHANNEL_CLOSE_INIT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "23",
    "txs_results": [
      {
        "code": 0,
        "data": "CioKKC9pYmMuY29yZS5jaGFubmVsLnYxLk1zZ0NoYW5uZWxDbG9zZUluaXQ=",
        "log": "[{\"events\":[{\"type\":\"channel_close_init\",\"attributes\":[{\"key\":\"port_id\",\"value\":\"transfer\"},{\"key\":\"channel_id\",\"value\":\"channel-1\"},{\"key\":\"counterparty_port_id\",\"value\":\"transfer\"},{\"key\":\"counterparty_channel_id\",\"value\":\"channel-0\"},{\"key\":\"connection_id\",\"value\":\"connection-1\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/ibc.core.channel.v1.MsgChannelCloseInit\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]}]}]",
        "info": "",
        "gas_wanted": "66392",
        "gas_used": "60266",
        "events": [
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "Y3JvMXQ3eWszZDRtZWVhcWY1emZlZ3Y4cDk0d2xmaHBjbnNmdHo1NWY3",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NjYzOTIwMDBiYXNlY3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "cmVjZWl2ZXI=",
                "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NjYzOTIwMDBiYXNlY3Jv",
                "index": true
              }
            ]
          },
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
                "value": "Y3JvMXQ3eWszZDRtZWVhcWY1emZlZ3Y4cDk0d2xmaHBjbnNmdHo1NWY3",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NjYzOTIwMDBiYXNlY3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMXQ3eWszZDRtZWVhcWY1emZlZ3Y4cDk0d2xmaHBjbnNmdHo1NWY3",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "ZmVl",
                "value": "NjYzOTIwMDBiYXNlY3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "Y3JvMXQ3eWszZDRtZWVhcWY1emZlZ3Y4cDk0d2xmaHBjbnNmdHo1NWY3Lzk=",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "dS84TzArR2g0dVZGa3FyYXFnSWJaTk83N1F2Y0FDcmc3Z2ZWSG1wdlV3c0srZ0dtbk1EUktuamg4MlhjWU1DVGRENjVZZGJGZE40SnNUSnRMalYwNXc9PQ==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2liYy5jb3JlLmNoYW5uZWwudjEuTXNnQ2hhbm5lbENsb3NlSW5pdA==",
                "index": true
              }
            ]
          },
          {
            "type": "channel_close_init",
            "attributes": [
              {
                "key": "cG9ydF9pZA==",
                "value": "dHJhbnNmZXI=",
                "index": true
              },
              {
                "key": "Y2hhbm5lbF9pZA==",
                "value": "Y2hhbm5lbC0x",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X3BvcnRfaWQ=",
                "value": "dHJhbnNmZXI=",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X2NoYW5uZWxfaWQ=",
                "value": "Y2hhbm5lbC0w",
                "index": true
              },
              {
                "key": "Y29ubmVjdGlvbl9pZA==",
                "value": "Y29ubmVjdGlvbi0x",
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
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc5Mzg4MTIwNDFiYXNlY3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coinbase",
        "attributes": [
          {
            "key": "bWludGVy",
            "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc5Mzg4MTIwNDFiYXNlY3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc5Mzg4MTIwNDFiYXNlY3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc5Mzg4MTIwNDFiYXNlY3Jv",
            "index": true
          }
        ]
      },
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
            "value": "MjA2MTc5Mzg4MTIwNDFiYXNlY3Jv",
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
            "value": "MC4wMDAwMDk5OTAwMDU0NjMxMzM=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMzAwMDA0NzM3Mjk4NDkxOTI=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTMwMTMwNTMzMTcwOTc3NTE2NzUwLjA2MDYyODE0MTEyOTA1NjE4NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc5Mzg4MTIwNDE=",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc5Mzg4MTIwNDFiYXNlY3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc5Mzg4MTIwNDFiYXNlY3Jv",
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
            "value": "MjA2MTc5Mzg4MTIwNDFiYXNlY3Jv",
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
            "value": "MTAzMDg5Njk0MDYwMi4wNTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFmcDYwZDBxeGd1aGsyY2xtZDl4bXhkcG0wejlsM3lhMDc1ajZxdA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMDg5Njk0MDYwLjIwNTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFmcDYwZDBxeGd1aGsyY2xtZDl4bXhkcG0wejlsM3lhMDc1ajZxdA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMDg5Njk0MDYwMi4wNTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFmcDYwZDBxeGd1aGsyY2xtZDl4bXhkcG0wejlsM3lhMDc1ajZxdA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNzQ2ODMwOTUxOS44MTMwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFmcDYwZDBxeGd1aGsyY2xtZDl4bXhkcG0wejlsM3lhMDc1ajZxdA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNzQ2ODMwOTUxOTguMTMwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFmcDYwZDBxeGd1aGsyY2xtZDl4bXhkcG0wejlsM3lhMDc1ajZxdA==",
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
        "pub_key_types": ["ed25519"]
      }
    }
  }
}
`

const TX_MSG_CHANNEL_CLOSE_INIT_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/ibc.core.channel.v1.MsgChannelCloseInit",
                    "port_id": "transfer",
                    "channel_id": "channel-1",
                    "signer": "cro1t7yk3d4meeaqf5zfegv8p94wlfhpcnsftz55f7"
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
                        "key": "A21jVocfixpkt0btHUSMNYeU7z07Lv7BxLQlIprpnHTC"
                    },
                    "mode_info": {
                        "single": {
                            "mode": "SIGN_MODE_DIRECT"
                        }
                    },
                    "sequence": "9"
                }
            ],
            "fee": {
                "amount": [
                    {
                        "denom": "basecro",
                        "amount": "66392000"
                    }
                ],
                "gas_limit": "66392",
                "payer": "",
                "granter": ""
            }
        },
        "signatures": [
            "u/8O0+Gh4uVFkqraqgIbZNO77QvcACrg7gfVHmpvUwsK+gGmnMDRKnjh82XcYMCTdD65YdbFdN4JsTJtLjV05w=="
        ]
    },
    "tx_response": {
        "height": "23",
        "txhash": "3E491C5B2404FC3C3C15A59630729355ED398BFB617DA7D61C1B548E98955F8D",
        "codespace": "",
        "code": 0,
        "data": "CioKKC9pYmMuY29yZS5jaGFubmVsLnYxLk1zZ0NoYW5uZWxDbG9zZUluaXQ=",
        "raw_log": "[{\"events\":[{\"type\":\"channel_close_init\",\"attributes\":[{\"key\":\"port_id\",\"value\":\"transfer\"},{\"key\":\"channel_id\",\"value\":\"channel-1\"},{\"key\":\"counterparty_port_id\",\"value\":\"transfer\"},{\"key\":\"counterparty_channel_id\",\"value\":\"channel-0\"},{\"key\":\"connection_id\",\"value\":\"connection-1\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/ibc.core.channel.v1.MsgChannelCloseInit\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]}]}]",
        "logs": [
            {
                "msg_index": 0,
                "log": "",
                "events": [
                    {
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "c3BlbmRlcg==",
                                "value": "Y3JvMXQ3eWszZDRtZWVhcWY1emZlZ3Y4cDk0d2xmaHBjbnNmdHo1NWY3",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "NjYzOTIwMDBiYXNlY3Jv",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "cmVjZWl2ZXI=",
                                "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "NjYzOTIwMDBiYXNlY3Jv",
                                "index": true
                            }
                        ]
                    },
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
                                "value": "Y3JvMXQ3eWszZDRtZWVhcWY1emZlZ3Y4cDk0d2xmaHBjbnNmdHo1NWY3",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "NjYzOTIwMDBiYXNlY3Jv",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMXQ3eWszZDRtZWVhcWY1emZlZ3Y4cDk0d2xmaHBjbnNmdHo1NWY3",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "ZmVl",
                                "value": "NjYzOTIwMDBiYXNlY3Jv",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "YWNjX3NlcQ==",
                                "value": "Y3JvMXQ3eWszZDRtZWVhcWY1emZlZ3Y4cDk0d2xmaHBjbnNmdHo1NWY3Lzk=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "c2lnbmF0dXJl",
                                "value": "dS84TzArR2g0dVZGa3FyYXFnSWJaTk83N1F2Y0FDcmc3Z2ZWSG1wdlV3c0srZ0dtbk1EUktuamg4MlhjWU1DVGRENjVZZGJGZE40SnNUSnRMalYwNXc9PQ==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "L2liYy5jb3JlLmNoYW5uZWwudjEuTXNnQ2hhbm5lbENsb3NlSW5pdA==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "channel_close_init",
                        "attributes": [
                            {
                                "key": "cG9ydF9pZA==",
                                "value": "dHJhbnNmZXI=",
                                "index": true
                            },
                            {
                                "key": "Y2hhbm5lbF9pZA==",
                                "value": "Y2hhbm5lbC0x",
                                "index": true
                            },
                            {
                                "key": "Y291bnRlcnBhcnR5X3BvcnRfaWQ=",
                                "value": "dHJhbnNmZXI=",
                                "index": true
                            },
                            {
                                "key": "Y291bnRlcnBhcnR5X2NoYW5uZWxfaWQ=",
                                "value": "Y2hhbm5lbC0w",
                                "index": true
                            },
                            {
                                "key": "Y29ubmVjdGlvbl9pZA==",
                                "value": "Y29ubmVjdGlvbi0x",
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
        "gas_wanted": "200000",
        "gas_used": "60266",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/ibc.core.channel.v1.MsgChannelCloseInit",
                        "port_id": "transfer",
                        "channel_id": "channel-1",
                        "signer": "cro1t7yk3d4meeaqf5zfegv8p94wlfhpcnsftz55f7"
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
                            "key": "A21jVocfixpkt0btHUSMNYeU7z07Lv7BxLQlIprpnHTC"
                        },
                        "mode_info": {
                            "single": {
                                "mode": "SIGN_MODE_DIRECT"
                            }
                        },
                        "sequence": "9"
                    }
                ],
                "fee": {
                    "amount": [
                        {
                            "denom": "basecro",
                            "amount": "66392000"
                        }
                    ],
                    "gas_limit": "66392",
                    "payer": "",
                    "granter": ""
                }
            },
            "signatures": [
                "u/8O0+Gh4uVFkqraqgIbZNO77QvcACrg7gfVHmpvUwsK+gGmnMDRKnjh82XcYMCTdD65YdbFdN4JsTJtLjV05w=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
