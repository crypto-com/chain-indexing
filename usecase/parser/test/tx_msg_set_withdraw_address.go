package usecase_parser_test

const TX_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "A5896BF9DCB04D6CBCA913F66A493CD3C3C76569011F135F707936B81C3672AA",
      "parts": {
        "total": 1,
        "hash": "06D8588A347B9CC7C429E0267416F652CA3BF1827A0B347792BA19FCE6BE3A3C"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-1",
        "height": "460060",
        "time": "2020-11-18T19:01:53.897059486Z",
        "last_block_id": {
          "hash": "5F097398A5568089E7C0AF55C63FC28F51D56F717594EF4B0F49C5F2843774E8",
          "parts": {
            "total": 1,
            "hash": "731CA8FAFC4CEF6D154ACAC92878BFDE51EB5130F512BA332AEADBBAE8260B6A"
          }
        },
        "last_commit_hash": "C6753AD0C0781009181BDC5D792ECD87B7F602A7ACF29173E408C56FB7E21939",
        "data_hash": "5E65C976A1E13E91BB4824B9938C3514EA328D1AD885C5C066E5FEC58AAC0D18",
        "validators_hash": "591581CA8A17BD2D2A6CEE21754B88B4C5DC6B1AD140BF879A60E5E4D5CD6CCA",
        "next_validators_hash": "BCBDE8CC52DEE9553BBEA5BA7C600CFE496D73245F3D663E263DBCB2163F2BB2",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "80C5B2A2F07C6C3F3E86A04C5B739388F339B00B842B9723250B848F4D08EE4D",
        "last_results_hash": "4B870D4F09AC178B4743DA6FABFC946647474B246427BDB7071A10745FCFBC5F",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "3705DA4F2E53A09025DAA8E6581EFCE851811914"
      },
      "data": {
        "txs": [
          "CpMBCpABCjIvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1NldFdpdGhkcmF3QWRkcmVzcxJaCit0Y3JvMWZtcHJtMHNqeTZsejlsbHY3cmx0bjB2MmF6endjd3p2azJsc3luEit0Y3JvMTRtNWE0a3h0MmU4MnVxcXM1Z3RxemEyOWRtNXdxenlhMmp3OXNoElgKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQNZoVS6IQxInaRmJqTWMcb4pHGi+9o0IWTdX8ShWJAfJhIECgIIARgIEgQQwJoMGkBkcbgfCjwZeROqWnjuQouDqzQ+hXk4bTCNokkb/aIBUUzhIWZ5OeqXAgT7TpjdQpCJsXdSsyjzNjW/n9HB5aF7"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "460059",
        "round": 0,
        "block_id": {
          "hash": "5F097398A5568089E7C0AF55C63FC28F51D56F717594EF4B0F49C5F2843774E8",
          "parts": {
            "total": 1,
            "hash": "731CA8FAFC4CEF6D154ACAC92878BFDE51EB5130F512BA332AEADBBAE8260B6A"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A1E8AAEBBC82929B852748734BA39D67A62F201B",
            "timestamp": "2020-11-18T19:01:53.799393339Z",
            "signature": "mLitN1qi+FadtvOkowKgTPlrexnOagIYK+GTBrPEPIylWOCJTvcHm76mWknQ75+R5OE3/vAnedQw6fwZdv42Bw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3705DA4F2E53A09025DAA8E6581EFCE851811914",
            "timestamp": "2020-11-18T19:01:54.105797705Z",
            "signature": "+u7C0LH/1kyoztF6FHWJ/dpQcPYrX79qb2jl1WC9411kIeOpiMT6a3p5137aBaAvmvkRyASXjEgnYa1i4RMdBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4B68F098199E7F565B02EF58115FB3CB9BAD52B0",
            "timestamp": "2020-11-18T19:01:53.883167068Z",
            "signature": "VAPt0+S+aj4N0Z81a5sYXwGYI7pDkUO2j+KfsOQfHEj263HNsLpaX0mXT27Jnz33ai8AB/enxrxnv/8bv36FBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "504C0C3FE72728946911C7956E1B012784446B64",
            "timestamp": "2020-11-18T19:01:53.691731697Z",
            "signature": "BUdjw3VW1TS/ByWQ3ql5+bkc2optXTJ7iVF+xf6+LLhf8H2Py5tYMPmbN2AXovNPjwv+CHmhYN54ieJ9tRArBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "95CDD1C2F0E79F62745D17A90D9A7B138DC8F922",
            "timestamp": "2020-11-18T19:01:53.997379508Z",
            "signature": "fDubk5KNqdsDZZI5/TjvmuLg0A+Yd0JXhAiREMKx3T4qgb+fyrbByxRWc/vrqpT+EwWpb2HzyYxG48D8eXenBg=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "460060",
    "txs_results": [
      {
        "code": 0,
        "data": "ChYKFHNldF93aXRoZHJhd19hZGRyZXNz",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"set_withdraw_address\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"}]},{\"type\":\"set_withdraw_address\",\"attributes\":[{\"key\":\"withdraw_address\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "42013",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "c2V0X3dpdGhkcmF3X2FkZHJlc3M=",
                "index": true
              }
            ]
          },
          {
            "type": "set_withdraw_address",
            "attributes": [
              {
                "key": "d2l0aGRyYXdfYWRkcmVzcw==",
                "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "ZGlzdHJpYnV0aW9u",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
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
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc2OTUzOTAxNDZiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC4wMDEwMTUyNDc3NDQwNDcxMjI=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTM5NDY3OTk2MjM5ODUzNDg=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTExNjg0ODA4ODE0NTQ2MTIzLjUyNTU0NTAyOTY0MTczOTMzNg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc2OTUzOTAxNDY=",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODMzOXA0bA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc2OTU0MTAxNDZiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "proposer_reward",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODg0NzcwNTA3LjMwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdHh0OTMweHV4bGZrd2Y4a25laDV6eXRlMmNoN3dwdjczc3d4eTI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODg0NzcwNTAuNzMwMDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdHh0OTMweHV4bGZrd2Y4a25laDV6eXRlMmNoN3dwdjczc3d4eTI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODg0NzcwNTA3LjMwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdHh0OTMweHV4bGZrd2Y4a25laDV6eXRlMmNoN3dwdjczc3d4eTI=",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": null,
    "validator_updates": [
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "fAkI6G9XcnXjaYH6y91T4lYxnrXQ9t1cBm/A2DZl7j8="
            }
          }
        },
        "power": "157927637"
      }
    ],
    "consensus_param_updates": {
      "block": {
        "max_bytes": "22020096",
        "max_gas": "-1"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      }
    }
  }
}`

const TX_MSG_SET_WITHDRAW_ADDRESS_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
                    "delegator_address": "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
                    "withdraw_address": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"
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
                        "key": "A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m"
                    },
                    "mode_info": {
                        "single": {
                            "mode": "SIGN_MODE_DIRECT"
                        }
                    },
                    "sequence": "8"
                }
            ],
            "fee": {
                "amount": [],
                "gas_limit": "200000",
                "payer": "",
                "granter": ""
            }
        },
        "signatures": [
            "ZHG4Hwo8GXkTqlp47kKLg6s0PoV5OG0wjaJJG/2iAVFM4SFmeTnqlwIE+06Y3UKQibF3UrMo8zY1v5/RweWhew=="
        ]
    },
    "tx_response": {
        "height": "460060",
        "txhash": "9C2501310E18EE69A7FE5CA1A684A0701C43BEB1A8D91EDA80CC598C924F9CBE",
        "codespace": "",
        "code": 0,
        "data": "ChYKFHNldF93aXRoZHJhd19hZGRyZXNz",
        "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"set_withdraw_address\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"}]},{\"type\":\"set_withdraw_address\",\"attributes\":[{\"key\":\"withdraw_address\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"}]}]}]",
        "logs": [
            {
                "msg_index": 0,
                "log": "",
                "events": [
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "c2V0X3dpdGhkcmF3X2FkZHJlc3M=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "set_withdraw_address",
                        "attributes": [
                            {
                                "key": "d2l0aGRyYXdfYWRkcmVzcw==",
                                "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "ZGlzdHJpYnV0aW9u",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
                                "index": true
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "42013",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
                        "delegator_address": "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
                        "withdraw_address": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"
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
                            "key": "A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m"
                        },
                        "mode_info": {
                            "single": {
                                "mode": "SIGN_MODE_DIRECT"
                            }
                        },
                        "sequence": "8"
                    }
                ],
                "fee": {
                    "amount": [],
                    "gas_limit": "200000",
                    "payer": "",
                    "granter": ""
                }
            },
            "signatures": [
                "ZHG4Hwo8GXkTqlp47kKLg6s0PoV5OG0wjaJJG/2iAVFM4SFmeTnqlwIE+06Y3UKQibF3UrMo8zY1v5/RweWhew=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
