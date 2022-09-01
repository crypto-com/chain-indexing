package usecase_parser_test

const TX_MSG_FUND_COMMUNITY_POOL_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "60CD10C5A635F8573F3C9D3293F3290ADF1D0615882DA018387F7B3FAB5ACF73",
      "parts": {
        "total": 1,
        "hash": "BC293207B50E8DFEE77B09DCC389D4287E601DE48C287148B542E4046F38D999"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-1",
        "height": "460662",
        "time": "2020-11-18T20:09:35.616503719Z",
        "last_block_id": {
          "hash": "72E2F52F933094510E1A8D0BBA4124C4CDDDA2B489C4A6AC4CB1716278A1778B",
          "parts": {
            "total": 1,
            "hash": "F698E767228D8C0270C602F6AE5E506BD8790330D60B09C80F7566671A9D182B"
          }
        },
        "last_commit_hash": "6A7D7D14F6932F2F8ADD8769D5E37678487A2F9F8B84605CBBE9B3D08168BF27",
        "data_hash": "65512958E8E1B51751F48A208C9F771432B18A07791CD68FBE201199C417A917",
        "validators_hash": "6D0A0787FF684D8AFBBDA99196F57A9F0BD044442FADAA59E21ADB6965DC1E1C",
        "next_validators_hash": "6D0A0787FF684D8AFBBDA99196F57A9F0BD044442FADAA59E21ADB6965DC1E1C",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "9F669CF4E2EAD60930F2E2B9F204B52F8BDF100692FD0DE26E186F941D60DFD6",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "73F3DAAE50A59F67990065ECAA33BF89F49C44B0"
      },
      "data": {
        "txs": [
          "CnMKcQoxL2Nvc21vcy5kaXN0cmlidXRpb24udjFiZXRhMS5Nc2dGdW5kQ29tbXVuaXR5UG9vbBI8Cg0KCGJhc2V0Y3JvEgExEit0Y3JvMWZtcHJtMHNqeTZsejlsbHY3cmx0bjB2MmF6endjd3p2azJsc3luElgKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQNZoVS6IQxInaRmJqTWMcb4pHGi+9o0IWTdX8ShWJAfJhIECgIIARgJEgQQwJoMGkDIo5p+5iJlLqqvJA0Zqt8B/V8LHaY26euOhVSQS1fYeXukVBka3z0+CuUlXmU732wTxx+EZwW79eFB+8g5zXX7"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "460661",
        "round": 0,
        "block_id": {
          "hash": "72E2F52F933094510E1A8D0BBA4124C4CDDDA2B489C4A6AC4CB1716278A1778B",
          "parts": {
            "total": 1,
            "hash": "F698E767228D8C0270C602F6AE5E506BD8790330D60B09C80F7566671A9D182B"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A1E8AAEBBC82929B852748734BA39D67A62F201B",
            "timestamp": "2020-11-18T20:09:35.504063977Z",
            "signature": "RVnU4kvzmFBqtT4qYv1e71LzUpzwWpxHUHmiH/Un08tr7mtz4nGsyYKqwT7M7kclVP6ECD1wAsDD327byugcBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3705DA4F2E53A09025DAA8E6581EFCE851811914",
            "timestamp": "2020-11-18T20:09:35.621342924Z",
            "signature": "NRAf9huSCu13zIKLQ2VwfgWLS7hB6pr9I+aLFcJCjIl/s2Dt/WUQJYCPrsUnbWSoVSeQZlzU6A9hQopNc1xVAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4B68F098199E7F565B02EF58115FB3CB9BAD52B0",
            "timestamp": "2020-11-18T20:09:35.627295831Z",
            "signature": "04f0C7KJG39vbKHMfMZeMh0V9YXGD/kC+p17EbpSwBU6DT1pi1TLcQ27vKhCkFLAeGou5yUjJ+qAhs0JGdhkAQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "504C0C3FE72728946911C7956E1B012784446B64",
            "timestamp": "2020-11-18T20:09:35.486810391Z",
            "signature": "m1raCUmoKteeNiGeRJwTnetcUiZsia7CG5dVaS+gy7RyGHg1mEJiaZnWUfIdQFOmkQhUCnp8i4cJqZHecYH/BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "95CDD1C2F0E79F62745D17A90D9A7B138DC8F922",
            "timestamp": "2020-11-18T20:09:35.616503719Z",
            "signature": "ksUPYOxtx0MGThgFRAATjV5YOzF3LHgBzJDruo4wG5tZ8oKQwV+Jb45FcJyMnO8rjEY5W7ERmI6nrzw0j6b7BA=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_FUND_COMMUNITY_POOL_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "460662",
    "txs_results": [
      {
        "code": 0,
        "data": "ChUKE2Z1bmRfY29tbXVuaXR5X3Bvb2w=",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"fund_community_pool\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"},{\"key\":\"amount\",\"value\":\"1basetcro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "54151",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "ZnVuZF9jb21tdW5pdHlfcG9vbA==",
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
                "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MWJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
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
            "value": "MTc2OTY5ODQ1MzBiYXNldGNybw==",
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
            "value": "MC4wMDEwMTYyODY3MDkzMDA3MDY=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTM5NDgwMzc2OTg5OTMyMDQ=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTExNjk0ODcxODA1OTAxNzY4LjYyOTY4NDk1Njk1OTYyODkwMA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc2OTY5ODQ1MzA=",
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
            "value": "MTc2OTY5ODQ1MzBiYXNldGNybw==",
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
            "value": "ODg0ODQ5MjI2LjUwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeGdkMDV2dWZuY2FmeDh0Y25zdjc3dWN1bWhoMHV6OHh0N2Q1N2c=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODg0ODQ5MjIuNjUwMDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeGdkMDV2dWZuY2FmeDh0Y25zdjc3dWN1bWhoMHV6OHh0N2Q1N2c=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODg0ODQ5MjI2LjUwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeGdkMDV2dWZuY2FmeDh0Y25zdjc3dWN1bWhoMHV6OHh0N2Q1N2c=",
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

const TX_MSG_FUND_COMMUNITY_POOL_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.distribution.v1beta1.MsgFundCommunityPool",
                    "amount": [
                        {
                            "denom": "basetcro",
                            "amount": "1"
                        }
                    ],
                    "depositor": "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
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
                    "sequence": "9"
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
            "yKOafuYiZS6qryQNGarfAf1fCx2mNunrjoVUkEtX2Hl7pFQZGt89PgrlJV5lO99sE8cfhGcFu/XhQfvIOc11+w=="
        ]
    },
    "tx_response": {
        "height": "460662",
        "txhash": "933052FD68B10549312F3CBA9AF4F4CC77536BE5ECD335638DD36ECCE681201E",
        "codespace": "",
        "code": 0,
        "data": "ChUKE2Z1bmRfY29tbXVuaXR5X3Bvb2w=",
        "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"fund_community_pool\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"},{\"key\":\"amount\",\"value\":\"1basetcro\"}]}]}]",
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
                                "value": "ZnVuZF9jb21tdW5pdHlfcG9vbA==",
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
                                "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MWJhc2V0Y3Jv",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
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
        "gas_used": "54151",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.distribution.v1beta1.MsgFundCommunityPool",
                        "amount": [
                            {
                                "denom": "basetcro",
                                "amount": "1"
                            }
                        ],
                        "depositor": "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
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
                        "sequence": "9"
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
                "yKOafuYiZS6qryQNGarfAf1fCx2mNunrjoVUkEtX2Hl7pFQZGt89PgrlJV5lO99sE8cfhGcFu/XhQfvIOc11+w=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
