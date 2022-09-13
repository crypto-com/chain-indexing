package usecase_parser_test

const TX_MSG_MULTI_SEND_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "0CB52E156572AA0DB9895CDE286FCED2BC2392F007FDE3C51182A4A19F63FAC6",
      "parts": {
        "total": 1,
        "hash": "2D9CEF684BF4C070C5FCD610DCFF05878633904F7E0C1266E1B22E17AD151559"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-1",
        "height": "421481",
        "time": "2020-11-15T18:37:54.019116751Z",
        "last_block_id": {
          "hash": "D63159B5C24077574CF6D6D5CEAC4192496533B1862BC06C5FF55C84E6BA6392",
          "parts": {
            "total": 1,
            "hash": "311632888BA1A2305A32F4764D9E219F9DCE712E8D23BDC2F4570AC0BE36F8B9"
          }
        },
        "last_commit_hash": "391F7764327E3E1923C5ABE9CC3BA877B7EDDBC0D09853969358E6844FABE06B",
        "data_hash": "9BAF0BB2C1DE1A3A65C9042F0149F7AA1BDBAD9FE9BA92F38C7B350892E85544",
        "validators_hash": "DEBA738701DDCEAFFA5D35C1B8EA80AB88F1DAD9E25AF0DC2D676E483A920D3C",
        "next_validators_hash": "DEBA738701DDCEAFFA5D35C1B8EA80AB88F1DAD9E25AF0DC2D676E483A920D3C",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "507BA4D7DE27C807F2E2F0F24C0968F138A5600AFDED4A81A8EB36C2685DC0D9",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "62E0A1625735852A49D0D42BE3D11EE6EE4EB198"
      },
      "data": {
        "txs": [
          "CqQCCqECCiEvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dNdWx0aVNlbmQS+wEKPQordGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bhIOCghiYXNldGNybxICNTESPAordGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaBINCghiYXNldGNybxIBMRI9Cit0Y3JvMTRtNWE0a3h0MmU4MnVxcXM1Z3RxemEyOWRtNXdxenlhMmp3OXNoEg4KCGJhc2V0Y3JvEgIyMBI9Cit0Y3JvMTRtNWE0a3h0MmU4MnVxcXM1Z3RxemEyOWRtNXdxenlhMmp3OXNoEg4KCGJhc2V0Y3JvEgIzMBJrClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDWaFUuiEMSJ2kZiak1jHG+KRxovvaNCFk3V/EoViQHyYSBAoCCAEYBBIXChEKCGJhc2V0Y3JvEgUyMDAwMBDAmgwaQLIQ3JmCXdkEmWsFnqWK2hB4KdL79m+3eykTDjSGAVNaQcLNrno9xJxZ4OJGgfGVfTgTF8wjJYYtjsYLUIWa2ws="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "421480",
        "round": 0,
        "block_id": {
          "hash": "D63159B5C24077574CF6D6D5CEAC4192496533B1862BC06C5FF55C84E6BA6392",
          "parts": {
            "total": 1,
            "hash": "311632888BA1A2305A32F4764D9E219F9DCE712E8D23BDC2F4570AC0BE36F8B9"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "A1E8AAEBBC82929B852748734BA39D67A62F201B",
            "timestamp": "2020-11-15T18:37:53.979944907Z",
            "signature": "MVauwzHpGpi5PfLO8wQGLdnMKyTFH/i0YrH6AFMWSo0eAh8jZNPePp+X+8xMFvUJAndbrPzDSfEd5Bu26iYwDw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3705DA4F2E53A09025DAA8E6581EFCE851811914",
            "timestamp": "2020-11-15T18:37:54.01555749Z",
            "signature": "cwFtpBZQzpOVqnVA41p/rI308W6ZxaRRYe75Nzhc727tiV4+CDCl0E0MSitwdeyOX43ipRr7vHmWb6gz7EnSDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4B68F098199E7F565B02EF58115FB3CB9BAD52B0",
            "timestamp": "2020-11-15T18:37:54.056322572Z",
            "signature": "jfWp8LLXCd+bf3GCLYwl0dvyoW65bwH7UAL193I8AtJ2w+EeNFwJIukfOCBu94QJ+qheJfdXG1F5TNUKPBIeBw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "504C0C3FE72728946911C7956E1B012784446B64",
            "timestamp": "2020-11-15T18:37:53.895093806Z",
            "signature": "eKMeTxHP/I7QZY0xQwyNcoGdKC3P5R4uOPKRcsHcVTMii0VDVh0TTwkI2qvau1wi3EjOAynjxYoTMPATtFedAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "95CDD1C2F0E79F62745D17A90D9A7B138DC8F922",
            "timestamp": "2020-11-15T18:37:54.0559605Z",
            "signature": "7zpmCzyRGsKrGnKgu1kRed/zd59Kw1YcHhUUF78+o0z9oRefdsJM3BK/tGzAvxhnC9w6I7G/fivw6qTNnCqPDw=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_MULTI_SEND_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "421481",
    "txs_results": [
      {
        "code": 0,
        "data": "CgsKCW11bHRpc2VuZA==",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"multisend\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh\"},{\"key\":\"amount\",\"value\":\"1basetcro\"},{\"key\":\"recipient\",\"value\":\"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh\"},{\"key\":\"amount\",\"value\":\"20basetcro\"},{\"key\":\"recipient\",\"value\":\"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh\"},{\"key\":\"amount\",\"value\":\"30basetcro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "74346",
        "events": [
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
                "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MjAwMDBiYXNldGNybw==",
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
                "key": "YWN0aW9u",
                "value": "bXVsdGlzZW5k",
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
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
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
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MjBiYXNldGNybw==",
                "index": true
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MzBiYXNldGNybw==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "YmFuaw==",
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
            "value": "MTc1OTMyMjE0OTFiYXNldGNybw==",
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
            "value": "MC4wMDA5MTQxOTcwNzIwMjM3MjQ=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTM4Njc0NTI3MzY4NzkzMzA=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTExMDM5OTY5MzA2ODY5NDc4LjI1MjQ0MzQwODgwNzgwNzQxMA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc1OTMyMjE0OTE=",
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
            "value": "MTc1OTMyMjE0OTFiYXNldGNybw==",
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
            "value": "ODc5NjYxMDc0LjU1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDM5ODMwNTM3LjI3NTAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
            "value": "ODc5NjYxMDc0LjU1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
    "validator_updates": [
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "YF5RlzKhq0KUreG18W3NiOMpyvXexmQko3PYFqsdcwQ="
            }
          }
        },
        "power": "197611722"
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

const TX_MSG_MULTI_SEND_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.bank.v1beta1.MsgMultiSend",
                    "inputs": [
                        {
                            "address": "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
                            "coins": [
                                {
                                    "denom": "basetcro",
                                    "amount": "51"
                                }
                            ]
                        }
                    ],
                    "outputs": [
                        {
                            "address": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
                            "coins": [
                                {
                                    "denom": "basetcro",
                                    "amount": "1"
                                }
                            ]
                        },
                        {
                            "address": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
                            "coins": [
                                {
                                    "denom": "basetcro",
                                    "amount": "20"
                                }
                            ]
                        },
                        {
                            "address": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
                            "coins": [
                                {
                                    "denom": "basetcro",
                                    "amount": "30"
                                }
                            ]
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
                        "key": "A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m"
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
                        "denom": "basetcro",
                        "amount": "20000"
                    }
                ],
                "gas_limit": "200000",
                "payer": "",
                "granter": ""
            }
        },
        "signatures": [
            "shDcmYJd2QSZawWepYraEHgp0vv2b7d7KRMONIYBU1pBws2uej3EnFng4kaB8ZV9OBMXzCMlhi2OxgtQhZrbCw=="
        ]
    },
    "tx_response": {
        "height": "421481",
        "txhash": "89D220EC5C551DA81EE6993D05B53F25912B072E7FFD863D93AA742923239A5F",
        "codespace": "",
        "code": 0,
        "data": "CgsKCW11bHRpc2VuZA==",
        "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"multisend\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh\"},{\"key\":\"amount\",\"value\":\"1basetcro\"},{\"key\":\"recipient\",\"value\":\"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh\"},{\"key\":\"amount\",\"value\":\"20basetcro\"},{\"key\":\"recipient\",\"value\":\"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh\"},{\"key\":\"amount\",\"value\":\"30basetcro\"}]}]}]",
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
                                "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MjAwMDBiYXNldGNybw==",
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
                                "key": "YWN0aW9u",
                                "value": "bXVsdGlzZW5k",
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
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
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
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MjBiYXNldGNybw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MzBiYXNldGNybw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "YmFuaw==",
                                "index": true
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "74346",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.bank.v1beta1.MsgMultiSend",
                        "inputs": [
                            {
                                "address": "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
                                "coins": [
                                    {
                                        "denom": "basetcro",
                                        "amount": "51"
                                    }
                                ]
                            }
                        ],
                        "outputs": [
                            {
                                "address": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
                                "coins": [
                                    {
                                        "denom": "basetcro",
                                        "amount": "1"
                                    }
                                ]
                            },
                            {
                                "address": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
                                "coins": [
                                    {
                                        "denom": "basetcro",
                                        "amount": "20"
                                    }
                                ]
                            },
                            {
                                "address": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
                                "coins": [
                                    {
                                        "denom": "basetcro",
                                        "amount": "30"
                                    }
                                ]
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
                            "key": "A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m"
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
                            "denom": "basetcro",
                            "amount": "20000"
                        }
                    ],
                    "gas_limit": "200000",
                    "payer": "",
                    "granter": ""
                }
            },
            "signatures": [
                "shDcmYJd2QSZawWepYraEHgp0vv2b7d7KRMONIYBU1pBws2uej3EnFng4kaB8ZV9OBMXzCMlhi2OxgtQhZrbCw=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
