package usecase_parser_test

const TX_MSG_VOTE_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "4905A138902D497BB2EB34B8FA26F8AF63B81798AAE81FAFFBB24D6CF58E258B",
      "parts": {
        "total": 1,
        "hash": "F2506EA57A23544E645378BBC446EA1467C393C53DEF2424CC01D114D468C622"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "chainmaind",
        "height": "100",
        "time": "2020-11-21T14:22:00.058859Z",
        "last_block_id": {
          "hash": "4A50D07143BEF1B1303FE3D1E35FE88172B158D15D1BF8CA497204C180798875",
          "parts": {
            "total": 1,
            "hash": "C0833123ABCF21F647F5B5339C6E797A3988AE3299606837F59493F8DB20F9B7"
          }
        },
        "last_commit_hash": "A6140C1EABB54D5FC31400B06A2FA6D58FDC9DA223CD09D6B15866A9D0C7DAAE",
        "data_hash": "E28A9F7CCB03B1D0831EC8F5C27CD3CEE165E6FFC4DDAB715EBB4D5E0D6BA3D7",
        "validators_hash": "E95FE501B6BF0271B9C3AE1BCC5502CF4DA75991A836D3823D8092BB4017440B",
        "next_validators_hash": "E95FE501B6BF0271B9C3AE1BCC5502CF4DA75991A836D3823D8092BB4017440B",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "9DD3A750E3C9AEB28CEC5B1B820478370DC8E901A18201EDE4B1765B9682A480",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "9375ACF34EB8767ABF12B2E778F53E3D9BD274CE"
      },
      "data": {
        "txs": [
          "ClEKTwobL2Nvc21vcy5nb3YudjFiZXRhMS5Nc2dWb3RlEjAIARIqY3JvMXRnNHhwcnl5ZTJ2NGZwM3NtcGZjM3Mya3Ftdm5ya3dmeWQ2M3k3GAESVgpOCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAnlNrnAkO3++Zkb+mDTBj7Mip3FUmFw02DggxxFld3BIEgQKAggBEgQQwJoMGkAVfEg8cXBjI3LQ8KuLYWUOKNAtpOndPfwX6hdEXudrviYPplRHQz5mjW/g0rAKjxZPFHAZWPstCUIhx4LLqheF"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "99",
        "round": 0,
        "block_id": {
          "hash": "4A50D07143BEF1B1303FE3D1E35FE88172B158D15D1BF8CA497204C180798875",
          "parts": {
            "total": 1,
            "hash": "C0833123ABCF21F647F5B5339C6E797A3988AE3299606837F59493F8DB20F9B7"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "7182E89DB50F2ACE00FBA243F22BA1159D1DE853",
            "timestamp": "2020-11-21T14:22:00.058859Z",
            "signature": "zXoNX+7aXuu8zkZtVHyxremwGQ1dLiNCeZOGhgDEQM5DZD9ydb4P0uy8ftT1UDtOnArdE8EbMMXW9g8XzDXlCQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "9375ACF34EB8767ABF12B2E778F53E3D9BD274CE",
            "timestamp": "2020-11-21T14:22:00.100116Z",
            "signature": "jk7Q7DWpSRnVMr5CKa7Jr15X0JWtTY+e1dqXs2/qH+gPgfgEltfrK/E6TSD49KSofFxY9uLpBKYisPCMOi/PBg=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_VOTE_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "100",
    "txs_results": [
      {
        "code": 0,
        "data": "CgYKBHZvdGU=",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"vote\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"cro1tg4xpryye2v4fp3smpfc3s2kqmvnrkwfyd63y7\"}]},{\"type\":\"proposal_vote\",\"attributes\":[{\"key\":\"option\",\"value\":\"VOTE_OPTION_YES\"},{\"key\":\"proposal_id\",\"value\":\"1\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "49266",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "dm90ZQ==",
                "index": true
              }
            ]
          },
          {
            "type": "proposal_vote",
            "attributes": [
              {
                "key": "b3B0aW9u",
                "value": "Vk9URV9PUFRJT05fWUVT",
                "index": true
              },
              {
                "key": "cHJvcG9zYWxfaWQ=",
                "value": "MQ==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "Z292ZXJuYW5jZQ==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMXRnNHhwcnl5ZTJ2NGZwM3NtcGZjM3Mya3Ftdm5ya3dmeWQ2M3k3",
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
            "value": "MTI3N2Jhc2Vjcm8=",
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
            "value": "MC4wMzIyNTc5OTg3Mzk0NjgxNTA=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMzAwMDE5NjA1NTc1MDMzOTA=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "ODA2MDEzNzk4OS44MDMwNjk3NDEyNTEwNzM5NzA=",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTI3Nw==",
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
            "value": "MTI3N2Jhc2Vjcm8=",
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
            "value": "NjMuODUwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF6cjh5em0wODZseG1yMnF5eTJ0NjdoOGw0dHJ4cnZ0emRlOGRrbQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Ni4zODUwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF6cjh5em0wODZseG1yMnF5eTJ0NjdoOGw0dHJ4cnZ0emRlOGRrbQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjMuODUwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF6cjh5em0wODZseG1yMnF5eTJ0NjdoOGw0dHJ4cnZ0emRlOGRrbQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkuMzgwNTAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF6cjh5em0wODZseG1yMnF5eTJ0NjdoOGw0dHJ4cnZ0emRlOGRrbQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkzLjgwNTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF6cjh5em0wODZseG1yMnF5eTJ0NjdoOGw0dHJ4cnZ0emRlOGRrbQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkuMzgwNTAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE1cng0dHNsczRwcGdsY3Zqdmd2czlnamZ6a3p4dHA4OGhwc2oycQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkzLjgwNTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE1cng0dHNsczRwcGdsY3Zqdmd2czlnamZ6a3p4dHA4OGhwc2oycQ==",
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

const TX_MSG_VOTE_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.gov.v1beta1.MsgVote",
                    "proposal_id": "1",
                    "voter": "cro1tg4xpryye2v4fp3smpfc3s2kqmvnrkwfyd63y7",
                    "option": "VOTE_OPTION_YES"
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
                        "key": "AnlNrnAkO3++Zkb+mDTBj7Mip3FUmFw02DggxxFld3BI"
                    },
                    "mode_info": {
                        "single": {
                            "mode": "SIGN_MODE_DIRECT"
                        }
                    },
                    "sequence": "0"
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
            "FXxIPHFwYyNy0PCri2FlDijQLaTp3T38F+oXRF7na74mD6ZUR0M+Zo1v4NKwCo8WTxRwGVj7LQlCIceCy6oXhQ=="
        ]
    },
    "tx_response": {
        "height": "100",
        "txhash": "6E6910024B74B16F3B9B14309D7F8CD89AF25E561F0FB3F56380F086218F1759",
        "codespace": "",
        "code": 0,
        "data": "CgYKBHZvdGU=",
        "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"vote\"},{\"key\":\"module\",\"value\":\"governance\"},{\"key\":\"sender\",\"value\":\"cro1tg4xpryye2v4fp3smpfc3s2kqmvnrkwfyd63y7\"}]},{\"type\":\"proposal_vote\",\"attributes\":[{\"key\":\"option\",\"value\":\"VOTE_OPTION_YES\"},{\"key\":\"proposal_id\",\"value\":\"1\"}]}]}]",
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
                                "value": "dm90ZQ==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "proposal_vote",
                        "attributes": [
                            {
                                "key": "b3B0aW9u",
                                "value": "Vk9URV9PUFRJT05fWUVT",
                                "index": true
                            },
                            {
                                "key": "cHJvcG9zYWxfaWQ=",
                                "value": "MQ==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "Z292ZXJuYW5jZQ==",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMXRnNHhwcnl5ZTJ2NGZwM3NtcGZjM3Mya3Ftdm5ya3dmeWQ2M3k3",
                                "index": true
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "49266",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.gov.v1beta1.MsgVote",
                        "proposal_id": "1",
                        "voter": "cro1tg4xpryye2v4fp3smpfc3s2kqmvnrkwfyd63y7",
                        "option": "VOTE_OPTION_YES"
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
                            "key": "AnlNrnAkO3++Zkb+mDTBj7Mip3FUmFw02DggxxFld3BI"
                        },
                        "mode_info": {
                            "single": {
                                "mode": "SIGN_MODE_DIRECT"
                            }
                        },
                        "sequence": "0"
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
                "FXxIPHFwYyNy0PCri2FlDijQLaTp3T38F+oXRF7na74mD6ZUR0M+Zo1v4NKwCo8WTxRwGVj7LQlCIceCy6oXhQ=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
