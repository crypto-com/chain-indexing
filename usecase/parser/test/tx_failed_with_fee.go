package usecase_parser_test

const TX_FAILED_WITH_FEE_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "653704613C8B5A2B0BF2F6834D76DCCEFCFE5B968536DC4DE92CDDAC6D2F8795",
      "parts": {
        "total": 1,
        "hash": "0B8F7A3700CB45ECF68C6847A2A090F437796AF87F3F25B76540247F7B10B194"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-1",
        "height": "420301",
        "time": "2020-11-14T15:12:04.515915178Z",
        "last_block_id": {
          "hash": "EB088172F526E5CA99194CCABF0F0DC005C5455739BDE64D21BAC466DE5482D1",
          "parts": {
            "total": 1,
            "hash": "C263B0B9E7078219078416DAB1FB254519744676EC66C5D04763C06C5826EEEA"
          }
        },
        "last_commit_hash": "4691F030493F5AE55435F4D4A8DEAD2BD9FE30EE40D42728D7DBC57A73D3857E",
        "data_hash": "0AA3804CCD08BFC0C53336768E56CC1ED248F68DD3657E67B24F0D346080401C",
        "validators_hash": "8CE0023D3327162430CC0FA93C4E3D8F46396A44C1CA2AA8D3B2738CE25F8434",
        "next_validators_hash": "8CE0023D3327162430CC0FA93C4E3D8F46396A44C1CA2AA8D3B2738CE25F8434",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "66F66F33AAC940F93C1EBD43E92E370C55EF19BDD62F122EE9934B50D772D27B",
        "last_results_hash": "5E0E284A94ADE5413F1EC168B4CA247C4241D0067011C4729B5FBC98697C1D90",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "504C0C3FE72728946911C7956E1B012784446B64"
      },
      "data": {
        "txs": [
          "CpUBCpIBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnIKK3Rjcm8xZmVxaDZhZDl5dGprcjc5a2prNW5obmw0dW4zd2V6MHludXJyd3YSK3Rjcm8xZmVxaDZhZDl5dGprcjc5a2prNW5obmw0dW4zd2V6MHludXJyd3YaFgoIYmFzZXRjcm8SCjEwMDAwMDAwMDASbgpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAx+Rgmd2ta8FxUOoFJ9Dvo3782nMWJzdYP0Jcyrk5XwOEgQKAggBGDsSGgoTCghiYXNldGNybxIHODAwMDAwMBCA6JImGkClcXvyfOzeWFKVOt6JNesyiqPEXTiSJ2tE7KPxsny+vE+/at95xSzHcgeD4/gBUc6y1rFqseI/vl9ZBIH0EGxH"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "406721",
        "round": 0,
        "block_id": {
          "hash": "EB088172F526E5CA99194CCABF0F0DC005C5455739BDE64D21BAC466DE5482D1",
          "parts": {
            "total": 1,
            "hash": "C263B0B9E7078219078416DAB1FB254519744676EC66C5D04763C06C5826EEEA"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A1E8AAEBBC82929B852748734BA39D67A62F201B",
            "timestamp": "2020-11-14T15:12:04.431274857Z",
            "signature": "urEhGwsgF/NUNzF/CTkngrewHAcHVxQXj4mne9YtRch9DTrO1dV6cRPGrJKB9ReJ8zOXWgasmpSsyt2QkIkBDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3705DA4F2E53A09025DAA8E6581EFCE851811914",
            "timestamp": "2020-11-14T15:12:04.529758908Z",
            "signature": "ge37zssTVOb8AuLHpxpJ2jWnLRt3buiDHjf45JsYSQB188skGhvu1ZsrMLk4b6rLBNZGIRUfnEPcQpCCHao4Bg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4B68F098199E7F565B02EF58115FB3CB9BAD52B0",
            "timestamp": "2020-11-14T15:12:04.525904924Z",
            "signature": "tkBVJixApww1iGlmy1htiriewgShOm3jSiOeRYBEd7+8uPOF8/mZBM1sKx68haafPwMwqbtcUW2ymC/2+raHDw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "504C0C3FE72728946911C7956E1B012784446B64",
            "timestamp": "2020-11-14T15:12:04.33393283Z",
            "signature": "wESm61iln7Aa+bWLuZNv1k13R661hR+6aWHd0/ALfPMGUqIdOREsPPQxxveCeSdIBjE6aIoIaYxb52xKHK/DAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "95CDD1C2F0E79F62745D17A90D9A7B138DC8F922",
            "timestamp": "2020-11-14T15:12:04.433355148Z",
            "signature": "wNTow7FmO439eiDI/CKQQNHU/ORmMeHCxBguSHQ/jX8u29ZAvkf5ibAkpa/rV3Wx+bZS4I5dgr1nleGK7T5WCQ=="
          }
        ]
      }
    }
  }
}`

const TX_FAILED_WITH_FEE_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "420301",
    "txs_results": [
      {
        "code": 11,
        "data": null,
        "log": "out of gas in location: WriteFlat; gasWanted: 80000000, gasUsed: 80150021: out of gas",
        "info": "",
        "gas_wanted": "80000000",
        "gas_used": "80150021",
        "events": [],
        "codespace": "sdk"
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
            "value": "MTc1NTQxMzc3NDNiYXNldGNybw==",
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
            "value": "MC4wMDA4ODEzODc2OTgyNzI2NTg=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTM4MzcwOTM5MTY2MTY5MzY=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTEwNzkzMjkxNDQ5MDQ3ODAwLjAyNTkyMjMxNDMzMzMyMjk3Ng==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc1NTQxMzc3NDM=",
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
            "value": "MTc1NTUxMzc2MjliYXNldGNybw==",
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
            "value": "ODc3NzU2ODgxLjQ1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODc3NzU2ODguMTQ1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
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
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODc3NzU2ODgxLjQ1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDY1MzYyMjk4LjU1NTY0NTAzMTQ4MDI2NTk3OGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeHdkM2s4eHRlcmRlZnQzbnhxZzkyc3pocHo2dng0M3FzcGRwdzY=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "OTMwNzI0NTk3LjExMTI5MDA2Mjk2MDUzMTk1NWJhc2V0Y3Jv",
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

const TX_FAILED_WITH_FEE_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.bank.v1beta1.MsgSend",
                    "from_address": "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
                    "to_address": "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
                    "amount": [
                        {
                            "denom": "basetcro",
                            "amount": "1000000000"
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
                        "key": "Ax+Rgmd2ta8FxUOoFJ9Dvo3782nMWJzdYP0Jcyrk5XwO"
                    },
                    "mode_info": {
                        "single": {
                            "mode": "SIGN_MODE_DIRECT"
                        }
                    },
                    "sequence": "59"
                }
            ],
            "fee": {
                "amount": [
                    {
                        "denom": "basetcro",
                        "amount": "8000000"
                    }
                ],
                "gas_limit": "80000000",
                "payer": "",
                "granter": ""
            }
        },
        "signatures": [
            "pXF78nzs3lhSlTreiTXrMoqjxF04kidrROyj8bJ8vrxPv2rfecUsx3IHg+P4AVHOstaxarHiP75fWQSB9BBsRw=="
        ]
    },
    "tx_response": {
        "height": "420301",
        "txhash": "2A2A64A310B3D0E84C9831F4353E188A6E63BF451975C859DF40C54047AC6324",
        "codespace": "sdk",
        "code": 11,
        "data": "",
        "raw_log": "out of gas in location: WriteFlat; gasWanted: 80000000, gasUsed: 80150021: out of gas",
        "logs": [
            {
                "msg_index": 0,
                "log": "",
                "events": []
            }
        ],
        "info": "",
        "gas_wanted": "80000000",
        "gas_used": "80150021",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.bank.v1beta1.MsgSend",
                        "from_address": "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
                        "to_address": "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
                        "amount": [
                            {
                                "denom": "basetcro",
                                "amount": "1000000000"
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
                            "key": "Ax+Rgmd2ta8FxUOoFJ9Dvo3782nMWJzdYP0Jcyrk5XwO"
                        },
                        "mode_info": {
                            "single": {
                                "mode": "SIGN_MODE_DIRECT"
                            }
                        },
                        "sequence": "59"
                    }
                ],
                "fee": {
                    "amount": [
                        {
                            "denom": "basetcro",
                            "amount": "8000000"
                        }
                    ],
                    "gas_limit": "80000000",
                    "payer": "",
                    "granter": ""
                }
            },
            "signatures": [
                "pXF78nzs3lhSlTreiTXrMoqjxF04kidrROyj8bJ8vrxPv2rfecUsx3IHg+P4AVHOstaxarHiP75fWQSB9BBsRw=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
