package usecase_parser_test

const TX_MSG_WITHDRAW_DELEGATOR_REWARD_NO_REWARD_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "7FC94E612BF874E6A87783529AF265AC034750E92C72B6A9A0D2FD19292F7FB4",
      "parts": {
        "total": 1,
        "hash": "FC61E39E14D7ACAEDD380E07F19D8A3A367C5F4594794C2C42CBAE2846C5E109"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-1",
        "height": "435599",
        "time": "2020-11-16T21:06:00.392846583Z",
        "last_block_id": {
          "hash": "37A9EEB1C60691B6CF929F37E35ECE9530EEA23B3AC3762E25AF69D5668F126B",
          "parts": {
            "total": 1,
            "hash": "646878FDA9D84B676A80D275B29E89F9DEA561EC693DAFAFC2793561E52D3A1D"
          }
        },
        "last_commit_hash": "161177EC07B729810459EA56956E88E0C4D8F72D617E8D0158AADFA5483DB48F",
        "data_hash": "712ADA81EED278138805D71FCB7E98625CC5AF6CBB033FD1E4A7D625497766E4",
        "validators_hash": "DE6B9984D709A78E9725D9192FDAFEF83132D7AEEEBE476B4F72182E560F2FDC",
        "next_validators_hash": "DE6B9984D709A78E9725D9192FDAFEF83132D7AEEEBE476B4F72182E560F2FDC",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "B70C1FE86204AF1F37B2C010BAFA04B50187F862C68C1B810BE128DAE268AB21",
        "last_results_hash": "C4E25D9A39D00D542B8E66755589A470DA76F42DEB503F5EE05FCD4AEE5FCCE5",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "DB15AEDC38A039940170D75A8920255FD8C9AAFF"
      },
      "data": {
        "txs": [
          "CpwBCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xZm1wcm0wc2p5Nmx6OWxsdjdybHRuMHYyYXp6d2N3enZrMmxzeW4SL3Rjcm9jbmNsMTVncmZ0Zzg4bDBnZHc0bWc5dDlwd25sMHBkZTJhc2p6ZWt6MGVrElgKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQNZoVS6IQxInaRmJqTWMcb4pHGi+9o0IWTdX8ShWJAfJhIECgIIARgFEgQQwJoMGkB2gelxAgaxpChFQvb3uKmnZawcUuTfowjb/3NbKvformkLncqBrGhx/jNzjfhccI3A+bsZXN/FoNA0zr6OrkVQ"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "435598",
        "round": 0,
        "block_id": {
          "hash": "37A9EEB1C60691B6CF929F37E35ECE9530EEA23B3AC3762E25AF69D5668F126B",
          "parts": {
            "total": 1,
            "hash": "646878FDA9D84B676A80D275B29E89F9DEA561EC693DAFAFC2793561E52D3A1D"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A1E8AAEBBC82929B852748734BA39D67A62F201B",
            "timestamp": "2020-11-16T21:06:00.341805089Z",
            "signature": "SCQjjyZr81CnqYqFz4yQMkYMFpfs5X4juiWvOZ4noZQXhwtshfbB+/jFzwP1QdIb+8CpxSHmo0pE/LeoJF8cCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3705DA4F2E53A09025DAA8E6581EFCE851811914",
            "timestamp": "2020-11-16T21:06:00.464899396Z",
            "signature": "UigH30/tjH47gQ7i/tk+CLJj1kb7DH2M2aajWRIShxp3ziuLjTtDRn/MUqGxT7ePM7MuBRSZAj5kbq9y2DuJBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4B68F098199E7F565B02EF58115FB3CB9BAD52B0",
            "timestamp": "2020-11-16T21:06:00.361755927Z",
            "signature": "K8az4pvEyAKKJIzgXKSIBBdtDx7nzXDXHCdKlTwpy4X4tuhqm4rKsq7HvF69NvQ3Dd0C6vV9Wvc+FZusgv4fDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "504C0C3FE72728946911C7956E1B012784446B64",
            "timestamp": "2020-11-16T21:06:00.271715841Z",
            "signature": "M3tKpe3bNA245u+M6MQc3xdly24Tf3u9msRDf+NziZl4m0MuCO7sFAHoSTtDOCFkg3a5LzyiPXipINkPtfkDAQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "95CDD1C2F0E79F62745D17A90D9A7B138DC8F922",
            "timestamp": "2020-11-16T21:06:00.418579053Z",
            "signature": "oBjY4cPDXT68aYrJe9Bh/1FHGxP0m9imC8gJ5YWMrJ62PeLWePxuGQ4PrQqf+9Dh7OctfOmISfhLuI2nIRu4CA=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_WITHDRAW_DELEGATOR_REWARD_NO_REWARD_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "435599",
    "txs_results": [{
        "code": 0,
        "data": "ChsKGXdpdGhkcmF3X2RlbGVnYXRvcl9yZXdhcmQ=",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "81466",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "d2l0aGRyYXdfZGVsZWdhdG9yX3Jld2FyZA==",
                "index": true
              }
            ]
          },
          {
            "type": "withdraw_rewards",
            "attributes": [
              {
                "key": "YW1vdW50",
                "value": null,
                "index": true
              },
              {
                "key": "dmFsaWRhdG9y",
                "value": "dGNyb2NuY2wxNWdyZnRnODhsMGdkdzRtZzl0OXB3bmwwcGRlMmFzanpla3owZWs=",
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
            "value": "MTc2MzA2MDg4NzViYXNldGNybw==",
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
            "value": "MC4wMDA5NjUyMjY1ODU1OTU4NDQ=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTM4OTY0OTExMjcwMjY5MDc=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTExMjc1OTQwNTMxNzE1OTUyLjY3MTEyOTEzMjY0MzQyOTgxMQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc2MzA2MDg4NzU=",
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
            "value": "MTc2MzA2NDg4NzViYXNldGNybw==",
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
            "value": "ODgxNTMyNDQzLjc1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZXpwejRhODhjcTdxdmthZGR3bDdoc3BocTc0eG56NHNkc2d6ZzI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODgxNTMyNDQuMzc1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZXpwejRhODhjcTdxdmthZGR3bDdoc3BocTc0eG56NHNkc2d6ZzI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODgxNTMyNDQzLjc1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZXpwejRhODhjcTdxdmthZGR3bDdoc3BocTc0eG56NHNkc2d6ZzI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDU3Mjg0ODczLjc3NTcwMjkxNTI1ODU3NjA3MmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeHdkM2s4eHRlcmRlZnQzbnhxZzkyc3pocHo2dng0M3FzcGRwdzY=",
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

const TX_MSG_WITHDRAW_DELEGATOR_REWARD_NO_REWARD_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                    "delegator_address": "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
                    "validator_address": "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek"
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
                    "sequence": "5"
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
            "doHpcQIGsaQoRUL297ipp2WsHFLk36MI2/9zWyr36K5pC53Kgaxocf4zc434XHCNwPm7GVzfxaDQNM6+jq5FUA=="
        ]
    },
    "tx_response": {
        "height": "435599",
        "txhash": "3643A4CA41EC52BCB5B10DB32EC9867B2FA6B6A7C48B4DE9D45E6EDBC39B31B5",
        "codespace": "",
        "code": 0,
        "data": "ChsKGXdpdGhkcmF3X2RlbGVnYXRvcl9yZXdhcmQ=",
        "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek\"}]}]}]",
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
                                "value": "d2l0aGRyYXdfZGVsZWdhdG9yX3Jld2FyZA==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "withdraw_rewards",
                        "attributes": [
                            {
                                "key": "YW1vdW50",
                                "value": null,
                                "index": true
                            },
                            {
                                "key": "dmFsaWRhdG9y",
                                "value": "dGNyb2NuY2wxNWdyZnRnODhsMGdkdzRtZzl0OXB3bmwwcGRlMmFzanpla3owZWs=",
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
        "gas_used": "81466",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                        "delegator_address": "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
                        "validator_address": "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek"
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
                        "sequence": "5"
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
                "doHpcQIGsaQoRUL297ipp2WsHFLk36MI2/9zWyr36K5pC53Kgaxocf4zc434XHCNwPm7GVzfxaDQNM6+jq5FUA=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
