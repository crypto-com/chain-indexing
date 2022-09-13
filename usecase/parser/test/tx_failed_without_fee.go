package usecase_parser_test

const TX_FAILED_WITHOUT_FEE_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "F3C1861B33027C3195ECD2656197A12D4FF1440D9B11D27CB9DAE57A5336D504",
      "parts": {
        "total": 1,
        "hash": "3FE34ADCD793A12A94E6F62ABB41D650B6530A47C72A3975A9ED66BE3E792EDB"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-1",
        "height": "3245",
        "time": "2020-10-15T14:11:59.011111839Z",
        "last_block_id": {
          "hash": "23A94EE12003E535B524CDDD854ED8E74E9D59088AB92CA8FAB64161E706E392",
          "parts": {
            "total": 1,
            "hash": "519B4B76458EA4243FD5CB5CBACDE0AB80BF160BEC4A4A87086032E554BFFB31"
          }
        },
        "last_commit_hash": "D677148D0CC92B45E67231069A23E0ADC803C4393543ACB70762FBA6864BA2B9",
        "data_hash": "F98100D0134CDAFF58062171BEB53CF2FD15C116D153207D70796D60C298DCDE",
        "validators_hash": "3E64FED13D19A7FEB6D56F12CF731D5B4D1A4CF79CE024A62C1DF8129F42F35A",
        "next_validators_hash": "3E64FED13D19A7FEB6D56F12CF731D5B4D1A4CF79CE024A62C1DF8129F42F35A",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "1FC912E9FA9B06679D65E19D43206ED2A9E28B89FF221BDA69EA6624C4655DD9",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "504C0C3FE72728946911C7956E1B012784446B64"
      },
      "data": {
        "txs": [
          "CowGCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xbDM4emU1Zm1ncmd6dzZybjNhZngzZ3RwYXJlM2pnczhrZTRuNjkSL3Rjcm9jbmNsMWZqYTVuc3h6N2dzcXc0emNjdXV5OHI3cGpuam1jN2RzY2RsMnZ6CpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xbDM4emU1Zm1ncmd6dzZybjNhZngzZ3RwYXJlM2pnczhrZTRuNjkSL3Rjcm9jbmNsMWo3cGVqOGtwbGVtNHd0NTBwNGhmdm5kaHV3NWpwcnh4eHRlbnZyCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xbDM4emU1Zm1ncmd6dzZybjNhZngzZ3RwYXJlM2pnczhrZTRuNjkSL3Rjcm9jbmNsMTZ5emN6M3R5OTRhd3I3bnIydHhlazlkcDJrbHAyYXY5dmg0MzdzCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xbDM4emU1Zm1ncmd6dzZybjNhZngzZ3RwYXJlM2pnczhrZTRuNjkSL3Rjcm9jbmNsMTZrcXIwMDlwdGdrZW42cXN4bnpmbnlqZnNxNnE5N2czdWVkY2VyCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xbDM4emU1Zm1ncmd6dzZybjNhZngzZ3RwYXJlM2pnczhrZTRuNjkSL3Rjcm9jbmNsMWwzOHplNWZtZ3Jnenc2cm4zYWZ4M2d0cGFyZTNqZ3M4cnhrMnp4ElgKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQIS2Ag1W6TNdiWn6qdKaChHSBKXk59N5iefpiIEc+D1mhIECgIIARgFEgQQwJoMGkBpgfqxZqQkt7K9UTyDdUrHsiKokzDCIVARyafw2+DaYmSOtgBrSiBdU+juhoX2ISjWEcFULwORWEgc5K4jhU/j"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "3244",
        "round": 0,
        "block_id": {
          "hash": "23A94EE12003E535B524CDDD854ED8E74E9D59088AB92CA8FAB64161E706E392",
          "parts": {
            "total": 1,
            "hash": "519B4B76458EA4243FD5CB5CBACDE0AB80BF160BEC4A4A87086032E554BFFB31"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "384E5F30F02538C0A34CBFF32F8D5554671C9029",
            "timestamp": "2020-10-15T14:11:59.029101258Z",
            "signature": "SHB0pvPEmRP2DgLDZDK8g8eFhpOfiR5d4j2YUbQSG8vYmuMacGvN79C8cz/c4V4lII+E0ibAdMkrZ8YB+6r6DA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4D9F47C5A19D5550685D2D55DF2FF8C5D7504CEB",
            "timestamp": "2020-10-15T14:11:59.018274046Z",
            "signature": "wjvHGNDCaBeNw+41dMwud2UrNpammKDkT0Ax9Ro3eJARwJBL+Q6/LuD/hMhoydiIdUzm2Qr3onvBRKoIRKbWCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "504C0C3FE72728946911C7956E1B012784446B64",
            "timestamp": "2020-10-15T14:11:59.005970221Z",
            "signature": "u01IYuhjnR2BhQGds0k2DcWKDCMuM466y6RkOsOmDVD4OmiOLAFgvwt4IXNXpJlWbF41NKy8b7Z4jMJv8G67Dg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "6D9E4B4995037D608E365CE90436C24580ABCC33",
            "timestamp": "2020-10-15T14:11:59.011111839Z",
            "signature": "B60/rCK/2YPnPWCEJGHaDrbymwMDT1pg22XoMelBCAgSWGjOg1VZyoVh5Qm0gYXQTG2JE/t+ekkYh06Q3QT2Dw=="
          },
          {
            "block_id_flag": 1,
            "validator_address": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "signature": null
          }
        ]
      }
    }
  }
}`

const TX_FAILED_WITHOUT_FEE_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "3245",
    "txs_results": [
      {
        "code": 11,
        "data": null,
        "log": "out of gas in location: WriteFlat; gasWanted: 200000, gasUsed: 201420: out of gas",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "201420",
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
            "value": "MTY0ODY1MTEzMjFiYXNldGNybw==",
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
            "value": "MC4wMDAwMDUwMTc0Mjg4MzUwOTU=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTMwMDY2ODM3NjA5MTU5OTY=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTA0MDU0OTQ1OTMzOTE1NjEwLjQ5Nzg2MDYzMDc4NzI0Mjk2NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTY0ODY1MTEzMjE=",
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
            "value": "MTY0ODY1MTEzMjFiYXNldGNybw==",
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
            "value": "ODIyNjgyNjY1LjA3MTMyNTM1OTc3MzQ5MzI2NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmtxcjAwOXB0Z2tlbjZxc3huemZueWpmc3E2cTk3ZzN1ZWRjZXI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODIyNjgyNjYuNTA3MTMyNTM1OTc3MzQ5MzI3YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmtxcjAwOXB0Z2tlbjZxc3huemZueWpmc3E2cTk3ZzN1ZWRjZXI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODIyNjgyNjY1LjA3MTMyNTM1OTc3MzQ5MzI2NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmtxcjAwOXB0Z2tlbjZxc3huemZueWpmc3E2cTk3ZzN1ZWRjZXI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwNjIwMTQxLjYxODk0ODc2MjExODg4NDM3MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZmphNW5zeHo3Z3NxdzR6Y2N1dXk4cjdwam5qbWM3ZHNjZGwydno=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwNjIwMTQxNi4xODk0ODc2MjExODg4NDM2OTViYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZmphNW5zeHo3Z3NxdzR6Y2N1dXk4cjdwam5qbWM3ZHNjZGwydno=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwNjIwMTQxLjYxODk0ODc2MjExODg4NDM3MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmtxcjAwOXB0Z2tlbjZxc3huemZueWpmc3E2cTk3ZzN1ZWRjZXI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwNjIwMTQxNi4xODk0ODc2MjExODg4NDM2OTViYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmtxcjAwOXB0Z2tlbjZxc3huemZueWpmc3E2cTk3ZzN1ZWRjZXI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwNjIwMTQxLjYxODk0ODc2MjExODg4NDM3MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxajdwZWo4a3BsZW00d3Q1MHA0aGZ2bmRodXc1anByeHh4dGVudnI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwNjIwMTQxNi4xODk0ODc2MjExODg4NDM2OTViYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxajdwZWo4a3BsZW00d3Q1MHA0aGZ2bmRodXc1anByeHh4dGVudnI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwNjIwMTQxLjYxODk0ODc2MjExODg4NDM3MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNnl6Y3ozdHk5NGF3cjducjJ0eGVrOWRwMmtscDJhdjl2aDQzN3M=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwNjIwMTQxNi4xODk0ODc2MjExODg4NDM2OTViYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNnl6Y3ozdHk5NGF3cjducjJ0eGVrOWRwMmtscDJhdjl2aDQzN3M=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwMjI5OS4xMTcwNzI0MTM5ODA3MzAzMzBiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxbDM4emU1Zm1ncmd6dzZybjNhZngzZ3RwYXJlM2pnczhyeGsyeng=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzkwMjI5OTEuMTcwNzI0MTM5ODA3MzAzMjk2YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxbDM4emU1Zm1ncmd6dzZybjNhZngzZ3RwYXJlM2pnczhyeGsyeng=",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "dGNyb2NuY2xjb25zMXF0MnN6dTkwazJlaHJwcmg0ZnM4ZzZkNzQ5a2dwbjVnbThjMHNz",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NDUw",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MzI0NQ==",
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

const TX_FAILED_WITHOUT_FEE_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                    "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                    "validator_address": "tcrocncl1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dscdl2vz"
                },
                {
                    "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                    "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                    "validator_address": "tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr"
                },
                {
                    "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                    "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                    "validator_address": "tcrocncl16yzcz3ty94awr7nr2txek9dp2klp2av9vh437s"
                },
                {
                    "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                    "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                    "validator_address": "tcrocncl16kqr009ptgken6qsxnzfnyjfsq6q97g3uedcer"
                },
                {
                    "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                    "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                    "validator_address": "tcrocncl1l38ze5fmgrgzw6rn3afx3gtpare3jgs8rxk2zx"
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
                        "key": "AhLYCDVbpM12Jafqp0poKEdIEpeTn03mJ5+mIgRz4PWa"
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
            "aYH6sWakJLeyvVE8g3VKx7IiqJMwwiFQEcmn8Nvg2mJkjrYAa0ogXVPo7oaF9iEo1hHBVC8DkVhIHOSuI4VP4w=="
        ]
    },
    "tx_response": {
        "height": "3245",
        "txhash": "CDBA166168176BF7ECA2EAC9E9B49054F1BF4C8799B8C26CC0B9EE85CB93AF27",
        "codespace": "sdk",
        "code": 11,
        "data": "",
        "raw_log": "out of gas in location: WriteFlat; gasWanted: 200000, gasUsed: 201420: out of gas",
        "logs": [
            {
                "msg_index": 0,
                "log": "",
                "events": []
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "201420",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                        "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                        "validator_address": "tcrocncl1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dscdl2vz"
                    },
                    {
                        "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                        "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                        "validator_address": "tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr"
                    },
                    {
                        "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                        "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                        "validator_address": "tcrocncl16yzcz3ty94awr7nr2txek9dp2klp2av9vh437s"
                    },
                    {
                        "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                        "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                        "validator_address": "tcrocncl16kqr009ptgken6qsxnzfnyjfsq6q97g3uedcer"
                    },
                    {
                        "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
                        "delegator_address": "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
                        "validator_address": "tcrocncl1l38ze5fmgrgzw6rn3afx3gtpare3jgs8rxk2zx"
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
                            "key": "AhLYCDVbpM12Jafqp0poKEdIEpeTn03mJ5+mIgRz4PWa"
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
                "aYH6sWakJLeyvVE8g3VKx7IiqJMwwiFQEcmn8Nvg2mJkjrYAa0ogXVPo7oaF9iEo1hHBVC8DkVhIHOSuI4VP4w=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
